package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SilverLuhtoja/TNVisual/src/api/auth/resources"
	"github.com/SilverLuhtoja/TNVisual/src/common"
)

type AuthController struct {
	Interactor AuthInteractor
}

func NewLoginController(inter AuthInteractor) *AuthController {
	return &AuthController{inter}
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req, err := common.GetParamsFromRequestBody(resources.LoginRequestResource{}, r)
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, fmt.Sprint("Login - ", err))
		return
	}

	httpStatus, err := controller.Interactor.Process(w, req)
	if err != nil {
		common.RespondWithError(w, httpStatus, fmt.Sprint("Login - ", err))
		return
	}

	common.RespondWithJSON(w, httpStatus, nil)
}

// Checks database for users with incoming apiKey
func (controller *AuthController) VerifyKey(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("tnsCookie")
	if err != nil {
		common.RespondWithError(w, http.StatusUnauthorized, http.ErrNoCookie.Error())
		return
	}

	common.RespondWithJSON(w, http.StatusOK, strings.Split(cookie.Value, " ")[1])
}

// TODO: DO I NEED THIS

// func (controller *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
// 	err := controller.Interactor.Verify(r.Header)
// 	if err != nil {
// 		common.RespondWithError(w, http.StatusUnauthorized, err.Error())
// 		return
// 	}

// 	common.RespondWithJSON(w, 200, "Key OK!")
// }
