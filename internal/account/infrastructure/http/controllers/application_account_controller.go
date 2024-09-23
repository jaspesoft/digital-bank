package accountcontroller

import (
	requestsaccount "digital-bank/internal/account/infrastructure/http/requests"
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

	//companyID, _ := c.Get("CompanyID")

	//appClient := systemusecase.NewSearchAppClient(
	//	systempersistence.NewSystemRedisRepository(),
	//).Run((companyID.(string)))
	//
	//jsonApplicationAccountCompanyRequest.OwnerRecord = appClient.GetValue().GetIdentifier()

	fmt.Println(jsonApplicationAccountCompanyRequest)

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": jsonApplicationAccountCompanyRequest})

}
