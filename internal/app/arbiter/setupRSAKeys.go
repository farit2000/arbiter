package arbiter

import (
	"crypto/rsa"
	"github.com/farit2000/arbiter/internal/pkg/crypt"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func SetupRSAKeys() {
	privateKey, publicKey = crypt.GetPrivateAndPubKeys()
}
