package arbiter

import (
	"github.com/farit2000/arbiter/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"html/template"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Set template delimiters
	router.Delims("{[{", "}]}")
	// Set template func map
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": utils.FormatAsDate,
	})
	// Load static templates
	router.LoadHTMLFiles("web/index.tmpl")

	// Main Arbiter page
	router.GET("/", getMainPage)

	// Get timestamp data with signature
	router.GET("/ts", getTimeStamp)

	// Verify hash with signature
	router.GET("/verify", getVerifyHash)

	// Verify hash with signature
	router.GET("/ts_and_digest_hex", getTsAndDigestHex)

	// Get private key
	router.GET("/private", getPrivateKey)

	// Get public key
	router.GET("/public", getPublicKey)

	return router
}
