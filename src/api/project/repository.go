package project

import "github.com/SilverLuhtoja/TNVisual/internal/database"

type ProjectRepositry interface {
}

type ProjectDatabase struct {
	db *database.Queries
}

func NewLoginRepository(db *database.Queries) *ProjectDatabase {
	return &ProjectDatabase{db}
}
