package arbiter

import (
	"encoding/hex"
	"github.com/farit2000/arbiter/internal/pkg/signature"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getVerifyHash(c *gin.Context) {
	signHexStr := c.DefaultQuery("sign", "")
	hashHexStr := c.DefaultQuery("hash", "")
	signHexByteArray, err := hex.DecodeString(signHexStr)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, "Invalid input sign: %s", signHexStr)
		return
	}
	hashHexByteArray, err := hex.DecodeString(hashHexStr)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, "Invalid input hash: %s", hashHexStr)
		return
	}
	if signature.Verify(publicKey, hashHexByteArray, signHexByteArray) == nil {
		c.String(http.StatusUnprocessableEntity, "%t", true)
	} else {
		c.String(http.StatusUnprocessableEntity, "%t", false)
	}
}
