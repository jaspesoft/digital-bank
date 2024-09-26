package systemcontrollers

import (
	"digital-bank/infrastructure/adapter"
	systempersistence "digital-bank/internal/system/infrastructure/persistence"
	systemreq "digital-bank/internal/system/infrastructure/requests"
	systemusecase "digital-bank/internal/system/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OnboardingAppClient(ctx *gin.Context) {
	var appClientRequest systemreq.AppClientRequest

	if err := ctx.ShouldBindJSON(&appClientRequest); err != nil {
		validationErrors := adapter.FormatValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	if appClientRequest.Commissions == (systemreq.Commissions{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": adapter.ErrorResponse{
			Field:   "commissions",
			Message: "required",
		}})
		return
	}

	result := systemusecase.NewOnboardingAppClient(
		adapter.NewUUIDEntityID(),
		systempersistence.NewAppClientMongoRepository(),
		systempersistence.NewSystemParametersMongoRepository(),
	).Run(appClientRequest)

	if !result.IsOk() {
		ctx.JSON(result.GetError().GetHTTPCode(), gin.H{"message": result.GetError().Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": result.GetValue()})
}
