package httpaccount

import "github.com/gin-gonic/gin"

func AccountRoute(r *gin.Engine) {
	account := r.Group("/account")
	{
		account.GET("/", func(context *gin.Context) {

		})
		//account.POST("/", CreateAccount)
		//account.PUT("/", UpdateAccount)
		//account.DELETE("/", DeleteAccount)
	}
}
