package accountcontroller

import (
	requestsaccount "digital-bank/internal/account/infrastructure/http/requests"
	accountpersistence "digital-bank/internal/account/infrastructure/persistence"
	usecaseaccount "digital-bank/internal/account/usecase"
	systempersistence "digital-bank/internal/system/infrastructure/persistence"
	systemusecase "digital-bank/internal/system/usecase"
	credentials "digital-bank/pkg/service_credentials"
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

	companyID, _ := c.Get("CompanyID")

	appClient := systemusecase.NewSearchAppClient(
		systempersistence.NewAppClientRedisRepository(),
	).Run(companyID.(string))

	if !appClient.IsOk() {
		c.JSON(appClient.GetError().GetHTTPCode(), gin.H{"error": appClient.GetError().Error()})
		return
	}

	log.Println(jsonApplicationAccountCompanyRequest)

	credential, err := credentials.FindApplicationClientCredentials(appClient.GetValue())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountUser, err := accountpersistence.NewAccountUserMongoRepository().FindByEmail(jsonApplicationAccountCompanyRequest.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := usecaseaccount.NewApplicationAccount(
		accountpersistence.NewAccountMongoRepository(),
		layer2.NewLayer2Application(credential.Layer2Credentials),
	).Run(accountUser, *appClient.GetValue(), jsonApplicationAccountCompanyRequest)

	if !resp.IsOk() {
		c.JSON(resp.GetError().GetHTTPCode(), gin.H{"error": resp.GetError().Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": resp.GetValue()})

}
