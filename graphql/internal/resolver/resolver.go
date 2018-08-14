package resolver

import (
	"context"

	"github.com/abbeyhrt/keep-up/graphql/internal/database"
	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	"github.com/abbeyhrt/keep-up/graphql/internal/session"
	graphql "github.com/graph-gophers/graphql-go"
	log "github.com/sirupsen/logrus"
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

func (r *viewerResolver) HomeID() *graphql.ID {
	homeID := graphql.ID(r.user.HomeID)
	return &homeID
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

func (r *Resolver) Home(ctx context.Context) (*homeResolver, error) {
	s, ok := session.FromContext(ctx)
	if !ok {
		return nil, nil
	}

	u, err := r.store.FindUserByID(ctx, s.User.ID)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	h, err := r.store.GetHomeByID(ctx, u.HomeID)

	return &homeResolver{h}, nil
}

func (r *Resolver) CreateHome(ctx context.Context, args *struct {
	Name        string
	Description string
}) (*homeResolver, error) {
	home := models.Home{
		Name:        args.Name,
		Description: args.Description,
	}

	s, ok := session.FromContext(ctx)
	if !ok {
		return nil, nil
	}

	h, err := r.store.CreateHome(ctx, home, s.User.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &homeResolver{h}, nil
}

type homeResolver struct {
	home models.Home
}

func (r *homeResolver) ID() graphql.ID {
	return graphql.ID(r.home.ID)
}

func (r *homeResolver) Name() string {
	return r.home.Name
}

func (r *homeResolver) Description() string {
	return r.home.Description
}

func (r *homeResolver) AvatarURL() *string {
	return &r.home.AvatarURL
}

func (r *homeResolver) CreatedAt() string {
	return r.home.CreatedAt.String()
}

func (r *homeResolver) UpdatedAt() string {
	return r.home.UpdatedAt.String()
}
