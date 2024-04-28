package resource

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
)

type ProjectImageData struct {
	Name string `json:"name"`
	File string `json:"file"`
}

type Project struct {
	Name        string             `json:"name"`
	Content     []ProjectImageData `json:"content"`
	Description *string            `json:"description"`
}

func ProjectToProjectEntity(domain Project) (database.CreateProjectsParams, error) {
	contentJson, err := json.Marshal(domain.Content)
	if err != nil {
		return database.CreateProjectsParams{}, errors.New(fmt.Sprint("Error with json marshaling: ", err.Error()))
	}

	var description sql.NullString
	if domain.Description != nil {
		description = sql.NullString{String: *domain.Description, Valid: true}
	} else {
		description = sql.NullString{Valid: false}
	}

	return database.CreateProjectsParams{
		Name:        domain.Name,
		Content:     contentJson,
		Description: description,
	}, nil
}

func AllDatabaseProjectsToProjects(databaseProjects []database.Project) ([]Project, error) {
	var projects []Project

	for _, project := range databaseProjects {
		content := []ProjectImageData{}
		err := json.Unmarshal(project.Content, &content)
		if err != nil {
			return []Project{}, errors.New(fmt.Sprint("Error with json unmarshaling: ", err.Error()))
		}

		var description *string
		if project.Description.Valid {
			desc := project.Description.String
			description = &desc
		}

		projects = append(projects, Project{
			Name:        project.Name,
			Content:     content,
			Description: description,
		})
	}

	return projects, nil
}
