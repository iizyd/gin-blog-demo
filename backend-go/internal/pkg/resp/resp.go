package resp

import (
	"github.com/gin-gonic/gin"
)

func Resp(c *gin.Context, status int, msg string, data any, code int) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
