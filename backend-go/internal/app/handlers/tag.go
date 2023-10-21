package handlers

import (
	"backend-go/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

func GetAllTags(c *gin.Context) {
	resp.Resp(c, 200, "", nil, 0)
}
