package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
)

type UserProvider interface {
	CreateUser(context.Context, database.CreateUserParams)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cfg *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req, err := GetParamsFromRequestBody(CreateUserRequest{}, r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprint("createUserHandler - ", err))
		return
	}

	hassed_pass, err := hashPassword(req.Password)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createUserHandler - ", err))
		return
	}

	userParams := database.CreateUserParams{
		Username: req.Username,
		Password: hassed_pass,
	}

	_, err = cfg.DB.CreateUser(r.Context(), userParams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("createUserHandler [couldn't create user to database] - ", err))
		return
	}

	RespondWithJSON(w, http.StatusCreated, "User created successfully")
}

func (cfg *ApiConfig) updateUserApiKey(r context.Context, id int32) error {
	key, err := generateAPIKey()
	if err != nil {
		return err
	}

	err = cfg.DB.UpdateUserKey(r, database.UpdateUserKeyParams{
		ID:     id,
		ApiKey: key,
	})

	if err != nil {
		return err
	}

	return nil
}
