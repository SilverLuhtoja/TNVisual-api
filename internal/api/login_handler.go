package api

import (
	"fmt"
	"net/http"
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

	user, err := cfg.DB.AuthenticateUser(r.Context(), req.Username)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprint("LoginHandler - ", err))
		return
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		RespondWithError(w, http.StatusNotAcceptable, "Invalid login credentials ")
		return
	}

	SetCookieHandler(w, r, user.ApiKey)

	// TODO: change apikey if once set
	RespondWithJSON(w, 200, "cookie set")
}
