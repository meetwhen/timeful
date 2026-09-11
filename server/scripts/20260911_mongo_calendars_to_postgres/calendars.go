package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"timeful/server/models"
	pgstore "timeful/server/postgres"
	"timeful/server/scripts/internal/legacybson"
	"timeful/server/utils"
)

// quarantineRecord is one unmigratable retained unit.
type quarantineRecord struct {
	legacyID string
	reason   string
	detail   string
}

// unitCalendar is one owner account and every retained calendar record it owns.
// It is copied in a single PostgreSQL transaction that also records the ledger
// completion marker.
type unitCalendar struct {
	legacyID    string
	externalID  string
	targetID    string
	connections []pgstore.CalendarAccount
	preferences *pgstore.CalendarPreferences
	quarantine  []quarantineRecord
}

// empty reports whether the unit carries no PostgreSQL target rows.
func (unit unitCalendar) empty() bool {
	return len(unit.connections) == 0 && unit.preferences == nil
}

// migrateCalendars pages the retained MongoDB users collection and migrates one
// owner account per unit. It returns complete=false when --limit stopped the run
// before the source was exhausted so the caller can resume.
func (m *migrator) migrateCalendars(ctx context.Context, config configuration) (migrationSummary, bool, error) {
	collection := m.database.Collection("users")
	summary := migrationSummary{}

	// hasCursor separates "no page read yet" from a legitimate zero ObjectID.
	var lastID primitive.ObjectID
	hasCursor := false
	for {
		cursor, err := collection.Find(ctx, pageFilter(lastID, hasCursor), options.Find().SetSort(bson.M{"_id": 1}).SetLimit(config.batchSize))
		if err != nil {
			return summary, false, err
		}
		users := make([]legacybson.User, 0, config.batchSize)
		if err := cursor.All(ctx, &users); err != nil {
			cursor.Close(ctx)
			return summary, false, err
		}
		cursor.Close(ctx)
		if len(users) == 0 {
			break
		}

		for _, user := range users {
			summary.Scanned++
			legacyID := user.Id.Hex()
			if _, completed, err := m.completedUnit(ctx, ledgerKindCalendarAccount, legacyID); err != nil {
				return summary, false, err
			} else if completed {
				summary.Skipped++
				lastID = user.Id
				hasCursor = true
				continue
			}

			unit, err := m.buildCalendarUnit(ctx, user)
			if err != nil {
				return summary, false, err
			}
			summary.Quarantined += len(unit.quarantine)
			if !unit.empty() || len(unit.quarantine) > 0 {
				if m.apply {
					if err := m.persistCalendarUnit(ctx, unit); err != nil {
						return summary, false, err
					}
				}
				if !unit.empty() {
					summary.Migrated++
				}
			}

			// Advance the source cursor only after the unit committed, so an
			// interruption resumes from the first incomplete account.
			lastID = user.Id
			hasCursor = true
			if config.limit > 0 && summary.Migrated >= config.limit {
				return summary, false, nil
			}
		}
		if int64(len(users)) < config.batchSize {
			break
		}
	}
	return summary, true, nil
}

// buildCalendarUnit resolves the owner through the PostgreSQL account and
// converts every retained connection, credential, sub-calendar, and preference.
// A retained document whose owner has no PostgreSQL account is quarantined and
// never creates, merges, or renames an account.
func (m *migrator) buildCalendarUnit(ctx context.Context, user legacybson.User) (unitCalendar, error) {
	unit := unitCalendar{legacyID: user.Id.Hex(), externalID: user.Id.Hex()}
	if !hasRetainedCalendarData(user) {
		return unit, nil
	}

	identity, err := m.ownerPlatformIdentityID(ctx, unit.externalID)
	if err != nil {
		return unit, err
	}
	if identity == "" {
		unit.quarantine = append(unit.quarantine, quarantineRecord{
			legacyID: unit.legacyID,
			reason:   reasonMissingOwnerAccount,
			detail:   "no PostgreSQL account resolves the retained calendar owner",
		})
		return unit, nil
	}
	unit.targetID = identity

	connections, err := convertCalendarConnections(user)
	if err != nil {
		return unit, err
	}
	unit.connections = connections
	preferences, err := convertCalendarPreferences(user)
	if err != nil {
		return unit, err
	}
	unit.preferences = preferences
	return unit, nil
}

// ownerPlatformIdentityID resolves the PostgreSQL platform identity backing a
// retained owner only when a PostgreSQL account row exists. A missing identity
// or account returns the empty string so the unit is quarantined rather than
// invented.
func (m *migrator) ownerPlatformIdentityID(ctx context.Context, externalUserID string) (string, error) {
	var identityID string
	err := m.pool.QueryRow(ctx, `SELECT p.id FROM platform_identities p
JOIN accounts a ON a.platform_identity_id = p.id
WHERE p.external_user_id = $1
ORDER BY a.created_at, a.id LIMIT 1`, externalUserID).Scan(&identityID)
	if err != nil {
		if isNoRows(err) {
			return "", nil
		}
		return "", err
	}
	return identityID, nil
}

