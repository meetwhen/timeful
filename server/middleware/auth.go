package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"timeful/server/accounts"
	"timeful/server/errs"
	"timeful/server/responses"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// The session carries the account's external user identifier, which is
		// the hexadecimal legacy MongoDB account identifier.
		session := sessions.Default(c)
		externalUserID, ok := session.Get("userId").(string)
		if !ok || externalUserID == "" {
			c.JSON(http.StatusUnauthorized, responses.Error{Error: errs.NotSignedIn})
			c.Abort()
			return
		}

		// Resolve the authoritative PostgreSQL account. A session that predates
		// the cutover adopts its legacy profile exactly once.
		account, err := accounts.Resolve(c.Request.Context(), externalUserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.Error{Error: errs.UserDoesNotExist})
			c.Abort()
			return
		}

		// Calendar connections, provider tokens, sub-calendars, and preferences
		// are PostgreSQL-authoritative. They are loaded through the accounts
		// boundary, never from the retained MongoDB document.
		user, err := accounts.LoadSessionUser(c.Request.Context(), account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Error{Error: "failed to load account integration data"})
			c.Abort()
			return
		}

		c.Set("authUser", user)
		c.Set("authAccount", account)

		c.Next()
	}
}
