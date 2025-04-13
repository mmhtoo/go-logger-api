package log

import (
	"context"
	"log"

	"github.com/mmhtoo/go-logger-api/config"
)

type LogReposistory struct {
	database *config.Database
}

func NewLogRepository(database *config.Database) *LogReposistory {
	return &LogReposistory{
		database: database,
	}
}

func (repo *LogReposistory) Save(
	input *SaveLogInput,
	ctx context.Context,
) (error) {
	query := `
		INSERT INTO log_groups
		(id, log_type, logged_at, logged_by, path_name, project_id, payload)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	if _, err := repo.database.Connection.ExecContext(
		ctx, 
		query,
		input.Id,
		input.LogType,
		input.LoggedAt,
		input.LoggedBy,
		input.PathName,
		input.ProjectId,
		input.Payload,
	); err != nil {
		log.Printf("Failed to save log with error: %s", err)
		return err
	}
	return nil
}