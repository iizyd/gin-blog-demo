package handlers

import (
	"backend-go/internal/app/model"
	"backend-go/internal/pkg/bind"
	"backend-go/internal/pkg/jwt"
	"backend-go/internal/pkg/resp"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	UserName string `gorm:"username" json:"username" valid:"required~姓名不能为空"`
	Password string `gorm:"password" json:"password" valid:"required~密码不能为空,minstringlength(6)~密码至少6位"`
}
type LoginRes struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	loginReq := LoginReq{}
	if err := bind.BindAndValid(c, &loginReq); err != nil {
		resp.Resp(c, 200, err.Error(), nil, 0)
		return
	}

	user := model.GetUserById(loginReq.UserName)
	if user.ID == 0 {
		resp.Resp(c, 200, "用户名或密码错误", nil, 0)
		return
	}

	token, err := jwt.GenToken(user.UserName)
	if err != nil {
		resp.Resp(c, 200, err.Error(), nil, 0)
		return
	}

	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	resp.Resp(c, 200, "", LoginRes{Token: token}, 0)
}
