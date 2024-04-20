package api

import (
	"fmt"
	"net/http"

	"github.com/SilverLuhtoja/TNVisual/internal/models"
)

func (cfg *ApiConfig) GetAllProjects(w http.ResponseWriter, r *http.Request) {

	dbProjects, err := cfg.DB.GetProjects(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("GetAllProjects - ", err))
		return
	}

	domainProjects, err := models.AllDatabaseProjectsToProjects(dbProjects)
	if err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("GetAllProjects mapping error [EntityToDomain] - ", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, domainProjects)
}

func (cfg *ApiConfig) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	request, err := GetParamsFromRequestBody(models.Project{}, r)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("CreateProjectHandler - ", err))
		return
	}

	projectEntity, err := models.ProjectToProjectEntity(request)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("CreateProjectHandler mapping error [DomainToEntity] - ", err))
		return
	}

	project, err := cfg.DB.CreateProjects(r.Context(), projectEntity)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("CreateProjectHandler [couldn't create project to database] - ", err))
		return
	}

	projectDomain, err := models.DatabaseProjectToProject(project)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprint("CreateProjectHandler mapping error [EntityToDomain] - ", err))
		return
	}
	RespondWithJSON(w, http.StatusCreated, projectDomain)
}
