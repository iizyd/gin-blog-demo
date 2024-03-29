package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iizyd/xigua-blog/global"
	"github.com/iizyd/xigua-blog/internal/middleware"
	"github.com/iizyd/xigua-blog/internal/routers/api"
	v1 "github.com/iizyd/xigua-blog/internal/routers/api/v1"

	"github.com/gin-contrib/cors"
	_ "github.com/iizyd/xigua-blog/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	r.Use(middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := v1.NewUpload()

	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	apiv1.Use(middleware.AccessLog())
	{
		// 标签
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		// 文章
		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
