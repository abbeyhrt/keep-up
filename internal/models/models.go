package models

import "time"

//User =`json:"id"` thing is a struct tag and is optional
type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	AvatarURL  string    `json:"avatar_url"`
	Provider   string    `json:"provider"`
	ProviderID string    `json:"provider_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
}
