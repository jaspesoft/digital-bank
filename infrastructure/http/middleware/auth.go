package middleware

import (
	systemdomain "digital-bank/internal/system/domain"
	systempersistence "digital-bank/internal/system/infrastructure/persistence"
	systemusecase "digital-bank/internal/system/usecase"
	"digital-bank/pkg"
	"digital-bank/pkg/cache"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

type middleware struct {
	CompanyID     string
	Authorization string
}

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")

	strToken, err := pkg.DecryptData(bearerToken, os.Getenv("PRIVATE_KEY"))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	var m middleware
	err = json.Unmarshal([]byte(strToken), &m)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	resAppClient := searchClient(m.CompanyID)

	if !resAppClient.IsOk() {
		c.JSON(resAppClient.GetError().GetHTTPCode(), gin.H{"error": resAppClient.GetError().Error()})
		c.Abort()
		return
	}

	if resAppClient.GetValue().GetTokenAPI() != strToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("CompanyID", m.CompanyID)

	c.Next()

}

func searchClient(companyID string) systemdomain.Result[*systemdomain.AppClient] {

	resUserApp := systemusecase.NewSearchAppClient(
		systempersistence.NewSystemRedisRepository(),
	).Run(companyID)

	if resUserApp.IsOk() {
		return resUserApp
	}

	resUserApp = systemusecase.NewSearchAppClient(
		systempersistence.NewSystemMongoRepository(),
	).Run(companyID)

	if resUserApp.IsOk() {
		_ = cache.SaveData(companyID, resUserApp.GetValue().ToMap(), 20*time.Minute)
	}

	return resUserApp

}
