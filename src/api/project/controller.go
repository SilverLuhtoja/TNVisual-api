package project

import "net/http"

type ProjectController struct {
	Interactor ProjectInteractor
}

func NewProjectController(inter ProjectInteractor) *ProjectController {
	return &ProjectController{inter}
}

func (contr ProjectController) Create(w http.ResponseWriter, r *http.Request) {

}

func (contr ProjectController) GetProjects(w http.ResponseWriter, r *http.Request) {

}
