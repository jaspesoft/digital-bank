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
	"log"
	"net/http"
	"os"
	"time"
)

type middleware struct {
	CompanyID     string `json:"companyId"`
	Authorization string `json:"authorization"`
}

func CompanyClientAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("X-Company-Token")

	strToken, err := pkg.DecryptData(authHeader, os.Getenv("PRIVATE_KEY"))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	var m middleware
	err = json.Unmarshal([]byte(strToken), &m)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	resAppClient := searchClient(m.CompanyID)

	if !resAppClient.IsOk() {
		log.Println("CompanyClientAuthMiddleware resAppClient Error:", resAppClient.GetError().Error())
		c.JSON(resAppClient.GetError().GetHTTPCode(), gin.H{"message": resAppClient.GetError().Error()})
		c.Abort()
		return
	}

	if resAppClient.GetValue().GetTokenAPI() != m.Authorization {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	log.Println("company client authorized")

	c.Set("AppClient", resAppClient.GetValue())

	c.Next()

}

func searchClient(companyID string) systemdomain.Result[*systemdomain.AppClient] {
	log.Println("searchClient companyID: " + companyID)

	resUserApp := systemusecase.NewSearchAppClient(
		systempersistence.NewAppClientRedisRepository(),
	).Run(companyID)

	if resUserApp.IsOk() {
		return resUserApp
	}

	resUserApp = systemusecase.NewSearchAppClient(
		systempersistence.NewAppClientMongoRepository(),
	).Run(companyID)

	if resUserApp.IsOk() {
		_ = cache.SaveData(companyID, resUserApp.GetValue().ToMap(), 20*time.Minute)
	}

	return resUserApp

}
