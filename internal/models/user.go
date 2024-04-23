package models

import (
	"github.com/SilverLuhtoja/TNVisual/internal/database"
)

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	ApiKey   string `json:"api_key"`
}

func DatabaseUserToUser(user database.User) User {
	return User{
		Username: user.Username,
		ApiKey:   user.ApiKey,
	}
}
