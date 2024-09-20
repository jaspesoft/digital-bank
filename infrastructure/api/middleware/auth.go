package middlewares

import (
	systempersistence "digital-bank/internal/system/infrastructure/persistence"
	systemusecase "digital-bank/internal/system/usecase"
	"digital-bank/pkg"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")

	base64Token, err := pkg.DecryptData(bearerToken, os.Getenv("PRIVATE_KEY"))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return

	}

	decodedAuthToken, err := base64.StdEncoding.DecodeString(base64Token)

	parts := strings.SplitN(string(decodedAuthToken), ":", 2)
	if len(parts) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	jsonInformationCompany := gin.H{
		"company": parts[0],
		"secret":  parts[1],
	}

	resUserApp := systemusecase.NewSearchAppClient(
		systempersistence.NewSystemMongoRepository(),
	).Run(jsonInformationCompany["company"].(string))

	if !resUserApp.IsOk() {
		c.JSON(resUserApp.GetError().HttpCode, gin.H{"error": resUserApp.GetError().Message})
		c.Abort()
		return
	}

	if resUserApp.GetValue().Secret != jsonInformationCompany["secret"] {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	//c.Set("CompanyID", cl.CompanyID)
	//c.Set("URLWebhook", cl.URLWebhook)

	c.Next()

}
