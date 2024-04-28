package project

import (
	"context"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/SilverLuhtoja/TNVisual/src/common"
)

type ProjectRepositry interface {
	Save(database.CreateProjectsParams) error
	FindAll() ([]database.Project, error)
}

type ProjectDatabase struct {
	db *database.Queries
}

func NewLoginRepository(db *database.Queries) *ProjectDatabase {
	return &ProjectDatabase{db}
}

func (repo *ProjectDatabase) Save(entity database.CreateProjectsParams) error {
	_, err := repo.db.CreateProjects(context.Background(), entity)
	if err != nil {
		return common.NewError("Failed to save ", err)
	}
	return nil
}

func (repo *ProjectDatabase) FindAll() ([]database.Project, error) {
	entities, err := repo.db.GetProjects(context.Background())
	if err != nil {
		return nil, common.NewError("Failed to get projects ", err)
	}

	return entities, nil
}
