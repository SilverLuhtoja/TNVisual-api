package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/SilverLuhtoja/TNVisual/internal/models"
	"github.com/google/uuid"
)

type UserProvider interface {
	CreateUser(context.Context, database.CreateUserParams)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TODO:  MAKE PASSWORD ENCRYPTION
func (cfg *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req, err := GetParamsFromRequestBody(CreateUserRequest{}, r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createUserHandler - ", err))
		return
	}

	userParams := database.CreateUserParams{
		ID:       uuid.New(),
		Username: req.Username,
		Password: req.Password,
	}

	user, err := cfg.DB.CreateUser(r.Context(), userParams)
	if err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createUserHandler [couldn't create user to database] - ", err))
		return
	}

	RespondWithJSON(w, http.StatusCreated, models.DatabaseUserToUser(user))
}
