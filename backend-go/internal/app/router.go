package app

import (
	"backend-go/internal/app/handlers"
	"backend-go/internal/pkg/config"
	"backend-go/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	api := r.Group("/api")

	userMgt(api)
	tagMgt(api)

	r.Run(config.Config.Server.HttpPort)
}

func userMgt(api *gin.RouterGroup) {
	api.POST("/login", handlers.Login)
}

func tagMgt(api *gin.RouterGroup) {
	t := api.Group("/tag", middleware.JWTAuthMiddleware())

	t.GET("", handlers.GetAllTags)
}
