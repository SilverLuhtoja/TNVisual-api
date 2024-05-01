package user

import (
	"context"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
)

type UserDatabase struct {
	db *database.Queries
}

func NewUserRepostitory(db *database.Queries) *UserDatabase {
	return &UserDatabase{db}
}

func (repo *UserDatabase) SaveUser(userParams database.CreateUserParams) error {
	_, err := repo.db.CreateUser(context.Background(), userParams)
	return err
}
