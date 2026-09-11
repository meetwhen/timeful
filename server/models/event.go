package models

import (
	"encoding/json"
)

type EventType string

const (
	SPECIFIC_DATES EventType = "specific_dates"
	DOW            EventType = "dow"
	GROUP          EventType = "group"
)

// Object containing information associated with the remindee
type Remindee struct {
	Email     string   `json:"email"`
	TaskIds   []string `json:"-"` // Task IDs of the scheduled emails
	Responded *bool    `json:"responded"`
}

type SignUpBlock struct {
	Id        ID        `json:"_id"`
	Name      string    `json:"name"`
	Capacity  *int      `json:"capacity"`
	StartDate *DateTime `json:"startDate"`
	EndDate   *DateTime `json:"endDate"`
}

type SignUpResponse struct {
	// The IDs of the sign up blocks that the user has signed up for
	SignUpBlockIds []ID `json:"signUpBlockIds"`

	// Guest information
	Name  string `json:"name"`
	Email string `json:"email"`

	// User information
	UserId ID    `json:"userId"`
	User   *User `json:"user"`
}

type SlotGeneration struct {
	StartTimeLocal       string `json:"startTimeLocal"`
	EndTimeLocal         string `json:"endTimeLocal"`
	TimeIncrementMinutes int    `json:"timeIncrementMinutes"`
}

type TimedRecurrence struct {
	Kind               string   `json:"kind"`
	SelectedDays       []string `json:"selectedDays"`
	SelectedDaysOfWeek []int    `json:"selectedDaysOfWeek"`
	StartOnMonday      *bool    `json:"startOnMonday"`
}

// Representation of an Event in the authoritative event store.
type Event struct {
	Id          ID      `json:"_id"`
	ShortId     *string `json:"shortId"`
	OwnerId     ID      `json:"ownerId"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IsArchived  *bool   `json:"isArchived"`
	IsDeleted   *bool   `json:"isDeleted"`

	Duration                 *float32         `json:"duration"`
	Dates                    []DateTime       `json:"dates"`
	NotificationsEnabled     *bool            `json:"notificationsEnabled"`
	SendEmailAfterXResponses *int             `json:"sendEmailAfterXResponses"`
	When2meetHref            *string          `json:"when2meetHref"`
	CollectEmails            *bool            `json:"collectEmails"`
	TimeIncrement            *int             `json:"timeIncrement"`
	ActiveSlots              []DateTime       `json:"activeSlots"`
	EventTimezone            *string          `json:"eventTimezone"`
	SlotGeneration           *SlotGeneration  `json:"slotGeneration"`
	TimedRecurrence          *TimedRecurrence `json:"timedRecurrence"`
	ScheduleVersion          int              `json:"-"`

	// Used for specific times for specific dates feature
	HasSpecificTimes *bool      `json:"hasSpecificTimes"`
	Times            []DateTime `json:"times"`

	Type EventType `json:"type"`

	// PostHog ID for the event creator
	CreatorPosthogId *string `json:"creatorPosthogId"`

	// Sign up form details
	IsSignUpForm    *bool                      `json:"isSignUpForm"`
	SignUpBlocks    *[]SignUpBlock             `json:"signUpBlocks"`
	SignUpResponses map[string]*SignUpResponse `json:"signUpResponses"`

	// Whether to start the event on Monday (as opposed to Sunday, used for DOW events)
	StartOnMonday *bool `json:"startOnMonday"`

	// Whether to enable blind availability
	BlindAvailabilityEnabled *bool `json:"blindAvailabilityEnabled"`

	// Whether to only poll for days, not times
	DaysOnly *bool `json:"daysOnly"`

	// Availability responses
	ResponsesMap map[string]*Response `json:"responses"`

	// Used to store the number of responses for the event
	NumResponses *int `json:"numResponses"`

	// Scheduled event
	ScheduledEvent  *CalendarEvent `json:"scheduledEvent"`
	CalendarEventId string         `json:"calendarEventId"`

	// Remindees
	Remindees *[]Remindee `json:"remindees"`

	// Whether the current viewer has responded to the availability group
	HasResponded *bool `json:"hasResponded"`
}

// MarshalJSON keeps legacy schedule columns out of the timed-event API. They
// remain in this persistence type only while stored payloads are unmarshaled.
// It also always emits the Mongo-era attendees key as null so non-group payloads
// keep their stored shape; group reads replace it at the route boundary.
func (event Event) MarshalJSON() ([]byte, error) {
	type eventJSON Event
	if event.DaysOnly != nil && *event.DaysOnly {
		return json.Marshal(struct {
			*eventJSON
			Attendees any `json:"attendees"`
		}{eventJSON: (*eventJSON)(&event)})
	}

	return json.Marshal(struct {
		*eventJSON
		Duration         *float32   `json:"duration,omitempty"`
		Dates            []DateTime `json:"dates,omitempty"`
		TimeIncrement    *int       `json:"timeIncrement,omitempty"`
		HasSpecificTimes *bool      `json:"hasSpecificTimes,omitempty"`
		Times            []DateTime `json:"times,omitempty"`
		StartOnMonday    *bool      `json:"startOnMonday,omitempty"`
		Attendees        any        `json:"attendees"`
	}{eventJSON: (*eventJSON)(&event)})
}
