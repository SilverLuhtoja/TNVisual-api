package auth

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/SilverLuhtoja/TNVisual/src/api/auth/resources"
	"github.com/SilverLuhtoja/TNVisual/src/common"
	"golang.org/x/crypto/bcrypt"
)

type AuthInteractor struct {
	LoginRepositry AuthRepositry
}

func NewLoginInteractor(repo AuthRepositry) *AuthInteractor {
	return &AuthInteractor{repo}
}

func (inter AuthInteractor) Process(w http.ResponseWriter, params resources.LoginRequestResource) (int, error) {
	user, err := inter.authenticate(params)
	if err != nil {
		return http.StatusBadRequest, err
	}
	SetCookieHandler(w, user.ApiKey)
	err = inter.updateUserApiKey(user.ID)
	if err != nil {
		return http.StatusInternalServerError, common.NewError("Process [updateKey] ", err)
	}

	return http.StatusOK, nil
}

// func (inter AuthInteractor) Verify(r http.Header) error {
// 	key, err := middleware.GetApiKeyFromHeaders(r)
// 	if err != nil {
// 		return common.NewError("Verify ", err)
// 	}

// 	err = inter.LoginRepositry.GetUserByKey(key)
// 	if err != nil {
// 		return errors.New("unauthorized request")
// 	}

// 	return nil
// }

// func (inter LoginInteractor) Authenticate(params resources.LoginRequestResource) (int, error) {
func (inter AuthInteractor) authenticate(params resources.LoginRequestResource) (database.User, error) {

	user, err := inter.LoginRepositry.FindUser(params.Username)
	if err != nil {
		return database.User{}, common.NewError("Authenticate [findUser] ", err)
	}

	if !checkPasswordHash(params.Password, user.Password) {
		return database.User{}, common.NewError("Invalid login credentials ", err)
	}

	return user, nil
}

func (inter AuthInteractor) updateUserApiKey(id int32) error {
	key, err := generateAPIKey()
	if err != nil {
		return err
	}

	return inter.LoginRepositry.UpdateUserApiKey(id, key)
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
