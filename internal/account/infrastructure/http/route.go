package accounthttp

import (
	"digital-bank/infrastructure/http/middleware"
	accountcontroller "digital-bank/internal/account/infrastructure/http/controllers"
	"github.com/gin-gonic/gin"
)

func AccountRoute(r *gin.Engine) {
	g := r.Group("/api/v1/account")
	g.Use(middleware.CompanyClientAuthMiddleware)
	{
		g.POST("/application", accountcontroller.ApplicationAccountCompanyController)
		g.POST("/signin", accountcontroller.AccountUserRegisterController)
		g.PUT("/change-password", accountcontroller.ChangePasswordController)
	}

}
