package httpsystem

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SystemRoute(r *gin.Engine) {
	r.GET("/public-key", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I live",
			"data":    os.Getenv("PUBLIC_KEY"),
		})
	})

	r.GET("/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I live",
		})
	})

}
