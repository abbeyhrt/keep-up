package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/abbeyhrt/keep-up/graphql/internal/config"
	"github.com/abbeyhrt/keep-up/graphql/internal/database"
	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	//"github.com/graphql-gophers/graphql-go"
	//	"github.com/graphql-gophers/graphql-go/relay"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	uuid "github.com/satori/go.uuid"
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
	r.Use(RequestIDMiddleware)
	r.Use(LoggingMiddleware)
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Keep Up</title>
			</head>
			<body>
			<h1>Keep Up</h1>
			<a href="/auth/google">Login with Google</a>
			</body>
			</html>
			`))
	})

	r.HandleFunc(
		"/auth/google",
		HandleGoogleAuth(ctx, cfg.Google.OAuth, cfg.CookieSecret),
	).Methods("GET")

	r.HandleFunc(
		"/auth/google/callback",
		HandleGoogleCallback(
			ctx,
			cfg.Google.OAuth,
			cfg.CookieSecret,
			cfg.Google.UserInfo,
			store,
		),
	).Methods("GET")

	r.Handle("/graphiql", GraphiqlHandler())

	s := r.PathPrefix("/").Subrouter()

	s.Use(SessionMiddleware(ctx, store, cfg.CookieSecret))
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
				<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Keep Up</title>
			</head>
			<body>
			<h1>Keep Up</h1>
			<a href="/auth/google">Login with Google</a>
			</body>
			</html>
		`))
	})
	s.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(sessionKey)
		if err != nil {
			log.Error()
			http.Error(w, "error finding cookie", http.StatusInternalServerError)
		}
		c.MaxAge = -1

		http.SetCookie(w, c)

		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	})

	return r
}

// GraphQLHandler will be the handler for all graphql queries
// func GraphQLHandler(db database.DAL) *relay.Handler {
// 	schema := graphql.MustParseSchema(schema.Schema, resolver.New(db))
// 	return &relay.Handler{Schema: schema}
// }

func GraphiqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(page)
	})
}

var graphiqlVersion = "0.11.2"
var page = []byte(fmt.Sprintf(`<!--
The request to this GraphQL server provided the header "Accept: text/html"
and as a result has been presented GraphiQL - an in-browser IDE for
exploring GraphQL.
If you wish to receive JSON, provide the header "Accept: application/json" or
add "&raw" to the end of the URL within a browser.
-->
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8" />
		<title>GraphiQL</title>
		<meta name="robots" content="noindex" />
		<style>
			html, body {
				height: 100%%;
				margin: 0;
				overflow: hidden;
				width: 100%%;
			}
		</style>
		<link href="//cdn.jsdelivr.net/npm/graphiql@%[1]s/graphiql.css" rel="stylesheet" />
		<script src="//cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
		<script src="//cdn.jsdelivr.net/react/15.4.2/react.min.js"></script>
		<script src="//cdn.jsdelivr.net/react/15.4.2/react-dom.min.js"></script>
		<script src="//cdn.jsdelivr.net/npm/graphiql@%[1]s/graphiql.min.js"></script>
	</head>
	<body>
		<div id="root" style="height: 100vh;">Loading...</div>
		<script>
			// Collect the URL parameters
			var parameters = {};
			window.location.search.substr(1).split('&').forEach(function (entry) {
			var eq = entry.indexOf('=');
			if (eq >= 0) {
				parameters[decodeURIComponent(entry.slice(0, eq))] =
				decodeURIComponent(entry.slice(eq + 1));
			}
			});
			// Produce a Location query string from a parameter object.
			function locationQuery(params) {
				return '?' + Object.keys(params).filter(function (key) {
					return Boolean(params[key]);
				}).map(function (key) {
					return encodeURIComponent(key) + '=' +
					encodeURIComponent(params[key]);
				}).join('&');
			}
			// Derive a fetch URL from the current URL, sans the GraphQL parameters.
			var graphqlParamNames = {
				query: true,
				variables: true,
				operationName: true
			};
			var otherParams = {};
			for (var k in parameters) {
				if (parameters.hasOwnProperty(k) && graphqlParamNames[k] !== true) {
					otherParams[k] = parameters[k];
				}
			}
			var fetchURL = locationQuery(otherParams);
			// Defines a GraphQL fetcher using the fetch API.
			function graphQLFetcher(graphQLParams) {
			return fetch("/graphql", {
				method: 'post',
				headers: {
					'Accept': 'application/json',
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(graphQLParams),
				credentials: 'include',
			}).then(function (response) {
				return response.text();
			}).then(function (responseBody) {
				try {
					return JSON.parse(responseBody);
				} catch (error) {
					return responseBody;
				}
			});
			}
			// When the query and variables string is edited, update the URL bar so
			// that it can be easily shared.
			function onEditQuery(newQuery) {
				parameters.query = newQuery;
				updateURL();
			}
			function onEditVariables(newVariables) {
				parameters.variables = newVariables;
				updateURL();
			}
			function onEditOperationName(newOperationName) {
				parameters.operationName = newOperationName;
				updateURL();
			}
			function updateURL() {
				history.replaceState(null, null, locationQuery(parameters));
			}
			// Render <GraphiQL /> into the body.
			ReactDOM.render(
				React.createElement(GraphiQL, {
					fetcher: graphQLFetcher,
					onEditQuery: onEditQuery,
					onEditVariables: onEditVariables,
					onEditOperationName: onEditOperationName,
				}),
				document.getElementById("root")
			);
		</script>
	</body>
</html>
`, graphiqlVersion))

