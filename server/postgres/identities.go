package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"timeful/server/models"
)

// CreatePlatformIdentity inserts a new platform identity. The identity's
// uuidv7() default supplies the account identifier, so sign-in stores no
// separate external value.
func (r *Repository) CreatePlatformIdentity(ctx context.Context) (*PlatformIdentity, error) {
	value := &PlatformIdentity{}
	err := r.db.QueryRow(ctx, `INSERT INTO platform_identities DEFAULT VALUES
 RETURNING id, created_at`).Scan(&value.ID, &value.CreatedAt)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// GetPlatformIdentity resolves the platform identity a sign-in session carries.
// A value that is not a canonical UUID, and a uuid with no live identity such
// as a deleted account's, both report no identity: there is no compatibility
// lookup for the retired 24-character external identifier.
func (r *Repository) GetPlatformIdentity(ctx context.Context, platformIdentityID string) (*PlatformIdentity, error) {
	if !validPlatformIdentityID(platformIdentityID) {
		return nil, pgx.ErrNoRows
	}
	value := &PlatformIdentity{}
	err := r.db.QueryRow(ctx, `SELECT id, created_at FROM platform_identities WHERE id = $1`, platformIdentityID).Scan(&value.ID, &value.CreatedAt)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// validPlatformIdentityID reports whether a value is the canonical wire form of
// a platform identity UUID. Non-canonical session values resolve to no account
// instead of reaching PostgreSQL as an invalid uuid literal.
func validPlatformIdentityID(value string) bool {
	_, ok := models.ParseUUID(value)
	return ok
}

func (r *Repository) CreateEventVisitorIdentity(ctx context.Context, eventID string) (*EventVisitorIdentity, error) {
	value := &EventVisitorIdentity{}
	err := r.db.QueryRow(ctx, `INSERT INTO event_visitor_identities (event_id) VALUES ($1)
 RETURNING id, event_id, public_id, platform_identity_id, created_at`, eventID).Scan(&value.ID, &value.EventID, &value.PublicID, &value.PlatformIdentityID, &value.CreatedAt)
	return value, err
}

func (r *Repository) GetEventVisitorIdentity(ctx context.Context, eventID, publicID string) (*EventVisitorIdentity, error) {
	value := &EventVisitorIdentity{}
	err := r.db.QueryRow(ctx, `SELECT id, event_id, public_id, platform_identity_id, created_at
 FROM event_visitor_identities WHERE event_id = $1 AND public_id::text = $2`, eventID, publicID).Scan(&value.ID, &value.EventID, &value.PublicID, &value.PlatformIdentityID, &value.CreatedAt)
	return value, err
}

// AssociateEventVisitorIdentity must be called only after proving browser authority.
// An association cannot be silently reassigned to another account.
func (r *Repository) AssociateEventVisitorIdentity(ctx context.Context, visitorID, platformID string) error {
	result, err := r.db.Exec(ctx, `UPDATE event_visitor_identities SET platform_identity_id = $2
 WHERE id = $1 AND (platform_identity_id IS NULL OR platform_identity_id = $2)`, visitorID, platformID)
	if err == nil && result.RowsAffected() != 1 {
		return errors.New("visitor is associated with another platform identity")
	}
	return err
}

func (r *Repository) CreateEventVisitorCredential(ctx context.Context, value *EventVisitorCredential) error {
	if value == nil || len(value.CredentialHash) != 32 {
		return errors.New("credential SHA-256 hash is required")
	}
	if value.Kind == "" {
		value.Kind = CredentialKindBase
	}
	return r.db.QueryRow(ctx, `INSERT INTO event_visitor_credentials (event_visitor_identity_id, credential_hash, kind, grants_owner)
 VALUES ($1, $2, $3, $4) RETURNING id, created_at`, value.EventVisitorIdentityID, value.CredentialHash, value.Kind, value.GrantsOwner).Scan(&value.ID, &value.CreatedAt)
}

func (r *Repository) GetEventVisitorCredential(ctx context.Context, visitorID, credentialID string) (*EventVisitorCredential, error) {
	value := &EventVisitorCredential{}
	err := r.db.QueryRow(ctx, `SELECT id, event_visitor_identity_id, credential_hash, created_at, revoked_at, kind, grants_owner
 FROM event_visitor_credentials WHERE event_visitor_identity_id = $1 AND id::text = $2`, visitorID, credentialID).Scan(&value.ID, &value.EventVisitorIdentityID, &value.CredentialHash, &value.CreatedAt, &value.RevokedAt, &value.Kind, &value.GrantsOwner)
	return value, err
}

func (r *Repository) RevokeEventVisitorCredentials(ctx context.Context, visitorID string) error {
	_, err := r.db.Exec(ctx, `UPDATE event_visitor_credentials SET revoked_at = clock_timestamp()
 WHERE event_visitor_identity_id = $1 AND revoked_at IS NULL`, visitorID)
	return err
}

// VisitorBelongsToAccount reports whether one Event Visitor Identity is
// associated with the given platform identity. A non-canonical identifier
// belongs to no account.
func (r *Repository) VisitorBelongsToAccount(ctx context.Context, visitorID, platformIdentityID string) (bool, error) {
	if !validPlatformIdentityID(platformIdentityID) {
		return false, nil
	}
	var authorized bool
	err := r.db.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM event_visitor_identities
 WHERE id = $1 AND platform_identity_id = $2)`, visitorID, platformIdentityID).Scan(&authorized)
	return authorized, err
}

func (r *Repository) GetResponseByPublicID(ctx context.Context, eventID, publicID string) (*Response, error) {
	return r.getResponse(ctx, `event_id = $1 AND public_id::text = $2`, eventID, publicID)
}

// LockEvent serializes response count changes across concurrent requests.
func (r *Repository) LockEvent(ctx context.Context, eventID string) (*Event, error) {
	var id string
	if err := r.db.QueryRow(ctx, `SELECT id FROM postgres_events WHERE id = $1 FOR UPDATE`, eventID).Scan(&id); err != nil {
		return nil, err
	}
	return r.GetEventByID(ctx, id)
}

// SetEventOwnerToken is used only during event creation; existing EVCCs cannot recover a token.
func (r *Repository) SetEventOwnerToken(ctx context.Context, eventID string, hash []byte) error {
	if len(hash) != 32 {
		return errors.New("owner token SHA-256 hash is required")
	}
	_, err := r.db.Exec(ctx, `UPDATE postgres_events SET owner_edit_token_hash = $2 WHERE id = $1 AND owner_edit_token_hash IS NULL`, eventID, hash)
	return err
}

// AssociateEventOwner must run under the event row lock after token proof.
// It deliberately does not reassign any Event Visitor Identity or response.
func (r *Repository) AssociateEventOwner(ctx context.Context, eventID, platformID string) error {
	_, err := r.db.Exec(ctx, `UPDATE postgres_events SET owner_platform_identity_id = $2, updated_at = clock_timestamp() WHERE id = $1`, eventID, platformID)
	return err
}

// EventOwnerBelongsToAccount reports whether an event is owned by the given
// platform identity. A non-canonical identifier owns no event.
func (r *Repository) EventOwnerBelongsToAccount(ctx context.Context, eventID, platformIdentityID string) (bool, error) {
	if !validPlatformIdentityID(platformIdentityID) {
		return false, nil
	}
	var authorized bool
	err := r.db.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM postgres_events
 WHERE id = $1 AND owner_platform_identity_id = $2)`, eventID, platformIdentityID).Scan(&authorized)
	return authorized, err
}
