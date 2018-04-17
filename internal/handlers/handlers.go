package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

// HandleGoogleAuth handles the Google Authentication route
func HandleGoogleAuth(cfg config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: generate random state and save in session
		state := "state"
		url := cfg.Google.OAuth.AuthCodeURL(state, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

//HandleGoogleCallback handles the google callback token exchange
func HandleGoogleCallback(cfg config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: validate state query param matches stored session state value
		code := r.FormValue("code")

		if code == "" {
			fmt.Printf("Excepted code to recieve value, return %s ", code)
		}

		token, err := cfg.Google.OAuth.Exchange(oauth2.NoContext, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		c := cfg.Google.UserInfo

		v := url.Values{}
		v.Set("access_token", token.AccessToken)

		c.RawQuery = v.Encode()

		response, err := http.Get(c.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Content: %s\n", contents)
	}
}

// New handler instance
func New(cfg config.Config) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})
	r.HandleFunc("/auth/google", HandleGoogleAuth(cfg))
	r.HandleFunc("/auth/google/callback", HandleGoogleCallback(cfg))

	return r
}
