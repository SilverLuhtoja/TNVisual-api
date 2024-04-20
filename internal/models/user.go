package models

import (
	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	ApiKey   string    `json:"api_key"`
}

func DatabaseUserToUser(user database.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		ApiKey:   user.ApiKey,
	}
}
