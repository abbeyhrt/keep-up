package models

import "time"

//User =`json:"id"` thing is a struct tag and is optional
type User struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	HomeID     *string   `json:"home_id"`
	Email      string    `json:"email"`
	AvatarURL  *string   `json:"avatar_url"`
	Provider   string    `json:"provider"`
	ProviderID string    `json:"provider_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Session struct for storing into the database
type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
}

//Home struct for manipulating home data
type Home struct {
	ID          string
	Name        string
	Description string
	AvatarURL   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//Task struct for users' tasks
type Task struct {
	ID          string
	UserID      string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
