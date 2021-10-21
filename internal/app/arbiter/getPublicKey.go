package arbiter

import (
	"github.com/farit2000/arbiter/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPublicKey(c *gin.Context) {
	c.String(http.StatusOK, utils.GetPublicKeyString(publicKey))
}
