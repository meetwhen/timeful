package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"timeful/server/errs"
	"timeful/server/models"
	pgstore "timeful/server/postgres"
	"timeful/server/respondents"
	"timeful/server/responses"
	"timeful/server/utils"
)

func postgresRepository(c *gin.Context) *pgstore.Repository {
	repository, err := pgstore.DefaultRepository()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, responses.Error{Error: "postgres-event-store-unavailable"})
		return nil
	}
	return repository
}

func postgresEvent(c *gin.Context, repository *pgstore.Repository) *pgstore.Event {
	event, err := repository.GetEventByShortID(c.Request.Context(), c.Param("eventId"))
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusNotFound, responses.Error{Error: errs.EventNotFound})
		return nil
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-load-event"})
		return nil
	}
	return event
}

func postgresEventModel(event *pgstore.Event) (models.Event, error) {
	var value models.Event
	if err := json.Unmarshal(event.Payload, &value); err != nil {
		return value, err
	}
	value.Id = primitive.NilObjectID
	value.ShortId = &event.ShortID
	value.OwnerId = primitive.NilObjectID
	value.Name = event.Name
	value.Type = models.EventType(event.Type)
	value.ScheduleVersion = event.ScheduleVersion
	value.NumResponses = &event.NumResponses
	value.CreatorPosthogId = event.CreatorPosthogID
	return value, nil
}

func postgresResponseModel(stored pgstore.Response) (*models.Response, string, error) {
	var value models.Response
	if err := json.Unmarshal(stored.Payload, &value); err != nil {
		return nil, "", err
	}
	if stored.RespondentKind == pgstore.RespondentKindAccount {
		value.UserId = utils.StringToObjectID(*stored.AccountUserID)
		key, ok := populateResponsePayloadIdentity(&value, *stored.AccountUserID)
		if !ok {
			return nil, "", errors.New("invalid account response")
		}
		return &value, key, nil
	}
	value.UserId = primitive.NilObjectID
	value.GuestId = dereference(stored.GuestID)
	value.GuestEditToken = dereference(stored.GuestEditToken)
	value.GuestEditPolicy = dereference(stored.GuestEditPolicy)
	value.GuestOwnershipMode = dereference(stored.GuestOwnershipMode)
	value.Name = dereference(stored.CanonicalGuestName)
	normalizeGuestResponseForPayload(&value)
	return &value, guestResponseLookupKey(models.EventResponse{Response: &value}), nil
}

func dereference(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func postgresResponses(ctx context.Context, repository *pgstore.Repository, event *pgstore.Event) (map[string]*models.Response, error) {
	stored, err := repository.ListResponses(ctx, event.ID)
	if err != nil {
		return nil, err
	}
	result := make(map[string]*models.Response, len(stored))
	for _, response := range stored {
		value, key, err := postgresResponseModel(response)
		if err != nil {
			return nil, err
		}
		result[key] = value
	}
	return result, nil
}

func postgresEventPayload(event *pgstore.Event, responseMap map[string]*models.Response) (map[string]any, error) {
	value, err := postgresEventModel(event)
	if err != nil {
		return nil, err
	}
	value.ResponsesMap = responseMap
	payload, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		return nil, err
	}
	result["_id"] = event.ShortID
	result["shortId"] = event.ShortID
	result["ownerId"] = primitive.NilObjectID.Hex()
	return result, nil
}

func postgresGetEventIDs(c *gin.Context) {
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	event := postgresEvent(c, repository)
	if event == nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"shortId": event.ShortID, "longId": event.ShortID})
}

