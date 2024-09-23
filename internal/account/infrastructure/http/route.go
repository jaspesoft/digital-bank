package httpaccount

import (
	accountcontroller "digital-bank/internal/account/infrastructure/http/controllers"
	"github.com/gin-gonic/gin"
)

func AccountRoute(r *gin.Engine) {
	//account := r.Group("/account").Use(middleware.AuthMiddleware)
	//{
	//	account.GET("/", func(context *gin.Context) {
	//
	//	})
	//	//account.POST("/", CreateAccount)
	//	//account.PUT("/", UpdateAccount)
	//	//account.DELETE("/", DeleteAccount)
	//}

	account := r.Group("/api/v1/account")
	account.POST("/application", accountcontroller.ApplicationAccountCompanyController)
}
