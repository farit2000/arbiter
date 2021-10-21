package arbiter

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/farit2000/arbiter/internal/pkg/signature"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type tsToken struct {
	Ts        string `json:"ts"`
	Signature string `json:"signature"`
}

type response struct {
	Status         int     `json:"status"`
	StatusString   string  `json:"status_string"`
	TimeStampToken tsToken `json:"time_stamp_token"`
}

func getTimeStamp(c *gin.Context) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	digestHexStr := c.DefaultQuery("digest", "")
	digestByteArray, err := hex.DecodeString(digestHexStr)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, "Invalid input digest: %s", digestHexStr)
		return
	}
	timeStamp := time.Now().In(loc).Format("2006-01-02T15:04:05.000+03")
	signingDataHash := sha256.Sum256(append([]byte(timeStamp), digestByteArray...))
	sign, err := signature.Sign(signingDataHash, privateKey)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, "Error while signing")
		return
	}
	resp := response{
		Status:       0,
		StatusString: "",
		TimeStampToken: tsToken{
			Ts:        timeStamp,
			Signature: hex.EncodeToString(sign),
		},
	}
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
