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

type JwtSecretResponseDto struct {
	Id string `json:"id"`
	KeyName string `json:"keyName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive bool `json:"isActive"`
	CreatedUserId string `json:"createdUserId"`
	UpdatedUserId string `json:"updatedUserId"`
	ProjectId string `json:"projectId"`
}

type JwtSecretDetailResponseDto struct {
	Id string `json:"id"`
	KeyName string `json:"keyName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive bool `json:"isActive"`
	CreatedUserId string `json:"createdUserId"`
	UpdatedUserId string `json:"updatedUserId"`
	ProjectId string `json:"projectId"`
	PrivateKey string `json:"privateKey"`	
	PublicKey string `json:"publicKey"`
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

func (jwtSecretEntity *JwtSecretEntity) ToResponseDto() *JwtSecretResponseDto {
	return &JwtSecretResponseDto{
		Id: jwtSecretEntity.id,
		KeyName: jwtSecretEntity.keyName,
		CreatedAt: jwtSecretEntity.createdAt,
		UpdatedAt: jwtSecretEntity.updatedAt,
		IsActive: jwtSecretEntity.isActive,
		CreatedUserId: jwtSecretEntity.createdUserId,
		UpdatedUserId: jwtSecretEntity.updatedUserId,	
		ProjectId: jwtSecretEntity.projectId,
	}
}

func (jwtSecretEntity *JwtSecretEntity) ToDetailResponseDto() *JwtSecretDetailResponseDto {
	return &JwtSecretDetailResponseDto{
		Id: jwtSecretEntity.id,
		KeyName: jwtSecretEntity.keyName,
		CreatedAt: jwtSecretEntity.createdAt,
		UpdatedAt: jwtSecretEntity.updatedAt,
		IsActive: jwtSecretEntity.isActive,
		CreatedUserId: jwtSecretEntity.createdUserId,
		UpdatedUserId: jwtSecretEntity.updatedUserId,		
		ProjectId: jwtSecretEntity.projectId,
		PrivateKey: jwtSecretEntity.privateKey,
		PublicKey: jwtSecretEntity.publicKey,
	}	
}

func MapJwtSecretEntitesToResDto(entities *[]JwtSecretEntity) *[]JwtSecretResponseDto{
	var dtos []JwtSecretResponseDto
	for _, entity := range *entities {
		dtos = append(dtos, *entity.ToResponseDto())
	}
	return &dtos
}