func postgresGetEvent(c *gin.Context) {
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	event := postgresEvent(c, repository)
	if event == nil {
		return
	}
	responseMap, err := postgresResponses(c.Request.Context(), repository, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-load-responses"})
		return
	}
	for key, response := range responseMap {
		stripSensitiveUserFields(response.User)
		response.Email = ""
		if response.User != nil {
			response.User.Email = ""
		}
		response.Availability = nil
		response.IfNeeded = nil
		response.ManualAvailability = nil
		responseMap[key] = response
	}
	value, err := postgresEventModel(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-serialize-event"})
		return
	}
	blind := utils.Coalesce(value.BlindAvailabilityEnabled)
	if blind {
		key := guestQueryLookupKey(c.Query("guestId"), c.Query("guestName"))
		if sessionID, ok := sessions.Default(c).Get("userId").(string); ok {
			key = sessionID
		}
		if key == "" {
			responseMap = nil
		} else if response, exists := responseMap[key]; exists {
			responseMap = map[string]*models.Response{key: response}
		} else {
			responseMap = map[string]*models.Response{}
		}
	}
	payload, err := postgresEventPayload(event, responseMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-serialize-event"})
		return
	}
	if blind {
		delete(payload, "numResponses")
		if responseMap == nil {
			delete(payload, "responses")
		}
	}
	c.JSON(http.StatusOK, payload)
}

func postgresGetResponses(c *gin.Context) {
	query := struct {
		TimeMin time.Time `form:"timeMin" binding:"required"`
		TimeMax time.Time `form:"timeMax" binding:"required"`
	}{}
	if err := c.BindQuery(&query); err != nil {
		return
	}
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	event := postgresEvent(c, repository)
	if event == nil {
		return
	}
	responseMap, err := postgresResponses(c.Request.Context(), repository, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-load-responses"})
		return
	}
	for key, response := range responseMap {
		response.Availability = filterResponseSlots(response.Availability, query.TimeMin, query.TimeMax)
		response.IfNeeded = filterResponseSlots(response.IfNeeded, query.TimeMin, query.TimeMax)
		stripSensitiveUserFields(response.User)
		response.Email = ""
		if response.User != nil {
			response.User.Email = ""
		}
		responseMap[key] = response
	}
	value, err := postgresEventModel(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-serialize-event"})
		return
	}
	if utils.Coalesce(value.BlindAvailabilityEnabled) {
		key := guestQueryLookupKey(c.Query("guestId"), c.Query("guestName"))
		if sessionID, ok := sessions.Default(c).Get("userId").(string); ok {
			key = sessionID
		}
		if key == "" {
			c.JSON(http.StatusOK, map[string]*models.Response{})
			return
		}
		if response, exists := responseMap[key]; exists {
			c.JSON(http.StatusOK, map[string]*models.Response{key: response})
			return
		}
		c.JSON(http.StatusOK, map[string]*models.Response{})
		return
	}
	c.JSON(http.StatusOK, responseMap)
}

func filterResponseSlots(slots []primitive.DateTime, minimum, maximum time.Time) []primitive.DateTime {
	filtered := make([]primitive.DateTime, 0, len(slots))
	for _, slot := range slots {
		if !slot.Time().Before(minimum) && !slot.Time().After(maximum) {
			filtered = append(filtered, slot)
		}
	}
	return filtered
}

func postgresEditEvent(c *gin.Context) {
	if err := rejectLegacyTimedScheduleFields(c); err != nil {
		c.JSON(http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var update models.Event
	if err := c.Bind(&update); err != nil {
		return
	}
	if update.Name == "" || update.Type == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if update.DaysOnly == nil || !*update.DaysOnly {
		fields, err := normalizeTimedEventPayloadFields(timedEventPayloadFields{ActiveSlots: update.ActiveSlots, EventTimezone: update.EventTimezone, SlotGeneration: update.SlotGeneration, TimedRecurrence: update.TimedRecurrence})
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Error{Error: err.Error()})
			return
		}
		update.ActiveSlots, update.EventTimezone, update.SlotGeneration, update.TimedRecurrence = fields.ActiveSlots, fields.EventTimezone, fields.SlotGeneration, fields.TimedRecurrence
	} else if len(update.Dates) == 0 {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "days-only-events-require-dates"})
		return
	}
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	event := postgresEvent(c, repository)
	if event == nil {
		return
	}
	current, err := postgresEventModel(event)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if _, present := raw["description"]; !present {
		update.Description = current.Description
	}
	if slots, present := raw["activeSlots"]; present && string(slots) == "[]" && len(current.ActiveSlots) > 0 {
		update.ActiveSlots = current.ActiveSlots
	}
	update.Id, update.ShortId, update.OwnerId, update.NumResponses, update.ResponsesMap = primitive.NilObjectID, nil, primitive.NilObjectID, nil, nil
	payload, err := json.Marshal(update)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	event.Name, event.Type, event.Payload, event.ScheduleVersion = update.Name, string(update.Type), payload, 1
	if err := repository.UpdateEvent(c.Request.Context(), event); err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-update-event"})
		return
	}
	c.Status(http.StatusOK)
}

