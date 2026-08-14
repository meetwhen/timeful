package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"timeful/server/errs"
	"timeful/server/responses"
)

// postgresEventRouteUnavailable is the PostgreSQL dispatch branch until the
// compatibility route handlers are implemented in this package.
func postgresEventRouteUnavailable(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, responses.Error{Error: errs.PostgreSQLEventUnavailable})
}
