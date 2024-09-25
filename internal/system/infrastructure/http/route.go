package httpsystem

import (
	"digital-bank/pkg"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SystemRoute(r *gin.Engine) {
	r.GET("/public-key", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I'm live",
			"data":    os.Getenv("PUBLIC_KEY"),
		})
	})

	r.POST("/generate-token", func(c *gin.Context) {

		var d map[string]interface{}
		if err := c.ShouldBindJSON(&d); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		jsonData, err := json.Marshal(d)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		jsonString := string(jsonData)
		t, err := pkg.EncryptData(jsonString, os.Getenv("PUBLIC_KEY"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return

		}

		c.String(http.StatusOK, t)
	})

}
