package models

import "time"

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	AvatarUrl  string    `json:"avatar_url"`
	Provider   string    `json:"provider"`
	ProviderID string    `json:"provider_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
