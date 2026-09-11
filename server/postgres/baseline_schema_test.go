package postgres

import (
	"testing"
)

// TestBaselineVisitorIdentityAndOwnerConstraints proves the baseline schema
// enforces the visitor identity and event owner relations the runtime depends
// on: every response belongs to a visitor identity of its own event, an event
// owner relation is event-scoped, and a base credential can never carry owner
// powers. It also proves the compatibility response columns remain usable.
func TestBaselineVisitorIdentityAndOwnerConstraints(t *testing.T) {
	ctx, _, tx := newMigrationTestRepository(t)
	eventID := seedSignupEvent(t, ctx, tx, signupTestShortID(t))
	otherEventID := seedSignupEvent(t, ctx, tx, signupTestShortID(t))
	visitorID := seedSignupVisitor(t, ctx, tx, eventID)

	// A response must carry a visitor identity scoped to the same event.
	expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `INSERT INTO postgres_event_responses (event_id, event_visitor_identity_id) VALUES ($1, $2)`, otherEventID, visitorID)
		return err
	})
	if _, err := tx.Exec(ctx, `INSERT INTO postgres_event_responses (event_id, event_visitor_identity_id, respondent_kind, account_user_id, guest_edit_token, payload)
VALUES ($1, $2, 'account', 'compat-account', 'compat-token', '{}'::jsonb)`, eventID, visitorID); err != nil {
		t.Fatalf("compatibility response insert: %v", err)
	}

	// The owner relation must reference a visitor identity of the same event.
	expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `UPDATE postgres_events SET owner_event_visitor_identity_id = $2 WHERE id = $1`, otherEventID, visitorID)
		return err
	})
	if _, err := tx.Exec(ctx, `UPDATE postgres_events SET owner_event_visitor_identity_id = $2 WHERE id = $1`, eventID, visitorID); err != nil {
		t.Fatalf("owner association: %v", err)
	}

	// A base credential can never grant owner powers; a granted credential can.
	if _, err := tx.Exec(ctx, `INSERT INTO event_visitor_credentials (event_visitor_identity_id, credential_hash)
VALUES ($1, decode(repeat('ab',32),'hex'))`, visitorID); err != nil {
		t.Fatal(err)
	}
	expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `UPDATE event_visitor_credentials SET grants_owner = true`)
		return err
	})
	if _, err := tx.Exec(ctx, `INSERT INTO event_visitor_credentials (event_visitor_identity_id, credential_hash, kind, grants_owner)
VALUES ($1, decode(repeat('cd',32),'hex'), 'granted', true)`, visitorID); err != nil {
		t.Fatalf("granted credential with owner powers: %v", err)
	}
}
