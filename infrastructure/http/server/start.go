package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r := NewRouter()

	r.Use(cors.New(config))

	fmt.Println("Server is running on port 8080")

	_ = r.Run(":8080")

}
