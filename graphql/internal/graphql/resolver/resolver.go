package resolver

import (
	"github.com/abbeyhrt/keep-up/graphql/internal/database"
)

type Resolver struct {
	db database.DAL
}

func New(db database.DAL) *Resolver {
	return &Resolver{db}
}

func (r *Resolver) Viewer() *viewerResolver {
	return &viewerResolver{r.db}
}

type viewerResolver struct {
	db database.DAL
}

func (r *viewerResolver) Name() string {
	return "hello"
}

// func (r *viewerResolver) Email() string {

// }
