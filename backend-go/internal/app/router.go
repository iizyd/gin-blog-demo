package app

import (
	"backend-go/internal/app/handlers"
	"backend-go/internal/pkg/config"
	"backend-go/internal/pkg/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	r.StaticFS("/file", http.Dir(config.Config.App.UploadSavePath))

	api := r.Group("/api")

	upload(api)
	userMgt(api)
	tagMgt(api)

	r.Run(config.Config.Server.HttpPort)
}

func userMgt(api *gin.RouterGroup) {
	api.POST("/login", handlers.Login)
}

func upload(api *gin.RouterGroup) {
	api.POST("/upload", handlers.Upload)
}

func tagMgt(api *gin.RouterGroup) {
	t := api.Group("/tag", middleware.JWTAuthMiddleware())

	t.GET("", handlers.GetAllTags)
}
