package jwt_secret

import (
	"context"
	"database/sql"
	"log"

	"github.com/mmhtoo/go-logger-api/helpers"
)

type JwtSecretService struct {
	jwtSecretRepository *JwtSecretRepository
}

func NewJwtSecretService(jwtSecretRepository *JwtSecretRepository) *JwtSecretService {
	return &JwtSecretService{
		jwtSecretRepository: jwtSecretRepository,
	}
}

func (service *JwtSecretService) CreateAndSaveJwtSecret(
	input *CreateAndSaveJwtSecretInput, 
	tx *sql.Tx,
) error {
	keyPair, err := helpers.NewRSAKeyPair(2048)
	if err != nil {
		log.Fatalf("Error while creating RSA key pair: %s", err)
		return err
	}
	privateKey, publicKey, err := keyPair.GetPEM()
	if err != nil {
		log.Fatalf("Error while getting PEM: %s", err)
		return err
	}
	if _, err = service.jwtSecretRepository.SaveWithTx(
		&SaveJwtSecretInput{
			KeyName: input.KeyName,
			PrivateKey: privateKey,
			PublicKey: publicKey,
			ProjectId: input.ProjectId,
			CreatedUserId: input.UserId,	
			IsActive: true,	
		},
		tx,
	); err != nil {
		return err
	}
	return nil
}

func (service *JwtSecretService) MakeJwtSecretUnactive(
	id string,
	ctx context.Context,
) (error) {
	err := service.jwtSecretRepository.MakeUnActive(id, ctx)
	return err
}

func (service *JwtSecretService) GetAllJwtSecretsByProjectId(
	projectId string,
	ctx context.Context,
) (*[]JwtSecretEntity,error) {
	jwtSecrets, err := service.jwtSecretRepository.GetAllByProjectId(projectId, ctx)
	if err != nil {
		return nil, err
	}
	return jwtSecrets, nil
}

func (service *JwtSecretService) GetTopActiveJwtSecretByProjectId(
	projectId string,
	ctx context.Context,
) (*JwtSecretEntity, error) {
	jwtSecret, err := service.jwtSecretRepository.GetTopActiveItemByProjectItem(
		projectId,
		ctx,
	)
	if err != nil {
		return nil, err
	}
	return jwtSecret, nil
}

func (service *JwtSecretService) GetDetailById(
	id string,
	ctx context.Context,
) (*JwtSecretEntity, error) {
	jwtSecret, err := service.jwtSecretRepository.GetById(
		id,
		ctx,
	)
	if err != nil {
		return nil, err
	}
	return &jwtSecret, nil
}