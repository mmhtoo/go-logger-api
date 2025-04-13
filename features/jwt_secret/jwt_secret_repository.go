package jwt_secret

import (
	"context"
	"database/sql"

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