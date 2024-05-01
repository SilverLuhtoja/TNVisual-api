package project

import (
	resource "github.com/SilverLuhtoja/TNVisual/internal/api/project/resources"
	"github.com/SilverLuhtoja/TNVisual/internal/common"
)

type ProjectInteractor struct {
	ProjectRepositry ProjectRepositry
}

func NewProjectInteractor(repo ProjectRepositry) *ProjectInteractor {
	return &ProjectInteractor{repo}
}

func (inter ProjectInteractor) Create(params resource.Project) error {
	projectEntity, err := resource.ProjectToProjectEntity(params)
	if err != nil {
		return common.NewError("Create [mapper]", err)
	}

	err = inter.ProjectRepositry.Save(projectEntity)
	if err != nil {
		return common.NewError("Create ", err)
	}
	return nil
}

func (inter ProjectInteractor) GetAllProjects() ([]resource.Project, error) {
	projectEntities, err := inter.ProjectRepositry.FindAll()
	if err != nil {
		return nil, common.NewError("GetAllProjects ", err)
	}

	projects, err := resource.AllDatabaseProjectsToProjects(projectEntities)
	if err != nil {
		return nil, common.NewError("GetAllProjects [mapper]", err)
	}

	return projects, nil
}
