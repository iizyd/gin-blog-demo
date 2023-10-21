package handlers

import (
	"backend-go/internal/app/model"
	"backend-go/internal/pkg/bind"
	"backend-go/internal/pkg/logger"
	"backend-go/internal/pkg/resp"
	"backend-go/utils"

	"github.com/gin-gonic/gin"
)

func GetAllTags(c *gin.Context) {
	tag := model.GetTags()
	resp.Resp(c, 200, "", tag, 0)
}

func GetTag(c *gin.Context) {
	id, err := utils.StrTo(c.Param("id")).Int()
	if err != nil {
		logger.Errorf("路径id为数字")
		resp.Resp(c, 400, "路径id为数字", nil, -1)
		return
	}

	tag := model.GetTagById(id)
	if tag.ID == 0 {
		resp.Resp(c, 200, "ID不存在", nil, -1)
		return
	}

	resp.Resp(c, 200, "", tag, 0)
}

type CreateTagReq struct {
	Name string `json:"name"  valid:"required~tag名不能为空"`
}

func CreateTag(c *gin.Context) {
	var createTagReq CreateTagReq
	if err := bind.BindAndValid(c, &createTagReq); err != nil {
		resp.Resp(c, 400, "参数错误:"+err.Error(), nil, -1)
		return
	}

	tag, err := model.CreateTag(createTagReq.Name)
	if err != nil {
		logger.Errorf("tag创建失败", err.Error())
		resp.Resp(c, 400, "tag创建失败", nil, -1)
		return
	}

	resp.Resp(c, 200, "", tag, 0)
}

func UpdateTag(c *gin.Context) {
	id, err := utils.StrTo(c.Param("id")).Int()
	if err != nil {
		logger.Errorf("路径id为数字")
		resp.Resp(c, 400, "路径id为数字", nil, -1)
		return
	}

	var createTagReq CreateTagReq
	if err := bind.BindAndValid(c, &createTagReq); err != nil {
		resp.Resp(c, 400, "参数错误:"+err.Error(), nil, -1)
		return
	}

	tag := model.GetTagById(id)
	if tag.ID == 0 {
		resp.Resp(c, 200, "ID不存在", nil, -1)
		return
	}

	if err := model.UpdateTag(id, createTagReq.Name); err != nil {
		logger.Errorf("tag更新失败", err.Error())
		resp.Resp(c, 200, "更新失败", nil, -1)
		return
	}

	resp.Resp(c, 200, "", nil, 0)
}

func DeleteTag(c *gin.Context) {
	id, _ := utils.StrTo(c.Param("id")).Int()
	tag := model.GetTagById(id)
	if tag.ID == 0 {
		resp.Resp(c, 200, "ID不存在", nil, -1)
		return
	}

	if err := model.DeleteTag(id); err != nil {
		logger.Errorf("tag删除失败", err.Error())
		resp.Resp(c, 500, "tag删除失败", nil, -1)
		return
	}

	resp.Resp(c, 200, "", nil, 0)
}
