package session

import (
	"context"

	"github.com/abbeyhrt/keep-up/graphql/internal/models"
)

// Reference: https://godoc.org/context#Context

// Session is the type representing our sessions in Context
type Session struct {
	User models.User
	// Task models.Task
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// sessionKey is the key for session.Session values in Contexts. It is
// unexported; clients use session.NewContext and session.FromContext
// instead of using this key directly.
var sessionKey key

// NewContext returns a new Context that carries value s.
func NewContext(ctx context.Context, s *Session) context.Context {
	return context.WithValue(ctx, sessionKey, s)
}

// FromContext returns the Session value stored in ctx, if any.
func FromContext(ctx context.Context) (*Session, bool) {
	s, ok := ctx.Value(sessionKey).(*Session)
	return s, ok
}
