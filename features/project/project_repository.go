package project

import (
	"context"
	"log"

	"github.com/mmhtoo/go-logger-api/config"
)

type ProjectRepository struct {
	database *config.Database
}

func NewProjectRepository(database *config.Database) *ProjectRepository {
	return &ProjectRepository{
		database: database,
	}
}

func (repo *ProjectRepository) GetAll(ctx context.Context) ([]ProjectEntity, error){
	query := `
		SELECT * FROM projects
	`
	rows, err := repo.database.Connection.QueryContext(
		ctx, 
		query,
	)
	defer rows.Close()
	if err != nil {
		log.Fatalf("Error while getting all projects: %s", err)
		return nil, err
	}

	var projects []ProjectEntity

	for rows.Next(){
		var project ProjectEntity
		if err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.ProjectType,
			&project.CreatedUserId,
			&project.CreatedAt,
		); err != nil {
			return nil , err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (repo *ProjectRepository) Save(input *ProjectCreateInput, ctx context.Context) (ProjectEntity,error){
	query := `
		INSERT INTO projects (id, name, description, project_type, created_user_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, description, project_type, created_user_id, created_at
	`
	row := repo.database.Connection.QueryRowContext(
		ctx, 
		query,
		input.Id,
		input.Name,
		input.Description,
		input.ProjectType,
		input.CreatedUserId,
	)
	var project ProjectEntity
	if err := row.Scan(
		&project.Id,
		&project.Name,
		&project.Description,
		&project.ProjectType,
		&project.CreatedUserId,
		&project.CreatedAt,
	); err != nil {
		log.Fatalf("Error while saving project: %s", err)
		return project, err
	}
	
	return project, nil
}