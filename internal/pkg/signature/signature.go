package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

func Sign(hash [32]byte, privateKey *rsa.PrivateKey) (signature []byte, err error) {
	signature, err = rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hash[:], nil)
	if err != nil {
		log.Printf("Error from signing: %s\n", err)
		return
	}
	return signature, nil
}

func Verify(publicKey *rsa.PublicKey, msg, signature []byte) (err error) {
	hash := sha256.Sum256(msg)
	return rsa.VerifyPSS(publicKey, crypto.SHA256, hash[:], signature, nil)
}