func postgresSaveSchedule(c *gin.Context)  { postgresUpdateSchedule(c, false) }
func postgresClearSchedule(c *gin.Context) { postgresUpdateSchedule(c, true) }

func postgresUpdateSchedule(c *gin.Context, clear bool) {
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	event := postgresEvent(c, repository)
	if event == nil {
		return
	}
	value, err := postgresEventModel(event)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if clear {
		value.ScheduledEvent = nil
	} else {
		payload := struct {
			StartDate primitive.DateTime `json:"startDate" binding:"required"`
			EndDate   primitive.DateTime `json:"endDate" binding:"required"`
		}{}
		if err := c.Bind(&payload); err != nil {
			return
		}
		if payload.EndDate <= payload.StartDate {
			c.JSON(http.StatusBadRequest, responses.Error{Error: "scheduled-event-end-must-follow-start"})
			return
		}
		value.ScheduledEvent = &models.CalendarEvent{Summary: event.Name, StartDate: payload.StartDate, EndDate: payload.EndDate}
	}
	value.Id, value.ShortId, value.OwnerId, value.NumResponses, value.ResponsesMap = primitive.NilObjectID, nil, primitive.NilObjectID, nil, nil
	event.Payload, err = json.Marshal(value)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if err := repository.UpdateEvent(c.Request.Context(), event); err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-save-scheduled-event"})
		return
	}
	c.Status(http.StatusOK)
}

