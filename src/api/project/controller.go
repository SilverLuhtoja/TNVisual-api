package project

import (
	"fmt"
	"net/http"

	resource "github.com/SilverLuhtoja/TNVisual/src/api/project/resources"
	"github.com/SilverLuhtoja/TNVisual/src/common"
)

type ProjectController struct {
	Interactor ProjectInteractor
}

func NewProjectController(inter ProjectInteractor) *ProjectController {
	return &ProjectController{inter}
}

func (contr ProjectController) CreateProjects(w http.ResponseWriter, r *http.Request) {
	req, err := common.GetParamsFromRequestBody(resource.Project{}, r)
	if err != nil {
		common.RespondWithError(w, http.StatusBadRequest, fmt.Sprint("CreateProjects [request]- ", err))
		return
	}

	err = contr.Interactor.Create(req)
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("CreateProjects - ", err))
		return
	}

	common.RespondWithJSON(w, http.StatusOK, "Project Created")
}

func (contr ProjectController) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := contr.Interactor.GetAllProjects()
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("CreateProjects - ", err))
		return
	}

	common.RespondWithJSON(w, http.StatusOK, projects)
}
