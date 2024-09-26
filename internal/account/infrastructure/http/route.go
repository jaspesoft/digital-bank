package httpaccount

import (
	"digital-bank/infrastructure/http/middleware"
	accountcontroller "digital-bank/internal/account/infrastructure/http/controllers"
	"github.com/gin-gonic/gin"
)

func AccountRoute(r *gin.Engine) {
	account := r.Group("/api/v1/account").Use(middleware.AuthMiddleware)
	account.POST("/application", accountcontroller.ApplicationAccountCompanyController)
}
