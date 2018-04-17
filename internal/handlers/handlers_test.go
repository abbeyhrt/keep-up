package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abbeyhrt/keep-up-graphql/internal/handlers"
	"golang.org/x/oauth2"
)

func TestHandleGoogleAuth(t *testing.T) {
	oauth := oauth2.Config{
		ClientID:     "client_id",
		ClientSecret: "client_secret",
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
	handler := handlers.HandleGoogleAuth(oauth)
	req := httptest.NewRequest("GET", "http://localhost:3000/auth/github", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusFound {
		t.Fatalf("expected status code to be: %d, instead received: %d", http.StatusFound, resp.StatusCode)
	}

	location, ok := resp.Header["Location"]
	if !ok {
		t.Fatalf("expected a location header to exist in the response")
	}

	u, _ := url.Parse("https://accounts.google.com/o/oauth2/auth")
	v := url.Values{}

	v.Set("access_type", "offline")
	v.Set("client_id", "client_id")
	v.Set("redirect_uri", "http://localhost:3000/auth/google/callback")
	v.Set("response_type", "code")
	v.Set("scope", "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")
	v.Set("state", "state")

	u.RawQuery = v.Encode()

	if u.String() != location[0] {
		t.Fatalf("expected location header to be: %s, instead recieved: %s", u.String(), location[0])
	}
}
