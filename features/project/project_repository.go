package project

import (
	"context"
	"errors"
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
		ORDER BY created_at DESC
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

func (repo *ProjectRepository) UpdateById(input *ProjectUpdateInput, ctx context.Context) (error) {
	query := `
		UPDATE projects
		SET name = $1, description = $2, project_type = $3
		WHERE id = $4
	`
	result, err := repo.database.Connection.ExecContext(
		ctx, 
		query, 
		input.Name,
		input.Description,
		input.ProjectType,
		input.Id,
	)
	if err != nil {
		return errors.New("Failed to update project with error:"+err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil && rowsAffected == 0 {
		return errors.New("Project was not found to update!")
	}
	return nil
}

func (repo *ProjectRepository) findById(id string, ctx context.Context) (ProjectEntity, error) {
	query := `
		SELECT * FROM projects
		WHERE id = $1
	`
	row := repo.database.Connection.QueryRowContext(
		ctx, 
		query, 
		id,
	)
	var project ProjectEntity
	if row != nil {
		return project, errors.New("Project was not found!")
	}
	if err := row.Scan(
		&project.Id,
		&project.Name,
		&project.Description,
		&project.ProjectType,
		&project.CreatedUserId,
		&project.CreatedAt,
	); err != nil {
		log.Fatalf("Error while retrieving project: %s", err)
		return  project, errors.New("Failed to retrieve!")
	}
	return project, nil
}
