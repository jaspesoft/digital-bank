package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r := NewRouter()

	r.Use(cors.New(config))

	_ = r.Run(":8080")
}
