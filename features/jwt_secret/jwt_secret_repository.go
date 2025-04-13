package jwt_secret

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/mmhtoo/go-logger-api/config"
)

type JwtSecretRepository struct {
	database *config.Database
}

func NewJwtSecretRepository(database *config.Database) *JwtSecretRepository {
	return &JwtSecretRepository{
		database: database,
	}
}

func (repo *JwtSecretRepository) Save(
	input *SaveJwtSecretInput, 
	ctx context.Context,
) (JwtSecretEntity, error) {
	query := `
		INSERT INTO jwt_secrets 
		(id, key_name, private_key, public_key, project_id, created_user_id, is_active, updated_user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, key_name, private_key, public_key, project_id, created_user_id, created_at, updated_at, is_active
	`
	row := repo.database.Connection.QueryRowContext(
		ctx,
		query,
		uuid.NewString(),
		input.KeyName,
		input.PrivateKey,
		input.PublicKey,
		input.ProjectId,
		input.CreatedUserId,
		input.IsActive,
		input.CreatedUserId,
	)
	var jwtSecret JwtSecretEntity
	if err := row.Scan(
		&jwtSecret.id,
		&jwtSecret.keyName,
		&jwtSecret.privateKey,
		&jwtSecret.publicKey,
		&jwtSecret.privateKey,
		&jwtSecret.createdUserId,
		&jwtSecret.createdAt,
		&jwtSecret.updatedAt,
	); err != nil {
		return jwtSecret, err
	}
	return jwtSecret, nil
}

func (repo *JwtSecretRepository) SaveWithTx(
	input *SaveJwtSecretInput, 
	tx *sql.Tx,
) (JwtSecretEntity, error) {
	query := `
		INSERT INTO jwt_secrets 
		(id, key_name, private_key, public_key, project_id, created_user_id, is_active, updated_user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, key_name, private_key, public_key, project_id, created_user_id, created_at, updated_at, is_active
	`
	row := tx.QueryRow(
		query,
		uuid.NewString(),
		input.KeyName,
		input.PrivateKey,
		input.PublicKey,
		input.ProjectId,
		input.CreatedUserId,
		input.IsActive,
		input.CreatedUserId,
	)
	var jwtSecret JwtSecretEntity
	if err := row.Scan(
		&jwtSecret.id,
		&jwtSecret.keyName,
		&jwtSecret.privateKey,
		&jwtSecret.publicKey,
		&jwtSecret.privateKey,
		&jwtSecret.createdUserId,
		&jwtSecret.createdAt,
		&jwtSecret.updatedAt,
		&jwtSecret.isActive,
	); err != nil {
		return jwtSecret, err
	}
	return jwtSecret, nil
}

func (repo *JwtSecretRepository) MakeUnActive(
	id string,
	ctx context.Context,
) (error) {
	query := `
		UPDATE jwt_secrets
		SET is_active = false
		WHERE id = $1
	`
	_, err := repo.database.Connection.ExecContext(
		ctx,
		query,
		id,
	)
	if err != nil {
		log.Printf("Failed to update jwt secret with error: %s", err)
		return err
	}
	return nil
}

func (repo *JwtSecretRepository) GetAllByProjectId(
	projectId string,
	ctx context.Context,
) (*[]JwtSecretEntity, error) {
	query := `
		SELECT * FROM jwt_secrets
		WHERE project_id = $1
	`
	rows, err := repo.database.Connection.QueryContext(
		ctx, 
		query,
		projectId,
	)
	defer rows.Close()
	if err != nil {
		log.Printf("Failed to get jwt secrets with error: %s", err)
		return nil, err
	}
	var jwtSecrets []JwtSecretEntity
	for rows.Next() {
		var jwtSecret JwtSecretEntity
		if err := rows.Scan(
			&jwtSecret.id,
			&jwtSecret.keyName,
			&jwtSecret.privateKey,
			&jwtSecret.publicKey,
			&jwtSecret.createdAt,
			&jwtSecret.updatedAt,	
			&jwtSecret.createdUserId,
			&jwtSecret.updatedUserId,
			&jwtSecret.projectId,
			&jwtSecret.isActive,	
		); err != nil {
			log.Printf("Failed to scan jwt secret with error: %s", err)
			return nil, err
		}
		jwtSecrets = append(jwtSecrets, jwtSecret)
	}
	return &jwtSecrets, nil
}

func (repo *JwtSecretRepository) GetTopActiveItemByProjectItem(
	projectId string,
	ctx context.Context,
)(*JwtSecretEntity, error){
	query := `
		SELECT * FROM jwt_secrets
		WHERE is_active = true and project_id = $1
		ORDER BY updated_at DESC
		LIMIT 1
	`
	row := repo.database.Connection.QueryRowContext(
		ctx, 
		query,
		projectId,
	)
	var jwtSecret JwtSecretEntity
	if err := row.Scan(
		&jwtSecret.id,
		&jwtSecret.keyName,
		&jwtSecret.privateKey,
		&jwtSecret.publicKey,
		&jwtSecret.createdAt,
		&jwtSecret.updatedAt,	
		&jwtSecret.createdUserId,
		&jwtSecret.updatedUserId,
		&jwtSecret.projectId,
		&jwtSecret.isActive,	
	); err != nil {
		log.Printf("Failed to get jwt secret with error: %s", err)
		return nil, err
	}
	return &jwtSecret, nil
}

func (repo *JwtSecretRepository) GetById(
	id string, 
	ctx context.Context,
) (JwtSecretEntity, error) {
	query := `
		SELECT * FROM jwt_secrets
		WHERE id = $1
	`
	row := repo.database.Connection.QueryRowContext(
		ctx,
		query,
		id,
	)
	var jwtSecret JwtSecretEntity
	if err := row.Scan(
		&jwtSecret.id,
		&jwtSecret.keyName,
		&jwtSecret.privateKey,
		&jwtSecret.publicKey,
		&jwtSecret.createdAt,
		&jwtSecret.updatedAt,	
		&jwtSecret.createdUserId,
		&jwtSecret.updatedUserId,
		&jwtSecret.projectId,
		&jwtSecret.isActive,		
	); err != nil {
		return jwtSecret,  err
	}
	return jwtSecret, nil
}