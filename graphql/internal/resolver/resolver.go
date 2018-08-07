package resolver

import (
	"context"

	"github.com/abbeyhrt/keep-up/graphql/internal/database"
	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	"github.com/abbeyhrt/keep-up/graphql/internal/session"
	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct {
	store database.DAL
}

func New(store database.DAL) *Resolver {
	return &Resolver{store}
}

func (_ *Resolver) Viewer(ctx context.Context) (*viewerResolver, error) {
	s, ok := session.FromContext(ctx)
	if !ok {
		return nil, nil
	}

	return &viewerResolver{s.User}, nil
}

type viewerResolver struct {
	user models.User
}

func (r *viewerResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}

func (r *viewerResolver) Name() string {
	return r.user.Name
}

func (r *viewerResolver) Email() string {
	return r.user.Email
}

func (r *viewerResolver) AvatarURL() *string {
	return &r.user.AvatarURL
}

func (r *viewerResolver) CreatedAt() string {
	return r.user.CreatedAt.String()
}

func (r *viewerResolver) UpdatedAt() string {
	return r.user.UpdatedAt.String()
}
