package arbiter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{
		"now": time.Date(2021, 07, 01, 0, 0, 0, 0, time.UTC),
	})
}
