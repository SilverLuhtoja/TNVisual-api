package api

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Handler func(http.ResponseWriter, *http.Request)

func (cfg *ApiConfig) middlewareAuth(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := getApiKey(r.Header)
		if err != nil {
			RespondWithError(w, http.StatusNotAcceptable, err.Error())
			return
		}

		_, err = cfg.DB.GetUserByKey(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized request")
			return
		}

		handler(w, r)
	}
}

func (cfg *ApiConfig) AuthenticateKeyHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := getApiKey(r.Header)
	if err != nil {
		RespondWithError(w, http.StatusNotAcceptable, err.Error())
		return
	}

	_, err = cfg.DB.GetUserByKey(r.Context(), apiKey)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized request")
		return
	}

	RespondWithJSON(w, 200, "Key OK!")
}

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

func getApiKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateAPIKey() (string, error) {
	key := make([]byte, 16) // 16 bytes = 128 bits
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
