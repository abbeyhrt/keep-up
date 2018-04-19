package handlers_test

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/handlers"
	"golang.org/x/oauth2"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type fixture struct {
	ctx     context.Context
	auth    *httptest.Server
	token   *httptest.Server
	user    *httptest.Server
	oauth   oauth2.Config
	store   database.DAL
	sqlmock struct {
		db   *sql.DB
		mock sqlmock.Sqlmock
	}
}

func newFixture(t *testing.T) *fixture {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	store := database.NewStoreFromClient(db)
	auth := getMockAuthService()
	token := getMockTokenService(false)
	user := getMockUserService(false)

	return &fixture{
		ctx:   context.Background(),
		auth:  auth,
		token: token,
		user:  user,
		store: store,
		oauth: oauth2.Config{
			ClientID:     "client_id",
			ClientSecret: "client_secret",
			RedirectURL:  "http://localhost:3000/auth/google/callback",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: oauth2.Endpoint{
				AuthURL:  auth.URL,
				TokenURL: token.URL,
			},
		},
		sqlmock: struct {
			db   *sql.DB
			mock sqlmock.Sqlmock
		}{
			db,
			mock,
		},
	}
}

func (f *fixture) teardown() {
	f.auth.Close()
	f.token.Close()
	f.user.Close()
	f.sqlmock.db.Close()
}

func TestHandleGoogleAuth(t *testing.T) {
	f := newFixture(t)
	defer f.teardown()

	handler := handlers.HandleGoogleAuth(f.ctx, f.oauth)
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

	u, _ := url.Parse(f.oauth.Endpoint.AuthURL)
	v := url.Values{}

	v.Set("access_type", "offline")
	v.Set("client_id", "client_id")
	v.Set("redirect_uri", f.oauth.RedirectURL)
	v.Set("response_type", "code")
	v.Set("scope", "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")
	v.Set("state", "state")

	u.RawQuery = v.Encode()

	if u.String() != location[0] {
		t.Fatalf("expected location header to be: %s, instead recieved: %s", u.String(), location[0])
	}
}

func it(should string, t *testing.T, callback func(t *testing.T, f *fixture)) {
	t.Run("it "+should, func(t *testing.T) {
		f := newFixture(t)
		callback(t, f)
		f.teardown()
	})
}

func TestHandleGoogleCallback(t *testing.T) {
	it("should throw if no code provided", t, func(t *testing.T, f *fixture) {
		userInfoURL, _ := url.Parse(f.user.URL)
		handler := handlers.HandleGoogleCallback(f.ctx, f.oauth, *userInfoURL, f.store)
		callbackURL, _ := url.Parse("http://localhost:3000/auth/github/callback")
		req := httptest.NewRequest("GET", callbackURL.String(), nil)
		w := httptest.NewRecorder()

		handler(w, req)

		result := w.Result()

		if result.StatusCode != 500 {
			t.Fatalf("expected result to have status code: %d, instead recieved: %d", 500, result.StatusCode)
		}
	})

	it(
		"should fail if no access_token received in exchange",
		t,
		func(t *testing.T, f *fixture) {
			// Setup service to fail
			tokenStub := getMockTokenService(true)
			f.token = tokenStub
			f.oauth.Endpoint.TokenURL = tokenStub.URL

			userInfoURL, _ := url.Parse(f.user.URL)
			handler := handlers.HandleGoogleCallback(
				f.ctx,
				f.oauth,
				*userInfoURL,
				f.store,
			)
			callbackURL, _ := url.Parse("http://localhost:3000/auth/github/callback")

			v := url.Values{}
			v.Set("code", "code")
			v.Set("state", "state")

			callbackURL.RawQuery = v.Encode()

			req := httptest.NewRequest("GET", callbackURL.String(), nil)
			w := httptest.NewRecorder()

			handler(w, req)

			result := w.Result()

			if result.StatusCode != 500 {
				t.Fatalf("expected result to have status code: %d, instead recieved: %d", 500, result.StatusCode)
			}
		},
	)
}

func getMockAuthService() *httptest.Server {
	authStub := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("auth"))
			},
		),
	)
	return authStub
}

func getMockTokenService(shouldFail bool) *httptest.Server {
	tokenStub := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if shouldFail {
					http.Error(w, "should fail", http.StatusInternalServerError)
					return
				}

				if r.FormValue("code") != "code" {
					http.Error(w, "no code recieved", http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte("access_token=access_token"))
			},
		),
	)
	return tokenStub
}

func getMockUserService(shouldFail bool) *httptest.Server {
	userStub := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("uh")
				fmt.Println(shouldFail)
				if shouldFail {
					fmt.Println("totally failing")
					http.Error(w, "should fail", http.StatusInternalServerError)
					return
				}
				w.Write([]byte("user"))
			},
		),
	)
	return userStub
}
