package arbiter

import (
	"github.com/farit2000/arbiter/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPrivateKey(c *gin.Context) {
	c.String(http.StatusOK, utils.GetPrivateKeyString(privateKey))
}
