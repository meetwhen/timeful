// Package postgres contains persistence types for PostgreSQL-owned anonymous
// events. These types deliberately do not reuse MongoDB BSON models.
package postgres

import (
	"encoding/json"
	"time"
)

const (
	EventTypeSpecificDates = "specific_dates"
	EventTypeDayOfWeek     = "dow"

	RespondentKindAccount = "account"
	RespondentKindGuest   = "guest"
)

// Event is the storage aggregate for a PostgreSQL-owned anonymous poll.
// ID is a hidden UUIDv7; ShortID is the sole public event identifier. Payload
// holds compatibility fields whose JSON shape matters.
type Event struct {
	ID               string
	ShortID          string
	OwnerExternalID  *string
	Name             string
	Type             string
	IsArchived       bool
	IsDeleted        bool
	NumResponses     int
	ScheduleVersion  int
	CreatorPosthogID *string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Payload          json.RawMessage
}

// Response is stored independently so response mutations and event response
// counts can be committed in one transaction. Identity columns drive lookup
// and uniqueness; Payload retains the current response wire shape.
type Response struct {
	ID                 string
	EventID            string
	RespondentKind     string
	AccountUserID      *string
	GuestID            *string
	CanonicalGuestName *string
	GuestEditPolicy    *string
	GuestOwnershipMode *string
	GuestEditToken     *string
	Payload            json.RawMessage
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
