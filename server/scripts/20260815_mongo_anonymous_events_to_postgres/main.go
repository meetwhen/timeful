// mongo_anonymous_events_to_postgres copies explicitly named anonymous events
// into PostgreSQL. It intentionally never modifies MongoDB source documents.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"timeful/server/models"
	pgstore "timeful/server/postgres"
	"timeful/server/respondents"
)

const usage = "usage: go run ./scripts/20260815_mongo_anonymous_events_to_postgres --apply <mongo-short-id> [<mongo-short-id> ...]"

type configuration struct {
	apply       bool
	mongoURI    string
	mongoDB     string
	postgresURI string
	shortIDs    []string
}

type migrationResult struct {
	SourceShortID string
	TargetShortID string
	Responses     int
}

func main() {
	config, err := parseConfiguration(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.mongoURI))
	if err != nil {
		fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	postgresPool, err := pgxpool.New(ctx, config.postgresURI)
	if err != nil {
		fatal(err)
	}
	defer postgresPool.Close()
	if err := postgresPool.Ping(ctx); err != nil {
		fatal(err)
	}

	results := make([]migrationResult, 0, len(config.shortIDs))
	for _, shortID := range config.shortIDs {
		result, err := migrateEvent(ctx, mongoClient.Database(config.mongoDB), postgresPool, shortID, config.apply)
		if err != nil {
			fatal(fmt.Errorf("migrate %s: %w", shortID, err))
		}
		results = append(results, result)
	}

	if !config.apply {
		fmt.Println("preflight passed; rerun with --apply to write PostgreSQL copies")
		return
	}
	for _, result := range results {
		fmt.Printf("%s -> %s (%d responses)\n", result.SourceShortID, result.TargetShortID, result.Responses)
	}
}

