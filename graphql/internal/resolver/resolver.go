package resolver

import (
	"context"

	"github.com/abbeyhrt/keep-up/graphql/internal/database"
	"github.com/abbeyhrt/keep-up/graphql/internal/models"
	"github.com/abbeyhrt/keep-up/graphql/internal/session"
	graphql "github.com/graph-gophers/graphql-go"
	log "github.com/sirupsen/logrus"
)

// Resolver contains all the methods to be resolved and gives them access to the store
type Resolver struct {
	store database.DAL
}

func New(store database.DAL) *Resolver {
	return &Resolver{store}
}

// Viewer gathers info on the currently logged in user
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
		log.Errorf("This is the %s: ", err)
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
	return r.user.AvatarURL
}

func (r *viewerResolver) CreatedAt() string {
	return r.user.CreatedAt.String()
}

func (r *viewerResolver) UpdatedAt() string {
	return r.user.UpdatedAt.String()
}

// ----------------------- USER RESOLVERS ----------------------------- //

type UserResolver struct {
	user models.User
	home *models.Home
}

// Users gathers a colleciton os users based on their name
func (r *Resolver) Users(ctx context.Context, args *struct {
	Name string
}) (*[]*UserResolver, error) {
	users, err := r.store.GetUsersByName(ctx, args.Name)
	if err != nil {
		return nil, err
	}

	resolvers := make([]*UserResolver, len(users))

	for i, user := range users {
		resolvers[i] = &UserResolver{user, nil}
	}
	return &resolvers, nil
}

//User resolves for one user, by their id
func (r *Resolver) User(ctx context.Context, args *struct {
	ID string
}) (*UserResolver, error) {

	user, err := r.store.GetUserByID(ctx, args.ID)
	if err != nil {
		log.Errorf("Error finding user: %s", err)
		return nil, nil
	}

	if user.HomeID == nil {
		return &UserResolver{
			user: user,
			home: nil,
		}, nil
	}

	home, err := r.store.GetHomeByID(ctx, user.HomeID)
	if err != nil {
		log.Errorf("Error finding home: %s", err)
	}

	return &UserResolver{
		user: user,
		home: &home,
	}, nil
}

// UpdateUser updates all fields on the user and returns the userResolver with that user
func (r *Resolver) UpdateUser(ctx context.Context, args struct {
	User struct {
		ID        string
		FirstName *string
		LastName  *string
		Email     *string
		HomeID    *string
		AvatarURL *string
	}
}) (*UserResolver, error) {

	user, err := r.store.GetUserByID(ctx, args.User.ID)
	if err != nil {
		log.Errorf("this is the error: %s", err)
		return nil, err
	}

	if args.User.FirstName != nil {
		user.FirstName = *args.User.FirstName
	}

	if args.User.LastName != nil {
		user.LastName = *args.User.LastName
	}

	if args.User.Email != nil {
		user.Email = *args.User.Email
	}

	if args.User.HomeID != nil {
		user.HomeID = args.User.HomeID
	}

	if args.User.AvatarURL != nil {
		user.AvatarURL = args.User.AvatarURL
	}

	u, err := r.store.UpdateUser(ctx, user)
	if err != nil {
		log.Errorf("this is the error: %s", err)
		return nil, err
	}

	return &UserResolver{u, nil}, nil
}

// DeleteUser deletes user
func (r *Resolver) DeleteUser(ctx context.Context, args struct {
	ID string
}) (*UserResolver, error) {
	err := r.store.DeleteUser(ctx, args.ID)
	if err != nil {
		log.Errorf("Error deleting user: %s", err)
		return nil, err
	}

	return &UserResolver{}, nil
}

// ID resolves for field id
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}

// FirstName resolves for field FirstName
func (r *UserResolver) FirstName() string {
	return r.user.FirstName
}

// LastName resolves for field LastName
func (r *UserResolver) LastName() string {
	return r.user.LastName
}

// Email resolves for field Email
func (r *UserResolver) Email() string {
	return r.user.Email
}

// Home resolves for field HomeID
func (r *UserResolver) Home() *homeResolver {
	if r.home == nil {
		return nil
	}
	return &homeResolver{*r.home}
}

// AvatarURL resolves for field AvatarURL
func (r *UserResolver) AvatarURL() *string {
	return r.user.AvatarURL
}

// ---------------- TASK RESOLVERS ------------------ //

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

func (r *Resolver) UpdateTask(ctx context.Context, args *struct {
	Task struct {
		ID          string
		Title       *string
		Description *string
		UserID      *string
	}
}) (*taskResolver, error) {

	task, err := r.store.GetTaskByID(ctx, args.Task.ID)
	if err != nil {
		log.Errorf("this is the error: %s", err)
		return nil, err
	}

	if args.Task.Title != nil {
		task.Title = *args.Task.Title
	}

	if args.Task.Description != nil {
		task.Description = *args.Task.Description
	}

	if args.Task.UserID != nil {
		task.UserID = *args.Task.UserID
	}

	t, err := r.store.UpdateTask(ctx, task)
	if err != nil {
		log.Errorf("this is the error: %s", err)
		return nil, err
	}

	return &taskResolver{t}, nil
}

func (r *Resolver) DeleteTask(ctx context.Context, args struct {
	ID string
}) (*taskResolver, error) {
	err := r.store.DeleteTask(ctx, args.ID)
	if err != nil {
		log.Errorf("Error deleting task: %s", err)
		return nil, err
	}

	return &taskResolver{}, nil
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

// ---------------------- HOME RESOLVERS ----------------------- //

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
		log.Errorf("This is the %s: ", err)
		return nil, err
	}

	return &homeResolver{h}, nil
}

func (r *Resolver) UpdateHome(ctx context.Context, args *struct {
	Home struct {
		ID          string
		Name        *string
		Description *string
		AvatarURL   *string
	}
}) (*homeResolver, error) {

	home, err := r.store.GetHomeByID(ctx, &args.Home.ID)
	if err != nil {
		log.Errorf("this is the error: %s", err)
		return nil, err
	}

	if args.Home.Name != nil {
		home.Name = *args.Home.Name
	}

	if args.Home.Description != nil {
		home.Description = *args.Home.Description
	}

	if args.Home.AvatarURL != nil {
		home.AvatarURL = *args.Home.AvatarURL
	}

	h, err := r.store.UpdateHome(ctx, home)
	if err != nil {
		log.Errorf("this is the error: %s", err)
		return nil, err
	}

	return &homeResolver{h}, nil
}

// DeleteHome deletes a user's home
func (r *Resolver) DeleteHome(ctx context.Context, args struct {
	ID string
}) (*homeResolver, error) {
	err := r.store.DeleteHome(ctx, args.ID)
	if err != nil {
		log.Errorf("Error deleting Home: %s", err)
		return nil, err
	}

	return &homeResolver{}, nil
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
