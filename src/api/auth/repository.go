package auth

import (
	"context"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/SilverLuhtoja/TNVisual/src/common"
)

type AuthRepositry interface {
	FindUser(string) (database.User, error)
	UpdateUserApiKey(int32, string) error
	// GetUserByKey(string) error
}

type AuthDatabase struct {
	db *database.Queries
}

func NewLoginRepository(db *database.Queries) *AuthDatabase {
	return &AuthDatabase{db}
}

func (repo *AuthDatabase) FindUser(username string) (database.User, error) {
	// TODO: CHANGE THIS database METHOD TO FIND USER
	user, err := repo.db.AuthenticateUser(context.Background(), username)
	if err != nil {
		return database.User{}, common.NewError("FindUser", err)
	}
	return user, nil
}

func (repo *AuthDatabase) UpdateUserApiKey(id int32, key string) error {
	err := repo.db.UpdateUserKey(context.Background(), database.UpdateUserKeyParams{
		ID:     id,
		ApiKey: key,
	})
	return err
}

// func (repo *AuthDatabase) GetUserByKey(key string) error {
// 	_, err := repo.db.GetUserByKey(context.Background(), key)
// 	if err != nil {
// 		return common.NewError("GetUserByKey ", err)
// 	}
// 	return nil
// }