// persistCalendarUnit writes every target row and the ledger completion marker
// in one transaction. A unit that is quarantined but carries no target rows only
// records the quarantine decision.
func (m *migrator) persistCalendarUnit(ctx context.Context, unit unitCalendar) error {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, record := range unit.quarantine {
		if err := m.recordQuarantine(ctx, tx, ledgerKindCalendarAccount, record.legacyID, record.reason, record.detail); err != nil {
			return err
		}
	}
	if unit.empty() {
		return tx.Commit(ctx)
	}

	repository := pgstore.NewRepositoryFromTx(tx)
	for i := range unit.connections {
		connection := &unit.connections[i]
		if err := repository.UpsertCalendarAccount(ctx, unit.externalID, connection); err != nil {
			return fmt.Errorf("upsert calendar connection %s: %w", connection.CalendarKey, err)
		}
		for subCalendarID, subCalendar := range connection.SubCalendars {
			sub := subCalendar
			if err := repository.UpsertCalendarSubCalendar(ctx, unit.externalID, connection.CalendarKey, &sub); err != nil {
				return fmt.Errorf("upsert sub-calendar %s: %w", subCalendarID, err)
			}
		}
	}
	if unit.preferences != nil {
		if err := repository.UpsertCalendarPreferences(ctx, unit.externalID, unit.preferences); err != nil {
			return fmt.Errorf("upsert calendar preferences: %w", err)
		}
	}
	if err := m.recordLedger(ctx, tx, ledgerKindCalendarAccount, unit.legacyID, unit.targetID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// hasRetainedCalendarData reports whether a retained user document carries any
// calendar integration field the retained-data contract moves to PostgreSQL.
func hasRetainedCalendarData(user legacybson.User) bool {
	return len(user.CalendarAccounts) > 0 ||
		user.PrimaryAccountKey != nil ||
		user.TokenOrigin != "" ||
		user.CalendarOptions != nil
}

// convertCalendarConnections converts every retained connection. Connections
// are ordered by their legacy key so a unit writes deterministically.
func convertCalendarConnections(user legacybson.User) ([]pgstore.CalendarAccount, error) {
	connections := make([]pgstore.CalendarAccount, 0, len(user.CalendarAccounts))
	for key, account := range user.CalendarAccounts {
		converted, err := convertCalendarConnection(key, account)
		if err != nil {
			return nil, err
		}
		connections = append(connections, converted)
	}
	sort.Slice(connections, func(i, j int) bool { return connections[i].CalendarKey < connections[j].CalendarKey })
	return connections, nil
}

// convertCalendarConnection translates one retained connection, decrypting the
// legacy AES-CFB Apple password so the repository re-encrypts it with the
// AES-256-GCM envelope. OAuth2 tokens and the ICS feed URL are plaintext in the
// retained document and are encrypted by the repository on write.
func convertCalendarConnection(key string, account legacybson.CalendarAccount) (pgstore.CalendarAccount, error) {
	calendarKey := strings.TrimSpace(key)
	if calendarKey == "" {
		calendarKey = utils.GetCalendarAccountKey(account.Email, models.CalendarType(account.CalendarType))
	}
	converted := pgstore.CalendarAccount{
		CalendarKey:  calendarKey,
		CalendarType: string(account.CalendarType),
		Email:        account.Email,
		Picture:      account.Picture,
		Enabled:      account.Enabled,
	}
	if account.OAuth2CalendarAuth != nil {
		credentials := &pgstore.CalendarOAuth2Credentials{
			AccessToken:  account.OAuth2CalendarAuth.AccessToken,
			RefreshToken: account.OAuth2CalendarAuth.RefreshToken,
			Scope:        account.OAuth2CalendarAuth.Scope,
		}
		if account.OAuth2CalendarAuth.AccessTokenExpireDate > 0 {
			expiresAt := account.OAuth2CalendarAuth.AccessTokenExpireDate.Time().UTC()
			credentials.AccessTokenExpiresAt = &expiresAt
		}
		converted.OAuth2 = credentials
	}
	if account.AppleCalendarAuth != nil {
		password := account.AppleCalendarAuth.Password
		if password != "" {
			decrypted, err := decryptLegacyCFB(password)
			if err != nil {
				return converted, fmt.Errorf("decrypt retained Apple password for %s: %w", calendarKey, err)
			}
			password = decrypted
		}
		if converted.Email == "" {
			converted.Email = account.AppleCalendarAuth.Email
		}
		converted.Apple = &pgstore.CalendarAppleCredentials{Password: password}
	}
	if account.ICSCalendarAuth != nil {
		if converted.Email == "" {
			converted.Email = strings.TrimSpace(account.ICSCalendarAuth.Label)
		}
		converted.ICS = &pgstore.CalendarICSCredentials{FeedURL: account.ICSCalendarAuth.FeedURL}
	}
	if account.SubCalendars != nil {
		converted.SubCalendars = make(map[string]pgstore.CalendarSubCalendar, len(*account.SubCalendars))
		for id, sub := range *account.SubCalendars {
			converted.SubCalendars[id] = pgstore.CalendarSubCalendar{SubCalendarID: id, Name: sub.Name, Enabled: sub.Enabled}
		}
	}
	return converted, nil
}

// convertCalendarPreferences converts the retained preference fields. A user
// with no retained preference field gets no preference row.
func convertCalendarPreferences(user legacybson.User) (*pgstore.CalendarPreferences, error) {
	if user.PrimaryAccountKey == nil && user.TokenOrigin == "" && user.CalendarOptions == nil {
		return nil, nil
	}
	preferences := &pgstore.CalendarPreferences{PrimaryAccountKey: user.PrimaryAccountKey}
	if user.TokenOrigin != "" {
		origin := string(user.TokenOrigin)
		preferences.TokenOrigin = &origin
	}
	if user.CalendarOptions != nil {
		encoded, err := json.Marshal(user.CalendarOptions)
		if err != nil {
			return nil, err
		}
		preferences.CalendarOptions = encoded
	}
	return preferences, nil
}
