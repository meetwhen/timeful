package routes

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"timeful/server/models"
	"timeful/server/respondents"
)

const guestOwnershipModeToken = "token"

func isTokenBackedGuestResponse(response *models.Response) bool {
	return response != nil &&
		response.GuestOwnershipMode == guestOwnershipModeToken &&
		response.GuestId != ""
}

func guestResponseLookupKey(eventResponse models.EventResponse) string {
	if isTokenBackedGuestResponse(eventResponse.Response) {
		return eventResponse.Response.GuestId
	}
	if eventResponse.Response != nil {
		if canonicalName := canonicalGuestName(eventResponse.Response.Name); canonicalName != "" {
			return canonicalName
		}
	}
	return eventResponse.UserId
}

func hasValidGuestName(name string) bool {
	return respondents.HasValidGuestName(name)
}

func shouldExposeGuestSignUpResponsePayload(_ string, response *models.SignUpResponse) bool {
	if response == nil {
		return true
	}

	if response.UserId != primitive.NilObjectID {
		return true
	}

	return hasValidGuestName(response.Name)
}
