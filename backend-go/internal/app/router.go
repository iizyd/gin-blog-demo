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
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Token", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	r.StaticFS("/file", http.Dir(config.Config.App.UploadSavePath))

	api := r.Group("/api")

	upload(api)
	userMgt(api)
	tagMgt(api)
	articleMgt(api)

	r.Run(config.Config.Server.HttpPort)
}

func userMgt(api *gin.RouterGroup) {
	api.POST("/login", handlers.Login)
}

func upload(api *gin.RouterGroup) {
	u := api.Group("/upload", middleware.JWTAuthMiddleware())
	u.POST("", handlers.Upload)
}

func tagMgt(api *gin.RouterGroup) {
	t := api.Group("/tag", middleware.JWTAuthMiddleware())
	// t := api.Group("/tag")

	t.GET("", handlers.GetAllTags)
	t.GET("/:id", handlers.GetTag)
	t.POST("", handlers.CreateTag)
	t.DELETE("/:id", handlers.DeleteTag)
	t.PUT("/:id", handlers.UpdateTag)
}

func articleMgt(api *gin.RouterGroup) {
	t := api.Group("/article")
	// t := api.Group("/article")

	t.GET("", middleware.JWTAuthMiddleware(), handlers.GetAllArticles)
	t.GET("/:id", handlers.GetArticleById)
	t.POST("", middleware.JWTAuthMiddleware(), handlers.CreateArticle)
	t.DELETE("/:id", middleware.JWTAuthMiddleware(), handlers.DeleteArticle)
	t.PUT("/:id", middleware.JWTAuthMiddleware(), handlers.UpdateArticle)
	t.GET("/published", handlers.GetAllPublishedArticle)
}
