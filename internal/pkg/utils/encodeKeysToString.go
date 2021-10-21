package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GetPrivateKeyString(privateKey *rsa.PrivateKey) string {
	privatePEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	return string(privatePEM)
}

func GetPublicKeyString(publicKey *rsa.PublicKey) string {
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		},
	)
	return string(pubPEM)
}
