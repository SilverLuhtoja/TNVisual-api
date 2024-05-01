package user

import (
	"github.com/SilverLuhtoja/TNVisual/internal/api/user/resources"
	"github.com/SilverLuhtoja/TNVisual/internal/common"
	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	SaveUser(database.CreateUserParams) error
}

type UserInteractor struct {
	UserRepository UserRepository
}

func NewUserInteractor(repo UserRepository) *UserInteractor {
	return &UserInteractor{repo}
}

func (inter UserInteractor) AddUser(params resources.CreateUserRequest) error {
	hassed_pass, err := hashPassword(params.Password)
	if err != nil {
		return common.NewError("AddUser [service]", err)
	}
	dbParams := database.CreateUserParams{
		Username: params.Username,
		Password: hassed_pass,
	}
	return inter.UserRepository.SaveUser(dbParams)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
