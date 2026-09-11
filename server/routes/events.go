/* The /events group contains all the routes to get and edit events */
package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"timeful/server/middleware"
	"timeful/server/models"
)

func rejectLegacyTimedScheduleFields(c *gin.Context) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))

	var payload map[string]json.RawMessage
	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}
	daysOnly := false
	if rawDaysOnly, exists := payload["daysOnly"]; exists {
		if err := json.Unmarshal(rawDaysOnly, &daysOnly); err != nil {
			return err
		}
	}
	legacyFields := []string{"duration", "times", "timeIncrement", "hasSpecificTimes", "startOnMonday"}
	if !daysOnly {
		legacyFields = append(legacyFields, "dates")
	}
	for _, field := range legacyFields {
		if _, exists := payload[field]; exists {
			return fmt.Errorf("legacy-timed-event-field:%s", field)
		}
	}
	return nil
}

func InitEvents(router *gin.RouterGroup) {
	eventRouter := router.Group("/events")

	eventRouter.POST("", postgresCreateEvent)
	eventRouter.POST("/:eventId/transfers", postgresCreateTransfer)
	eventRouter.POST("/:eventId/transfers/:transferId/:action", postgresTransferAction)
	eventRouter.POST("/:eventId/grant-association", postgresGrantAssociation)
	eventRouter.PUT("/:eventId", postgresEditEvent)
	eventRouter.GET("/:eventId/ids", postgresGetEventIDs)
	eventRouter.GET("/:eventId", postgresGetEvent)
	eventRouter.GET("/:eventId/responses", postgresGetResponses)
	eventRouter.POST("/:eventId/response", postgresUpdateResponse)
	eventRouter.DELETE("/:eventId/response", postgresDeleteResponse)
	eventRouter.PUT("/:eventId/schedule", postgresSaveSchedule)
	eventRouter.DELETE("/:eventId/schedule", postgresClearSchedule)
	eventRouter.POST("/:eventId/rename-user", postgresRenameUser)
	eventRouter.POST("/:eventId/decline", middleware.AuthRequired(), postgresDeclineInvite)
	eventRouter.GET("/:eventId/calendar-availabilities", middleware.AuthRequired(), postgresGetCalendarAvailabilities)
	eventRouter.DELETE("/:eventId", postgresDeleteEvent)
	eventRouter.POST("/:eventId/archive", postgresArchiveEvent)
}

func normalizeTimedResponseAvailabilitySlots(
	availability []models.DateTime,
	ifNeeded []models.DateTime,
) ([]models.DateTime, []models.DateTime) {
	normalizedAvailability := make([]models.DateTime, 0, len(availability))
	availabilitySet := make(map[models.DateTime]struct{}, len(availability))
	for _, slot := range availability {
		if _, exists := availabilitySet[slot]; exists {
			continue
		}
		availabilitySet[slot] = struct{}{}
		normalizedAvailability = append(normalizedAvailability, slot)
	}

	normalizedIfNeeded := make([]models.DateTime, 0, len(ifNeeded))
	ifNeededSet := make(map[models.DateTime]struct{}, len(ifNeeded))
	for _, slot := range ifNeeded {
		if _, exists := availabilitySet[slot]; exists {
			continue
		}
		if _, exists := ifNeededSet[slot]; exists {
			continue
		}
		ifNeededSet[slot] = struct{}{}
		normalizedIfNeeded = append(normalizedIfNeeded, slot)
	}

	return normalizedAvailability, normalizedIfNeeded
}

// stripSensitiveUserFields removes fields from a User that should never be
// exposed in the event page API response (calendar accounts, etc.).
// Email is NOT stripped here as callers handle email visibility separately based
// on the collectEmails setting and owner status.
func stripSensitiveUserFields(user *models.User) {
	if user == nil {
		return
	}
	user.CalendarAccounts = nil
	user.CalendarOptions = nil
	user.PrimaryAccountKey = nil
}
