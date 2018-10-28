package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abbeyhrt/keep-up/graphql/internal/database"
	"github.com/abbeyhrt/keep-up/graphql/internal/handlers"
)

// Cookie needs to be encoded, and properly decoded
// Set the session cookie for the current user
// Need to get the session by the id
// We need to redirect if there is no session key
// We need to remember the path if we redirect because there is no session key

func TestSessionRedirectPath(t *testing.T) {
	mockCtx := context.Background()
	mockStore := database.NewMockStore()
	mockCookieSecret := "abcd"
	m := handlers.NewMiddleware(
		mockCtx,
		mockStore,
		mockCookieSecret,
	)
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	ts := httptest.NewServer(m.Session(next))
	defer ts.Close()
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusSeeOther {
		t.Fatalf(
			"Expected status code to be: %d, instead received: %d",
			http.StatusSeeOther,
			res.StatusCode,
		)
	}

	location := res.Header.Get("Location")
	if location != handlers.RedirectPath {
		t.Fatalf(
			"Expected location to be: %s, instead received: %s",
			handlers.RedirectPath,
			location,
		)
	}
}