func parseConfiguration(arguments []string) (configuration, error) {
	flags := flag.NewFlagSet("mongo_anonymous_events_to_postgres", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	config := configuration{}
	flags.BoolVar(&config.apply, "apply", false, "write PostgreSQL rows")
	flags.StringVar(&config.mongoURI, "mongo-uri", os.Getenv("MONGODB_URI"), "MongoDB connection URI")
	flags.StringVar(&config.mongoDB, "mongo-database", os.Getenv("MONGODB_DATABASE"), "MongoDB database name")
	flags.StringVar(&config.postgresURI, "postgres-uri", os.Getenv("POSTGRES_APPLICATION_URI"), "PostgreSQL connection URI")
	if err := flags.Parse(arguments); err != nil {
		return config, err
	}
	config.shortIDs = flags.Args()
	if config.mongoURI == "" || config.mongoDB == "" || config.postgresURI == "" {
		return config, errors.New("MONGODB_URI, MONGODB_DATABASE, and POSTGRES_APPLICATION_URI must be set")
	}
	if len(config.shortIDs) == 0 {
		return config, errors.New("at least one Mongo short ID is required")
	}
	for _, shortID := range config.shortIDs {
		if strings.TrimSpace(shortID) == "" {
			return config, errors.New("short IDs cannot be empty")
		}
	}
	return config, nil
}

func migrateEvent(ctx context.Context, database *mongo.Database, pool *pgxpool.Pool, shortID string, apply bool) (migrationResult, error) {
	var event models.Event
	err := database.Collection("events").FindOne(ctx, bson.M{"shortId": shortID}).Decode(&event)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return migrationResult{}, errors.New("source event not found")
	}
	if err != nil {
		return migrationResult{}, err
	}
	if err := validateEvent(event); err != nil {
		return migrationResult{}, err
	}

	responses, err := loadResponses(ctx, database, event.Id)
	if err != nil {
		return migrationResult{}, err
	}
	if event.NumResponses == nil || *event.NumResponses != len(responses) {
		return migrationResult{}, fmt.Errorf("numResponses=%v does not match %d stored responses", event.NumResponses, len(responses))
	}

	migrated, err := buildEvent(event)
	if err != nil {
		return migrationResult{}, err
	}
	migratedResponses := make([]migratedResponse, 0, len(responses))
	for _, response := range responses {
		migratedResponse, err := buildResponse(response)
		if err != nil {
			return migrationResult{}, err
		}
		migratedResponses = append(migratedResponses, migratedResponse)
	}

	result := migrationResult{SourceShortID: shortID, Responses: len(migratedResponses)}
	if !apply {
		return result, nil
	}
	tx, err := pool.Begin(ctx)
	if err != nil {
		return migrationResult{}, err
	}
	defer tx.Rollback(ctx) // Commit below makes this a no-op.
	postgresID, postgresShortID, err := insertEvent(ctx, tx, migrated)
	if err != nil {
		return migrationResult{}, err
	}
	for _, response := range migratedResponses {
		if err := insertResponse(ctx, tx, postgresID, response); err != nil {
			return migrationResult{}, err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return migrationResult{}, err
	}
	result.TargetShortID = postgresShortID
	return result, nil
}

func validateEvent(event models.Event) error {
	if event.Id.IsZero() || event.ShortId == nil || *event.ShortId == "" {
		return errors.New("source event has no stable identifier")
	}
	if event.OwnerId.IsZero() == false {
		return errors.New("only anonymous events can move to PostgreSQL")
	}
	if event.Type != models.SPECIFIC_DATES && event.Type != models.DOW {
		return fmt.Errorf("unsupported event type %q", event.Type)
	}
	if event.DaysOnly == nil || !*event.DaysOnly {
		if event.EventTimezone == nil || *event.EventTimezone == "" || event.SlotGeneration == nil || event.TimedRecurrence == nil || len(event.ActiveSlots) == 0 {
			return errors.New("timed event is missing canonical slot metadata")
		}
	}
	return nil
}

type migratedEvent struct {
	name             string
	eventType        string
	isArchived       bool
	isDeleted        bool
	numResponses     int
	scheduleVersion  int
	creatorPosthogID *string
	payload          []byte
	createdAt        time.Time
	updatedAt        time.Time
}

func buildEvent(event models.Event) (migratedEvent, error) {
	payloadEvent := event
	payloadEvent.Id = primitive.NilObjectID
	payloadEvent.ShortId = nil
	payloadEvent.OwnerId = primitive.NilObjectID
	payloadEvent.NumResponses = nil
	payloadEvent.ResponsesMap = nil
	payloadEvent.Attendees = nil
	payloadEvent.HasResponded = nil
	payload, err := json.Marshal(payloadEvent)
	if err != nil {
		return migratedEvent{}, err
	}
	responses := 0
	if event.NumResponses != nil {
		responses = *event.NumResponses
	}
	scheduleVersion := event.ScheduleVersion
	if scheduleVersion == 0 {
		scheduleVersion = 1
	}
	createdAt := event.Id.Timestamp().UTC()
	return migratedEvent{
		name:             event.Name,
		eventType:        string(event.Type),
		isArchived:       boolValue(event.IsArchived),
		isDeleted:        boolValue(event.IsDeleted),
		numResponses:     responses,
		scheduleVersion:  scheduleVersion,
		creatorPosthogID: event.CreatorPosthogId,
		payload:          payload,
		createdAt:        createdAt,
		updatedAt:        createdAt,
	}, nil
}

type migratedResponse struct {
	respondentKind     string
	accountUserID      *string
	guestID            *string
	canonicalGuestName *string
	guestEditPolicy    *string
	guestOwnershipMode *string
	guestEditToken     *string
	payload            []byte
	createdAt          time.Time
	updatedAt          time.Time
}

func loadResponses(ctx context.Context, database *mongo.Database, eventID primitive.ObjectID) ([]models.EventResponse, error) {
	cursor, err := database.Collection("eventResponses").Find(ctx, bson.M{"eventId": eventID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var responses []models.EventResponse
	if err := cursor.All(ctx, &responses); err != nil {
		return nil, err
	}
	return responses, nil
}

func buildResponse(stored models.EventResponse) (migratedResponse, error) {
	if stored.Response == nil {
		return migratedResponse{}, fmt.Errorf("response %s has no payload", stored.Id.Hex())
	}
	response := *stored.Response
	availability, ifNeeded := normalizeAvailability(response.Availability, response.IfNeeded)
	response.Availability, response.IfNeeded = availability, ifNeeded
	createdAt := stored.Id.Timestamp().UTC()

	if response.GuestId != "" || response.Name != "" {
		validated := respondents.ValidateGuestName(response.Name)
		if validated.Code != respondents.GuestNameValid {
			return migratedResponse{}, fmt.Errorf("response %s has invalid guest name", stored.Id.Hex())
		}
		response.Name = validated.Name
		if response.GuestOwnershipMode == "token" && (response.GuestId == "" || response.GuestEditToken == "") {
			return migratedResponse{}, fmt.Errorf("response %s has incomplete token ownership", stored.Id.Hex())
		}
		if response.GuestEditPolicy != "" && response.GuestEditPolicy != "open" && response.GuestEditPolicy != "protected" {
			return migratedResponse{}, fmt.Errorf("response %s has invalid guest edit policy", stored.Id.Hex())
		}
		payload, err := json.Marshal(response)
		if err != nil {
			return migratedResponse{}, err
		}
		return migratedResponse{
			respondentKind:     pgstore.RespondentKindGuest,
			guestID:            optionalString(response.GuestId),
			canonicalGuestName: &validated.Name,
			guestEditPolicy:    optionalString(response.GuestEditPolicy),
			guestOwnershipMode: optionalString(response.GuestOwnershipMode),
			guestEditToken:     optionalString(response.GuestEditToken),
			payload:            payload,
			createdAt:          createdAt,
			updatedAt:          createdAt,
		}, nil
	}

	accountID := response.UserId
	if accountID.IsZero() {
		parsed, err := primitive.ObjectIDFromHex(stored.UserId)
		if err != nil {
			return migratedResponse{}, fmt.Errorf("response %s has no account or guest identity", stored.Id.Hex())
		}
		accountID = parsed
	}
	response.UserId = accountID
	payload, err := json.Marshal(response)
	if err != nil {
		return migratedResponse{}, err
	}
	accountIDString := accountID.Hex()
	return migratedResponse{
		respondentKind: pgstore.RespondentKindAccount,
		accountUserID:  &accountIDString,
		payload:        payload,
		createdAt:      createdAt,
		updatedAt:      createdAt,
	}, nil
}

func insertEvent(ctx context.Context, tx pgx.Tx, event migratedEvent) (string, string, error) {
	for attempt := 0; attempt < 5; attempt++ {
		shortID, err := pgstore.GenerateEventShortID()
		if err != nil {
			return "", "", err
		}
		var id string
		err = tx.QueryRow(ctx, `INSERT INTO postgres_events (short_id, owner_external_id, name, type, is_archived, is_deleted, num_responses, schedule_version, creator_posthog_id, payload, created_at, updated_at)
VALUES ($1, NULL, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`, shortID, event.name, event.eventType, event.isArchived, event.isDeleted, event.numResponses, event.scheduleVersion, event.creatorPosthogID, event.payload, event.createdAt, event.updatedAt).Scan(&id)
		if err == nil {
			return id, shortID, nil
		}
		if !isUniqueViolation(err) {
			return "", "", err
		}
	}
	return "", "", errors.New("generate unique PostgreSQL event identifier")
}

func insertResponse(ctx context.Context, tx pgx.Tx, eventID string, response migratedResponse) error {
	_, err := tx.Exec(ctx, `INSERT INTO postgres_event_responses (event_id, respondent_kind, account_user_id, guest_id, canonical_guest_name, guest_edit_policy, guest_ownership_mode, guest_edit_token, payload, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, eventID, response.respondentKind, response.accountUserID, response.guestID, response.canonicalGuestName, response.guestEditPolicy, response.guestOwnershipMode, response.guestEditToken, response.payload, response.createdAt, response.updatedAt)
	return err
}

func normalizeAvailability(availability, ifNeeded []primitive.DateTime) ([]primitive.DateTime, []primitive.DateTime) {
	available := deduplicate(availability, nil)
	availableSet := make(map[primitive.DateTime]struct{}, len(available))
	for _, value := range available {
		availableSet[value] = struct{}{}
	}
	return available, deduplicate(ifNeeded, availableSet)
}

func deduplicate(values []primitive.DateTime, excluded map[primitive.DateTime]struct{}) []primitive.DateTime {
	result := make([]primitive.DateTime, 0, len(values))
	seen := make(map[primitive.DateTime]struct{}, len(values))
	for _, value := range values {
		if _, exists := excluded[value]; exists {
			continue
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func boolValue(value *bool) bool {
	return value != nil && *value
}

func optionalString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func isUniqueViolation(err error) bool {
	var postgresError *pgconn.PgError
	return errors.As(err, &postgresError) && postgresError.Code == "23505"
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
