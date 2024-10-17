package accountcontroller

import (
	requestsaccount "digital-bank/internal/account/infrastructure/http/requests"
	accountpersistence "digital-bank/internal/account/infrastructure/persistence"
	usecaseaccount "digital-bank/internal/account/usecase"
	"digital-bank/internal/system/infrastructure/service_credentials"
	"digital-bank/pkg/services/layer2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ApplicationAccountCompanyController(c *gin.Context) {
	var jsonApplicationAccountCompanyRequest requestsaccount.ApplicationAccountCompanyRequest

	if err := c.ShouldBindJSON(&jsonApplicationAccountCompanyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(jsonApplicationAccountCompanyRequest)

	credential, err := credentials.FindApplicationClientCredentials(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountUser, err := accountpersistence.NewAccountUserMongoRepository().FindByEmail(
		jsonApplicationAccountCompanyRequest.Email,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := usecaseaccount.NewApplicationAccount(
		accountpersistence.NewAccountMongoRepository(),
		layer2.NewLayer2Application(credential.Layer2Credentials),
	).Run(accountUser, jsonApplicationAccountCompanyRequest)

	if !resp.IsOk() {
		c.JSON(resp.GetError().GetHTTPCode(), gin.H{"error": resp.GetError().Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": resp.GetValue()})

}
