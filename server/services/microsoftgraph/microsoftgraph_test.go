package microsoftgraph

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"timeful/server/models"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func TestGetUserInfo(t *testing.T) {
	previousTransport := http.DefaultTransport
	http.DefaultTransport = roundTripperFunc(func(request *http.Request) (*http.Response, error) {
		if request.Method != http.MethodGet {
			t.Errorf("method = %q, want %q", request.Method, http.MethodGet)
		}
		if request.URL.String() != "https://graph.microsoft.com/v1.0/me?$select=givenName,surname,mail" {
			t.Errorf("URL = %q", request.URL.String())
		}
		if request.Header.Get("Authorization") != "Bearer test-access-token" {
			t.Errorf("Authorization = %q", request.Header.Get("Authorization"))
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"givenName":"Ada","surname":"Lovelace","mail":"ada@example.com"}`)),
			Header:     make(http.Header),
			Request:    request,
		}, nil
	})
	t.Cleanup(func() { http.DefaultTransport = previousTransport })

	userInfo := GetUserInfo(nil, &models.OAuth2CalendarAuth{AccessToken: "test-access-token"})
	if userInfo != (UserInfo{FirstName: "Ada", LastName: "Lovelace", Email: "ada@example.com"}) {
		t.Fatalf("user info = %#v", userInfo)
	}
}
