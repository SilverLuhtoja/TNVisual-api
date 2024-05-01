package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/SilverLuhtoja/TNVisual/internal/common"
	"github.com/SilverLuhtoja/TNVisual/internal/infrastructure"
)

type Handler func(http.ResponseWriter, *http.Request)

func Authenticate(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := getApiKeyFromHeaders(r.Header)
		if err != nil {
			common.RespondWithError(w, http.StatusNotAcceptable, err.Error())
			return
		}

		err = isUserInDatabase(r.Context(), apiKey)
		if err != nil {
			common.RespondWithError(w, http.StatusUnauthorized, "Unauthorized request")
			return
		}

		handler(w, r)
	}
}

func getApiKeyFromHeaders(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization header included")
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

func isUserInDatabase(r context.Context, key string) error {
	db := infrastructure.NewDatabase()
	_, err := db.GetUserByKey(r, key)
	return err
}
