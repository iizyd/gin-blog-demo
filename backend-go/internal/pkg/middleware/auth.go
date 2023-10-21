package middleware

import (
	"backend-go/internal/pkg/jwt"
	"backend-go/internal/pkg/resp"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			resp.Resp(c, 200, "未登录", nil, -1)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			resp.Resp(c, 400, "请求头token格式错误", nil, -1)
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			fmt.Println(err.Error())
			resp.Resp(c, 200, "token无效", nil, -1)
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		//后续的处理函数可以通过c.Get("username")来获取请求的用户信息
		c.Next()
	}

}
