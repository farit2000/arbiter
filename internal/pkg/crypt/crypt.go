package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func GetPrivateAndPubKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		log.Fatal(err.Error())
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}
