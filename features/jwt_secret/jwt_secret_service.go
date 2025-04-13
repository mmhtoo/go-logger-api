package jwt_secret

import (
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