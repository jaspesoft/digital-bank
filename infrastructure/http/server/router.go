package server

import (
	"digital-bank/infrastructure/http/middleware"
	accounthttp "digital-bank/internal/account/infrastructure/http"
	systemhttp "digital-bank/internal/system/infrastructure/http"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware)

	accounthttp.AccountRoute(r)
	systemhttp.SystemRoute(r)

	return r
}
