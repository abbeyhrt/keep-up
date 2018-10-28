package handlers

import (
	"context"
	"net/http"

	"github.com/abbeyhrt/keep-up/graphql/internal/cookie"
	"github.com/abbeyhrt/keep-up/graphql/internal/database"
	"github.com/gorilla/securecookie"
	log "github.com/sirupsen/logrus"
)

type Middleware struct {
	Context      context.Context
	Store        database.DAL
	CookieSecret string
}

func NewMiddleware(
	ctx context.Context,
	store database.DAL,
	secret string,
) *Middleware {
	return &Middleware{
		Context:      ctx,
		Store:        store,
		CookieSecret: secret,
	}
}

const RedirectPath = "/login"

func (m *Middleware) Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie(cookie.SessionKey)
		if err != nil {
			http.Redirect(w, r, RedirectPath, http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SessionMiddleware(
	ctx context.Context,
	store database.DAL,
	secret string,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// hashKey := []byte(secret)
			// blockKey := []byte(nil)
			// sc := securecookie.New(hashKey, blockKey)

			// cookie, err := r.Cookie(sessionKey)
			// if err != nil {
			// http.Redirect(w, r, "/login", http.StatusSeeOther)
			// return
			// }
			// value := make(map[string]string)
			// if err == nil {
			// err = sc.Decode(sessionKey, cookie.Value, &value)
			// }
			// if err != nil {
			// log.Error(err)
			// http.Error(w, "error decoding cookie value", http.StatusInternalServerError)
			// return
			// }

			// s, err := store.GetSessionByID(ctx, value["sessID"])
			// if err != nil {
			// log.Error(err)
			// http.Error(w, "error finding session", http.StatusInternalServerError)
			// return
			// }

			// user, err := store.GetUserByID(ctx, s.UserID)
			// if err != nil {
			// log.Error(err)
			// http.Error(w, "error finding user", http.StatusInternalServerError)
			// return
			// }

			// ctx = session.NewContext(ctx, &session.Session{User: user})

			// next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
