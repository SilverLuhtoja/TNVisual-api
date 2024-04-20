package api

import (
	"fmt"
	"net/http"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
)

type LoginRequestResource struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cfg *ApiConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	req, err := GetParamsFromRequestBody(LoginRequestResource{}, r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprint("LoginHandler - ", err))
		return
	}

	user, err := cfg.DB.AuthenticateUser(r.Context(), database.AuthenticateUserParams(req))
	if err != nil {
		RespondWithError(w, http.StatusNotAcceptable, "Invalid login credentials ")
		return
	}

	SetCookieHandler(w, r, user.ApiKey)
	RespondWithJSON(w, 200, "cookie set")
}
