package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/mmhtoo/go-logger-api/config"
	"github.com/mmhtoo/go-logger-api/entities"
)

type ProjectRepository struct {
	Database *config.Database
}

type ProjectSaveInput struct {
	Id string
	Name string
	Description string
	ProjectType string
	CreatedUserId string
}

func (repo *ProjectRepository) Save(input *ProjectSaveInput, ctx context.Context) (error){
	_, err := repo.Database.Connection.ExecContext(
		ctx, 
		`
			INSERT INTO projects (id, name, description, project_type) 
			VALUES ($1, $2, $3, $4)
		`,
		input.Id,
		input.Name,
		input.Description,
		input.ProjectType,
	)
	if err != nil {
		log.Fatalf("Error while saving project: %s", err)
	}
	return nil
} 

func (repo *ProjectRepository) GetAll(ctx context.Context) ([]entities.ProjectEntity, error){
	rows, err := repo.Database.Connection.QueryContext(
		ctx,
		`SELECT * FROM projects`,
	)
	defer rows.Close()

	if err != nil {
		log.Fatalf("Error while getting all projects: %s", err)
		return nil, err
	}

	var projects []entities.ProjectEntity

	for rows.Next() {
		var project entities.ProjectEntity
		if err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.ProjectType,
			&project.CreatedUserId,
			&project.CreatedAt,
		);  err != nil {
			return nil, fmt.Errorf("Error while scanning row: %s", err)
		}
		projects = append(projects, project)
	}
	return projects, nil
}