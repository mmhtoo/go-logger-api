package log

import (
	"context"
	"fmt"
	"log"
	"time"

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

func buildSelectByProjectIdWithFilterQuery(input *SelectByProjectIdWithFilterInput) (string, []interface{}) {
	query := `SELECT * FROM log_groups WHERE project_id = $1 `
	argPosition := 1
	args := []any{
		input.projectId,
	}

	if input.logType != "" {
		argPosition += 1
		query += fmt.Sprintf("AND log_type = $%d ", argPosition)
		args = append(args, input.logType)
	}
	if input.pathName != "" {
		argPosition += 1
		query += fmt.Sprintf("AND path_name ILIKE $%d ", argPosition)
		args = append(args, input.pathName)
	}
	if input.keyword != "" {
		argPosition += 1
		query += fmt.Sprintf("AND payload ILIKE $%d ", argPosition)
		args = append(args, input.keyword)
	}
	if input.fromTime != (time.Time{}) && input.toTime == (time.Time{}) {
		argPosition += 1
		query += fmt.Sprintf("AND logged_at >= $%d ", argPosition)
		args  = append(args, input.fromTime)
	}
	if input.fromTime != (time.Time{}) && input.toTime != (time.Time{}) {
		query += fmt.Sprintf("AND logged_at BETWEEN $%d AND $%d ", argPosition + 1, argPosition + 2)
		args  = append(args, input.fromTime, input.toTime)
		argPosition += 2
	}
	query += fmt.Sprintf("ORDER BY logged_at DESC  LIMIT $%d OFFSET $%d", argPosition + 1, argPosition + 2)
	args = append(args, input.pageSize, (input.page -1) * input.pageSize)
	return query, args
}

func (repo *LogReposistory) SelectByProjectIdWithFilter(
	input *SelectByProjectIdWithFilterInput,
	ctx context.Context,
) (*[]LogEntity, error) {
	query, args := buildSelectByProjectIdWithFilterQuery(input)
	rows, err := repo.database.Connection.QueryContext(
		ctx, 
		query,
		args...
	)
	if err != nil {
		return nil, err
	}
	var logs []LogEntity
	for rows.Next(){
		var log LogEntity
		if err := rows.Scan(
			&log.id,
			&log.logType,
			&log.loggedAt,
			&log.loggedBy,
			&log.pathName,
			&log.projectId,
			&log.payload,
		); err != nil {
			return nil, nil
		} 
		logs = append(logs, log)
	}
	return &logs, nil
}