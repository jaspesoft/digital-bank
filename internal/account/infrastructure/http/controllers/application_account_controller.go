package accountcontroller

import (
	requestsaccount "digital-bank/internal/account/infrastructure/http/requests"
	accountpersistence "digital-bank/internal/account/infrastructure/persistence"
	usecaseaccount "digital-bank/internal/account/usecase"
	systempersistence "digital-bank/internal/system/infrastructure/persistence"
	systemusecase "digital-bank/internal/system/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
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

	jsonApplicationAccountCompanyRequest.OwnerRecord = appClient.GetValue().GetIdentifier()

	fmt.Println(jsonApplicationAccountCompanyRequest)

	resp := usecaseaccount.NewApplicationAccount(
		accountpersistence.NewAccountMongoRepository(),
	).Run(jsonApplicationAccountCompanyRequest)

	if !resp.IsOk() {
		c.JSON(resp.GetError().GetHTTPCode(), gin.H{"error": resp.GetError().Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": resp.GetValue()})

}
