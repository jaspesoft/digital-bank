package server

import (
	"digital-bank/infrastructure/api/server/routes"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	routes.AccountRoute(r)

	return r
}