func postgresUpdateResponse(c *gin.Context) {
	payload := struct {
		Availability    []primitive.DateTime `json:"availability"`
		IfNeeded        []primitive.DateTime `json:"ifNeeded"`
		Guest           *bool                `json:"guest" binding:"required"`
		Name            string               `json:"name"`
		Email           string               `json:"email"`
		GuestId         string               `json:"guestId"`
		GuestEditToken  string               `json:"guestEditToken"`
		GuestEditPolicy *string              `json:"guestEditPolicy"`
	}{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	event := postgresEvent(c, repository)
	if event == nil {
		return
	}
	availability, ifNeeded := normalizeTimedResponseAvailabilitySlots(payload.Availability, payload.IfNeeded)
	result := guestResponseMutationResult{}
	err := repository.WithTransaction(c.Request.Context(), func(ctx context.Context, tx *pgstore.Repository) error {
		if *payload.Guest {
			return postgresMutateGuestResponse(ctx, tx, event, payload, availability, ifNeeded, &result, c)
		}
		userID, ok := sessions.Default(c).Get("userId").(string)
		if !ok {
			return errNotSignedIn
		}
		existing, err := tx.GetResponseByAccountUserID(ctx, event.ID, userID)
		response := models.Response{UserId: utils.StringToObjectID(userID), Availability: availability, IfNeeded: ifNeeded}
		encoded, _ := json.Marshal(response)
		if errors.Is(err, pgx.ErrNoRows) {
			id := userID
			if err := tx.CreateResponse(ctx, &pgstore.Response{EventID: event.ID, RespondentKind: pgstore.RespondentKindAccount, AccountUserID: &id, Payload: encoded}); err != nil {
				return err
			}
			event.NumResponses++
			return tx.UpdateEvent(ctx, event)
		}
		if err != nil {
			return err
		}
		existing.Payload = encoded
		return tx.UpdateResponse(ctx, existing)
	})
	if errors.Is(err, errNotSignedIn) {
		c.JSON(http.StatusUnauthorized, responses.Error{Error: errs.NotSignedIn})
		return
	}
	if err != nil {
		postgresMutationError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

var errNotSignedIn = errors.New("not signed in")

func postgresMutateGuestResponse(ctx context.Context, tx *pgstore.Repository, event *pgstore.Event, input struct {
	Availability    []primitive.DateTime `json:"availability"`
	IfNeeded        []primitive.DateTime `json:"ifNeeded"`
	Guest           *bool                `json:"guest" binding:"required"`
	Name            string               `json:"name"`
	Email           string               `json:"email"`
	GuestId         string               `json:"guestId"`
	GuestEditToken  string               `json:"guestEditToken"`
	GuestEditPolicy *string              `json:"guestEditPolicy"`
}, availability, ifNeeded []primitive.DateTime, result *guestResponseMutationResult, c *gin.Context) error {
	validated := respondents.ValidateGuestName(input.Name)
	if validated.Code != respondents.GuestNameValid {
		return guestNameError{guestNameValidationErrorMessage(validated.Code)}
	}
	existing, err := tx.GetResponseByGuestID(ctx, event.ID, input.GuestId)
	if errors.Is(err, pgx.ErrNoRows) {
		existing, err = tx.GetResponseByGuestName(ctx, event.ID, validated.Name)
	}
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	response := models.Response{Name: validated.Name, Email: input.Email, Availability: availability, IfNeeded: ifNeeded}
	if existing != nil {
		stored, _, err := postgresResponseModel(*existing)
		if err != nil {
			return err
		}
		if !canMutateGuestResponse(stored, guestQueryLookupKey(c.Query("guestId"), c.Query("guestName")), input.GuestEditToken) {
			return guestForbidden{"This guest response is protected and cannot be edited without its edit token"}
		}
		response.GuestId, response.GuestEditToken, response.GuestOwnershipMode = stored.GuestId, stored.GuestEditToken, stored.GuestOwnershipMode
		response.GuestEditPolicy = stored.GuestEditPolicy
		if input.GuestEditPolicy != nil {
			response.GuestEditPolicy = normalizeGuestEditPolicy(*input.GuestEditPolicy)
		}
		result.GuestCredentials = buildGuestCredentialsResponse(&response, validated.Name)
		encoded, _ := json.Marshal(response)
		existing.CanonicalGuestName = &validated.Name
		existing.GuestEditPolicy = &response.GuestEditPolicy
		existing.Payload = encoded
		return tx.UpdateResponse(ctx, existing)
	}
	policy := guestEditPolicyProtected
	if input.GuestEditPolicy != nil {
		policy = *input.GuestEditPolicy
	}
	result.GuestCredentials = ensureGuestTokenOwnership(&response, policy)
	encoded, _ := json.Marshal(response)
	guestID, token, ownership, editPolicy := response.GuestId, response.GuestEditToken, response.GuestOwnershipMode, response.GuestEditPolicy
	if err := tx.CreateResponse(ctx, &pgstore.Response{EventID: event.ID, RespondentKind: pgstore.RespondentKindGuest, GuestID: &guestID, CanonicalGuestName: &validated.Name, GuestEditPolicy: &editPolicy, GuestOwnershipMode: &ownership, GuestEditToken: &token, Payload: encoded}); err != nil {
		return err
	}
	event.NumResponses++
	return tx.UpdateEvent(ctx, event)
}

type guestNameError struct{ message string }

func (e guestNameError) Error() string { return e.message }

type guestForbidden struct{ message string }

func (e guestForbidden) Error() string { return e.message }

func postgresDeleteResponse(c *gin.Context) {
	payload := struct {
		UserId         string `json:"userId"`
		Guest          *bool  `json:"guest" binding:"required"`
		Name           string `json:"name"`
		GuestId        string `json:"guestId"`
		GuestEditToken string `json:"guestEditToken"`
	}{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	event := postgresEvent(c, repository)
	if event == nil {
		return
	}
	if !*payload.Guest {
		userID, ok := sessions.Default(c).Get("userId").(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, responses.Error{Error: errs.NotSignedIn})
			return
		}
		if payload.UserId != userID {
			c.JSON(http.StatusForbidden, responses.Error{Error: errs.UserNotEventOwner})
			return
		}
		if err := postgresDeleteStoredResponse(c.Request.Context(), repository, event, func(tx *pgstore.Repository) (*pgstore.Response, error) {
			return tx.GetResponseByAccountUserID(c.Request.Context(), event.ID, userID)
		}); err != nil {
			postgresMutationError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	canonicalName := canonicalGuestName(payload.Name)
	if payload.GuestId == "" && canonicalName == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Guest name is required"})
		return
	}
	err := repository.WithTransaction(c.Request.Context(), func(ctx context.Context, tx *pgstore.Repository) error {
		var stored *pgstore.Response
		var err error
		if payload.GuestId != "" {
			stored, err = tx.GetResponseByGuestID(ctx, event.ID, payload.GuestId)
		} else {
			stored, err = tx.GetResponseByGuestName(ctx, event.ID, canonicalName)
		}
		if err != nil {
			return err
		}
		value, _, err := postgresResponseModel(*stored)
		if err != nil {
			return err
		}
		if !canMutateGuestResponse(value, guestQueryLookupKey(c.Query("guestId"), c.Query("guestName")), payload.GuestEditToken) {
			return guestForbidden{"This guest response is protected and cannot be deleted without its edit token"}
		}
		if err := tx.DeleteResponse(ctx, stored.ID); err != nil {
			return err
		}
		if event.NumResponses > 0 {
			event.NumResponses--
		}
		return tx.UpdateEvent(ctx, event)
	})
	if err != nil {
		postgresMutationError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func postgresDeleteStoredResponse(ctx context.Context, repository *pgstore.Repository, event *pgstore.Event, lookup func(*pgstore.Repository) (*pgstore.Response, error)) error {
	return repository.WithTransaction(ctx, func(txCtx context.Context, tx *pgstore.Repository) error {
		stored, err := lookup(tx)
		if err != nil {
			return err
		}
		if err := tx.DeleteResponse(txCtx, stored.ID); err != nil {
			return err
		}
		if event.NumResponses > 0 {
			event.NumResponses--
		}
		return tx.UpdateEvent(txCtx, event)
	})
}

func postgresRenameUser(c *gin.Context) {
	payload := struct {
		OldName        string `json:"oldName"`
		NewName        string `json:"newName"`
		GuestId        string `json:"guestId"`
		GuestEditToken string `json:"guestEditToken"`
	}{}
	if err := c.Bind(&payload); err != nil {
		return
	}
	oldName := canonicalGuestName(payload.OldName)
	if payload.GuestId == "" && oldName == "" {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "Existing guest name is required"})
		return
	}
	validated := respondents.ValidateGuestName(payload.NewName)
	if validated.Code != respondents.GuestNameValid {
		c.JSON(http.StatusBadRequest, responses.Error{Error: guestNameValidationErrorMessage(validated.Code)})
		return
	}
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	event := postgresEvent(c, repository)
	if event == nil {
		return
	}
	result := guestResponseMutationResult{}
	err := repository.WithTransaction(c.Request.Context(), func(ctx context.Context, tx *pgstore.Repository) error {
		var stored *pgstore.Response
		var err error
		if payload.GuestId != "" {
			stored, err = tx.GetResponseByGuestID(ctx, event.ID, payload.GuestId)
		} else {
			stored, err = tx.GetResponseByGuestName(ctx, event.ID, oldName)
		}
		if err != nil {
			return err
		}
		value, _, err := postgresResponseModel(*stored)
		if err != nil {
			return err
		}
		if !canMutateGuestResponse(value, guestQueryLookupKey(c.Query("guestId"), c.Query("guestName")), payload.GuestEditToken) {
			return guestForbidden{"This guest response is protected and cannot be renamed without its edit token"}
		}
		value.Name = validated.Name
		if !isTokenBackedGuestResponse(value) {
			result.GuestCredentials = ensureGuestTokenOwnership(value, guestEditPolicyProtected)
		} else {
			result.GuestCredentials = buildGuestCredentialsResponse(value, validated.Name)
		}
		encoded, _ := json.Marshal(value)
		stored.CanonicalGuestName, stored.GuestID, stored.GuestEditToken, stored.GuestEditPolicy, stored.GuestOwnershipMode, stored.Payload = &validated.Name, &value.GuestId, &value.GuestEditToken, &value.GuestEditPolicy, &value.GuestOwnershipMode, encoded
		return tx.UpdateResponse(ctx, stored)
	})
	if err != nil {
		postgresMutationError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func postgresMutationError(c *gin.Context, err error) {
	if errors.Is(err, pgx.ErrNoRows) {
		c.JSON(http.StatusNotFound, responses.Error{Error: errs.EventNotFound})
		return
	}
	var forbidden guestForbidden
	if errors.As(err, &forbidden) {
		c.JSON(http.StatusForbidden, responses.Error{Error: forbidden.message})
		return
	}
	var nameError guestNameError
	if errors.As(err, &nameError) {
		c.JSON(http.StatusBadRequest, responses.Error{Error: nameError.message})
		return
	}
	if pgstore.IsUniqueViolation(err) {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "A guest with this name already exists for this event"})
		return
	}
	c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-update-response"})
}

func postgresEventRouteUnavailable(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, responses.Error{Error: errs.PostgreSQLEventUnsupported})
}

func postgresCreationEnabled(c *gin.Context) bool {
	if !strings.EqualFold(os.Getenv("POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED"), "true") {
		return false
	}
	if _, signedIn := sessions.Default(c).Get("userId").(string); signedIn {
		return false
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return false
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var payload struct {
		Type            models.EventType `json:"type"`
		DaysOnly        bool             `json:"daysOnly"`
		IsSignUpForm    bool             `json:"isSignUpForm"`
		ActiveSlots     json.RawMessage  `json:"activeSlots"`
		SlotGeneration  json.RawMessage  `json:"slotGeneration"`
		TimedRecurrence json.RawMessage  `json:"timedRecurrence"`
	}
	if json.Unmarshal(body, &payload) != nil || payload.IsSignUpForm || (payload.Type != models.SPECIFIC_DATES && payload.Type != models.DOW) {
		return false
	}
	return payload.DaysOnly || (len(payload.ActiveSlots) > 0 && len(payload.SlotGeneration) > 0 && len(payload.TimedRecurrence) > 0)
}

func postgresCreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.Bind(&event); err != nil {
		return
	}
	if event.Name == "" || (event.Type != models.SPECIFIC_DATES && event.Type != models.DOW) {
		c.Status(http.StatusBadRequest)
		return
	}
	if event.DaysOnly == nil || !*event.DaysOnly {
		fields, err := normalizeTimedEventPayloadFields(timedEventPayloadFields{ActiveSlots: event.ActiveSlots, EventTimezone: event.EventTimezone, SlotGeneration: event.SlotGeneration, TimedRecurrence: event.TimedRecurrence})
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Error{Error: err.Error()})
			return
		}
		event.ActiveSlots, event.EventTimezone, event.SlotGeneration, event.TimedRecurrence = fields.ActiveSlots, fields.EventTimezone, fields.SlotGeneration, fields.TimedRecurrence
	} else if len(event.Dates) == 0 {
		c.JSON(http.StatusBadRequest, responses.Error{Error: "days-only-events-require-dates"})
		return
	}
	event.Id, event.ShortId, event.OwnerId, event.NumResponses, event.ResponsesMap = primitive.NilObjectID, nil, primitive.NilObjectID, nil, nil
	encoded, err := json.Marshal(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-serialize-event"})
		return
	}
	repository := postgresRepository(c)
	if repository == nil {
		return
	}
	stored := &pgstore.Event{Name: event.Name, Type: string(event.Type), ScheduleVersion: 1, CreatorPosthogID: event.CreatorPosthogId, Payload: encoded}
	if err := repository.CreateEvent(c.Request.Context(), stored); err != nil {
		c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed-to-create-event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"eventId": stored.ShortID})
}
