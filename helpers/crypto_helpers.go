package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type RSAKeyPair struct {
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

func NewRSAKeyPair(bits int) (*RSAKeyPair, error){
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	keyPair := &RSAKeyPair{
		privateKey: privateKey,
		publicKey: &privateKey.PublicKey,
	}
	return keyPair, nil
}

func (keyPair *RSAKeyPair) GetPEM() (string, string, error){
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(keyPair.privateKey)
	privatePem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	publicKeyBytes := x509.MarshalPKCS1PublicKey(keyPair.publicKey)
	publicPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return string(privatePem), string(publicPem), nil
}