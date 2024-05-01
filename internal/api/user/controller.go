package user

import (
	"fmt"
	"net/http"

	"github.com/SilverLuhtoja/TNVisual/internal/api/user/resources"
	"github.com/SilverLuhtoja/TNVisual/internal/common"
)

type UserController struct {
	Interactor UserInteractor
}

func NewUserController(inter UserInteractor) *UserController {
	return &UserController{inter}
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	req, err := common.GetParamsFromRequestBody(resources.CreateUserRequest{}, r)
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, fmt.Sprint("CreateUser [controller] - ", err))
		return
	}

	err = controller.Interactor.AddUser(req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("CreateUser [controller]  - ", err))
		return
	}

	common.RespondWithJSON(w, http.StatusCreated, "User created successfully")
}
