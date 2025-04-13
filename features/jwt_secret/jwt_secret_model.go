package jwt_secret

import "time"

type JwtSecretEntity struct {
	id string `db:"id"`
	keyName string `db:"key_name"`
	privateKey string `db:"private_key"`
	publicKey string `db:"public_key"`
	createdAt time.Time `db:"created_at"`
	updatedAt time.Time `db:"updated_at"`
	isActive bool `db:"is_active"`
	createdUserId string `db:"created_user_id"`
	updatedUserId string `db:"updated_user_id"`
	projectId string `db:"project_id"`
}

type SaveJwtSecretInput struct {
	KeyName string
  PrivateKey string
	PublicKey string
	ProjectId string
	IsActive bool
	CreatedUserId string
}

type CreateAndSaveJwtSecretInput struct {
	KeyName string
	ProjectId string
	UserId string
}


