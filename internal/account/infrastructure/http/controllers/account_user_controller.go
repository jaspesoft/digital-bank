package accountcontroller

import (
	"digital-bank/infrastructure/adapter"
	accountadapter "digital-bank/internal/account/infrastructure/adapter"
	accountpersistence "digital-bank/internal/account/infrastructure/persistence"
	usecaseaccount "digital-bank/internal/account/usecase"
	credentials "digital-bank/internal/system/infrastructure/service_credentials"
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

	appClient := credentials.SearchApplicationClient(c)

	resp := usecaseaccount.NewAccountUserRegister(
		accountpersistence.NewAccountUserMongoRepository(),
		accountadapter.NewHashPasswordAdapter(),
		adapter.NewUUIDEntityID(),
		eventbus.NewAWSEventBus(),
	).Run(req, *appClient)

	if !resp.IsOk() {
		c.JSON(resp.GetError().GetHTTPCode(), gin.H{"error": resp.GetError().Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": resp.GetValue()})

}

func ChangePasswordController(c *gin.Context) {
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
