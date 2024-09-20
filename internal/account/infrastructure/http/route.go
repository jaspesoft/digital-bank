package httpaccount

import (
	"digital-bank/infrastructure/api/middleware"
	"github.com/gin-gonic/gin"
)

func AccountRoute(r *gin.Engine) {
	account := r.Group("/account").Use(middleware.AuthMiddleware)
	{
		account.GET("/", func(context *gin.Context) {

		})
		//account.POST("/", CreateAccount)
		//account.PUT("/", UpdateAccount)
		//account.DELETE("/", DeleteAccount)
	}
}