const sessionKey = "keepup.sid"

//SessionMiddleware for creating a session on all routes, once the user is logged in.
func SessionMiddleware(
	ctx context.Context,
	store database.DAL,
	secret string,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hashKey := []byte(secret)

			blockKey := []byte(nil)

			s := securecookie.New(hashKey, blockKey)

			cookie, err := r.Cookie(sessionKey)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			value := make(map[string]string)
			if err == nil {
				err = s.Decode(sessionKey, cookie.Value, &value)
			}
			if err != nil {
				log.Error(err)
				http.Error(w, "error decoding cookie value", http.StatusInternalServerError)
				return
			}

			session, err := store.FindSessionByID(ctx, value["sessID"])
			if err != nil {
				log.Error(err)
				http.Error(w, "error finding session", http.StatusInternalServerError)
				return
			}

			user, err := store.FindUserByID(ctx, session.UserID)
			if err != nil {
				log.Error(err)
				http.Error(w, "error finding session", http.StatusInternalServerError)
				return
			}

			type contextUser struct{}
			currentUser := contextUser{}
			ctx = context.WithValue(ctx, currentUser, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequestIDMiddleware middleware set the requestID
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			r.Header.Set("X-Request-ID", uuid.NewV4().String())
		}

		next.ServeHTTP(w, r)
	})
}

//LoggingMiddleware to help with find errors
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		requestID := r.Header.Get("X-Request-ID")
		elapsed := time.Now().Sub(start)
		log.WithFields(log.Fields{
			"request-id": requestID,
			"method":     r.Method,
			"url":        r.URL.String(),
			"elapsed":    elapsed,
		}).Info()
	})
}

// HandleGoogleAuth handles the Google Authentication route
func HandleGoogleAuth(ctx context.Context, cfg oauth2.Config, secret string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hashKey := []byte(secret)

		blockKey := []byte(nil)

		s := securecookie.New(hashKey, blockKey)

		stateID := uuid.NewV4().String()

		state := map[string]string{
			"stateValue": stateID,
		}

		if encoded, err := s.Encode("state", state); err == nil {
			c := &http.Cookie{
				Name:   "state",
				Value:  encoded,
				Path:   "/",
				MaxAge: 10000,
			}
			if err != nil {
				log.Error(err)
				http.Error(w, "error making cookie", http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, c)
		}

		url := cfg.AuthCodeURL(stateID, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

// HandleGoogleCallback handles the google callback token exchange
func HandleGoogleCallback(
	ctx context.Context,
	cfg oauth2.Config,
	secret string,
	userInfo url.URL,
	store database.DAL,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryState := r.FormValue("state")
		if queryState == "" {
			http.Error(w, "no state recieved", http.StatusInternalServerError)
			return
		}

		hashKey := []byte(secret)

		blockKey := []byte(nil)

		s := securecookie.New(hashKey, blockKey)

		value := make(map[string]string)
		if cookie, err := r.Cookie("state"); err == nil {
			err = s.Decode("state", cookie.Value, &value)
		} else {
			log.Error(err)
			http.Error(w, "error decoding cookie value", http.StatusInternalServerError)
			return
		}

		if value["stateValue"] != queryState {
			http.Error(w, "stored state does not equal query state", http.StatusInternalServerError)
			return
		}

		code := r.FormValue("code")
		if code == "" {
			http.Error(w, "no code received", http.StatusInternalServerError)
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
			ID:         `json:"id"`,
			Name:       info.Name,
			Email:      info.Email,
			AvatarURL:  info.Picture,
			Provider:   "google",
			ProviderID: info.ID,
		}

		err = store.FindOrCreateUser(ctx, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, err := store.CreateSession(ctx, user.ID)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}

		cookieValue := map[string]string{
			"sessID": session.ID,
		}

		if encoded, err := s.Encode(sessionKey, cookieValue); err == nil {
			c := &http.Cookie{
				Name:   sessionKey,
				Value:  encoded,
				Path:   "/",
				MaxAge: 10000,
			}
			http.SetCookie(w, c)
			r.AddCookie(c)

		}

		http.Redirect(w, r, "https://localhost:3001", http.StatusSeeOther)
		fmt.Printf("Content: %s\n", userJSON)

	}
}
