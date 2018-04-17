package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type googleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func New(ctx context.Context, cfg config.Config, store database.DAL) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})
	r.HandleFunc("/auth/google", HandleGoogleAuth(ctx, cfg.Google.OAuth))
	r.HandleFunc("/auth/google/callback", HandleGoogleCallback(
		ctx,
		cfg.Google.OAuth,
		cfg.Google.UserInfo,
		store,
	))

	return r
}

// HandleGoogleAuth handles the Google Authentication route
func HandleGoogleAuth(ctx context.Context, cfg oauth2.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: generate random state and save in session
		state := "state"
		url := cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

// HandleGoogleCallback handles the google callback token exchange
func HandleGoogleCallback(
	ctx context.Context,
	cfg oauth2.Config,
	userInfo url.URL,
	store database.DAL,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: validate state query param matches stored session state value
		code := r.FormValue("code")

		if code == "" {
			http.Error(w, "code mismatch error", http.StatusInternalServerError)
			return
		}

		token, err := cfg.Exchange(oauth2.NoContext, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		v := url.Values{}
		v.Set("access_token", token.AccessToken)

		userInfo.RawQuery = v.Encode()

		response, err := http.Get(userInfo.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var info googleUserInfo
		if err := json.Unmarshal(data, &info); err != nil {
			log.Error(err)
			http.Error(w, "unmarshal error", http.StatusInternalServerError)
			return
		}

		user := models.User{
			Name:       info.Name,
			Email:      info.Email,
			AvatarUrl:  info.Picture,
			Provider:   "google",
			ProviderID: info.ID,
		}

		err = store.FindOrCreateUser(ctx, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userJson, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Content: %s\n", userJson)
	}
}
