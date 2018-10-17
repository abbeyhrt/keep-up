package resolver

import (
	"context"
	"fmt"

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

func (r *Resolver) Viewer(ctx context.Context) (*viewerResolver, error) {

	s, ok := session.FromContext(ctx)
	if !ok {
		return nil, nil
	}

	if s.User.HomeID == nil {
		return &viewerResolver{
			user:  s.User,
			home:  nil,
			tasks: nil,
		}, nil
	}

	home, err := r.store.GetHomeByID(ctx, s.User.HomeID)
	if err != nil {
		log.Error(err)
	}

	tasks, err := r.store.GetTasksByUserID(ctx, s.User.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resolvers := make([]*taskResolver, len(tasks))

	for i, task := range tasks {
		resolvers[i] = &taskResolver{task}
	}

	return &viewerResolver{
		user:  s.User,
		home:  &home,
		tasks: resolvers,
	}, nil
}

type viewerResolver struct {
	user  models.User
	home  *models.Home
	tasks []*taskResolver
}

func (r *viewerResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}

func (r *viewerResolver) FirstName() string {
	return r.user.FirstName
}

func (r *viewerResolver) LastName() string {
	return r.user.LastName
}

func (r *viewerResolver) Home() *homeResolver {
	if r.home == nil {
		return nil
	}
	return &homeResolver{*r.home}
}

func (r *viewerResolver) Tasks() []*taskResolver {
	if r.tasks == nil {
		return nil
	}
	return r.tasks
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

type userResolver struct {
	user models.User
}

func (r *Resolver) Users(ctx context.Context, args *struct {
	Name string
}) (*[]*userResolver, error) {
	users, err := r.store.GetUsersByName(ctx, args.Name)
	if err != nil {
		return nil, err
	}

	resolvers := make([]*userResolver, len(users))

	for i, user := range users {
		resolvers[i] = &userResolver{user}
	}
	return &resolvers, nil
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}

func (r *userResolver) FirstName() string {
	return r.user.FirstName
}

func (r *userResolver) LastName() string {
	return r.user.LastName
}
func (r *userResolver) Email() string {
	return r.user.Email
}

func (r *userResolver) HomeID() *string {
	if r.user.HomeID == nil {
		return nil
	}

	return r.user.HomeID
}

func (r *userResolver) AvatarURL() *string {
	return &r.user.AvatarURL
}

// func (r *userResolver) CreatedAt() string {
// 	return r.user.CreatedAt.String()
// }

// func (r *userResolver) UpdatedAt() string {
// 	return r.user.UpdatedAt.String()
// }

func (r *Resolver) Tasks(ctx context.Context) ([]*taskResolver, error) {
	s, ok := session.FromContext(ctx)
	if !ok {
		return nil, nil
	}

	tasks, err := r.store.GetTasksByUserID(ctx, s.User.ID)
	if err != nil {
		return nil, err
	}

	resolvers := make([]*taskResolver, len(tasks))

	for i, task := range tasks {
		resolvers[i] = &taskResolver{task}
	}

	return resolvers, nil
}

func (r *Resolver) Task(ctx context.Context, args *struct {
	ID string
}) (*taskResolver, error) {
	task, err := r.store.GetTaskByID(ctx, args.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &taskResolver{task}, nil
}

func (r *Resolver) CreateTask(ctx context.Context, args *struct {
	Title       string
	Description string
}) (*taskResolver, error) {

	s, ok := session.FromContext(ctx)
	if !ok {
		return nil, nil
	}

	t := models.Task{
		Title:       args.Title,
		Description: args.Description,
	}

	task, err := r.store.CreateTask(ctx, t, s.User.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &taskResolver{task}, nil
}

type taskResolver struct {
	task models.Task
}

func (r *taskResolver) ID() graphql.ID {
	return graphql.ID(r.task.ID)
}
func (r *taskResolver) UserID() string {
	return r.task.UserID
}

func (r *taskResolver) Title() string {
	return r.task.Title
}

func (r *taskResolver) Description() string {
	return r.task.Description
}

func (r *taskResolver) CreatedAt() string {
	return r.task.CreatedAt.String()
}

func (r *taskResolver) UpdatedAt() string {
	return r.task.UpdatedAt.String()
}

func (r *Resolver) Home(ctx context.Context) (*homeResolver, error) {
	s, ok := session.FromContext(ctx)
	if !ok {
		return nil, nil
	}

	u, err := r.store.GetUserByID(ctx, s.User.ID)

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
