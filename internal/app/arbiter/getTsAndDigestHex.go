package arbiter

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getTsAndDigestHex(c *gin.Context) {
	tsStr := c.DefaultQuery("ts", "")
	digestHexStr := c.DefaultQuery("digest", "")
	log.Println(tsStr)
	log.Println(digestHexStr)
	digestByteArray, err := hex.DecodeString(digestHexStr)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, "Invalid input digest: %s", digestHexStr)
		return
	}
	hashHexStr := hex.EncodeToString(append([]byte(tsStr), digestByteArray...))
	c.String(http.StatusUnprocessableEntity, hashHexStr)
}
