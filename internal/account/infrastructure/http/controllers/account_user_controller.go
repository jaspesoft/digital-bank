package accountcontroller

import (
	"digital-bank/infrastructure/adapter"
	accountadapter "digital-bank/internal/account/infrastructure/adapter"
	accountpersistence "digital-bank/internal/account/infrastructure/persistence"
	usecaseaccount "digital-bank/internal/account/usecase"
	systempersistence "digital-bank/internal/system/infrastructure/persistence"
	systemusecase "digital-bank/internal/system/usecase"
	eventbus "digital-bank/pkg/event_bus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountUserRegisterController(c *gin.Context) {
	var req usecaseaccount.AccountUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
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

	resp := usecaseaccount.NewAccountUserRegister(
		accountpersistence.NewAccountUserMongoRepository(),
		accountadapter.NewHashPasswordAdapter(),
		adapter.NewUUIDEntityID(),
		eventbus.NewAWSEventBus(),
	).Run(req, *appClient.GetValue())

	if !resp.IsOk() {
		c.JSON(resp.GetError().GetHTTPCode(), gin.H{"error": resp.GetError().Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": resp.GetValue()})

}

func ChangePassswordController(c *gin.Context) {
	var req usecaseaccount.ChangePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := usecaseaccount.NewAccountUserChangePassword(
		accountpersistence.NewAccountUserMongoRepository(),
		accountadapter.NewHashPasswordAdapter(),
	).Run(req)

	if !resp.IsOk() {
		c.JSON(resp.GetError().GetHTTPCode(), gin.H{"error": resp.GetError().Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})

}
