package app

import (
	"backend-go/internal/app/model"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	api := r.Group("/api")

	userMgt(api)

	r.Run()
}

func userMgt(api *gin.RouterGroup) {
	api.POST("/login", model.Login)
}
