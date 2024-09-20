package server

import (
	"digital-bank/infrastructure/http/middleware"
	httpaccount "digital-bank/internal/account/infrastructure/http"
	httpsystem "digital-bank/internal/system/infrastructure/http"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware)

	httpaccount.AccountRoute(r)
	httpsystem.SystemRoute(r)

	return r
}
