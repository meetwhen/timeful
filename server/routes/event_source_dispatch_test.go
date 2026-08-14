package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"timeful/server/errs"
)

func TestEventSourceHandlerBypassesMongoForPostgreSQLIDs(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	calledMongoHandler := false
	router.GET("/:eventId", eventSourceHandler(func(c *gin.Context) {
		calledMongoHandler = true
		c.Status(http.StatusOK)
	}, postgresEventRouteUnavailable))

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/p_01J3NYJ4ABCD1234EFGH5678JK", nil)
	router.ServeHTTP(recorder, request)

	if calledMongoHandler {
		t.Fatal("PostgreSQL ID reached the Mongo handler")
	}
	if recorder.Code != http.StatusNotImplemented {
		t.Fatalf("expected status %d, got %d: %s", http.StatusNotImplemented, recorder.Code, recorder.Body.String())
	}
	if recorder.Body.String() != `{"error":"`+errs.PostgreSQLEventUnsupported+`"}` {
		t.Fatalf("unexpected response body: %s", recorder.Body.String())
	}
}

func TestEventSourceHandlerKeepsMongoIDsOnMongoHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	calledMongoHandler := false
	router.GET("/:eventId", eventSourceHandler(func(c *gin.Context) {
		calledMongoHandler = true
		c.Status(http.StatusNoContent)
	}, postgresEventRouteUnavailable))

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/64f5e4d3c2b1a09876543210", nil)
	router.ServeHTTP(recorder, request)

	if !calledMongoHandler {
		t.Fatal("Mongo ID did not reach the Mongo handler")
	}
	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, recorder.Code)
	}
}
