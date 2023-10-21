package handlers

import (
	"backend-go/internal/app/model"
	"backend-go/internal/pkg/bind"
	"backend-go/internal/pkg/logger"
	"backend-go/internal/pkg/page"
	"backend-go/internal/pkg/resp"
	"backend-go/utils"

	"github.com/gin-gonic/gin"
)

type GetAllArticlesReq struct {
	page.Page
	Published bool `form:"published" json:"published"`
}
type GetAllResp struct {
	Total int             `json:"total"`
	Data  []model.Article `json:"data"`
}

func GetAllArticles(c *gin.Context) {
	var query GetAllArticlesReq
	if err := bind.BindAndValid(c, &query); err != nil {
		resp.Resp(c, 200, err.Error(), nil, 0)
		return
	}

	articles := model.GetAllArticles(query.Page.GetLimit(), query.Page.GetOffset())
	count := model.CountArticles(0)

	resp.Resp(c, 200, "", &GetAllResp{
		Total: int(count),
		Data:  articles,
	}, 0)
}

func GetArticleById(c *gin.Context) {
	id, _ := utils.StrTo(c.Param("id")).Int()

	article := model.GetArticleById(id)
	if article.ID == 0 {
		resp.Resp(c, 200, "ID不存在", nil, 0)
		return
	}

	resp.Resp(c, 200, "", article, 0)
}

type CreateArticleReq struct {
	Title         string `json:"title" valid:"required~title必填"`
	Description   string `json:"description" valid:"required~description必填"`
	CoverImageUrl string `json:"cover_image_url" valid:"required~cover_image_url必填"`
	Content       string `json:"content" valid:"required~content必填"`
	Published     bool   `json:"published" valid:"-"`
	Tags          []int  `json:"tags" valid:"required~标签数组必填"`
}

func CreateArticle(c *gin.Context) {
	var queryForm CreateArticleReq
	if err := bind.BindAndValid(c, &queryForm); err != nil {
		resp.Resp(c, 400, "参数格式错误: "+err.Error(), nil, -1)
		return
	}

	article, err := model.CreateArticle(queryForm.Title, queryForm.Description, queryForm.Content, queryForm.CoverImageUrl, queryForm.Published, queryForm.Tags)
	if article.ID == 0 || err != nil {
		logger.Errorf("文章创建失败：", err.Error())
		resp.Resp(c, 500, "文章创建失败：", nil, -1)
		return
	}
	resp.Resp(c, 200, "", nil, 0)
}

func UpdateArticle(c *gin.Context) {
	var queryForm CreateArticleReq
	if err := bind.BindAndValid(c, &queryForm); err != nil {
		resp.Resp(c, 400, "参数格式错误: "+err.Error(), nil, -1)
		return
	}

	id, _ := utils.StrTo(c.Param("id")).Int()
	article := model.GetArticleById(id)
	if article.ID == 0 {
		resp.Resp(c, 200, "ID不存在", nil, -1)
		return
	}

	article, err := model.UpdateArticle(id, queryForm.Title, queryForm.Description, queryForm.Content, queryForm.CoverImageUrl, queryForm.Published, queryForm.Tags)
	if article.ID == 0 || err != nil {
		logger.Errorf("文章修改失败：", err.Error())
		resp.Resp(c, 500, "文章修改失败：", nil, -1)
		return
	}
	resp.Resp(c, 200, "", nil, 0)
}

func DeleteArticle(c *gin.Context) {
	id, _ := utils.StrTo(c.Param("id")).Int()
	article := model.GetArticleById(id)
	if article.ID == 0 {
		resp.Resp(c, 200, "ID不存在", nil, -1)
		return
	}

	if err := model.DeleteArticle(id); err != nil {
		logger.Errorf("文章删除失败", err.Error())
		resp.Resp(c, 500, "文章删除失败", nil, -1)
		return
	}

	resp.Resp(c, 200, "", nil, 0)
}

func GetAllPublishedArticle(c *gin.Context) {
	var query GetAllArticlesReq
	if err := bind.BindAndValid(c, &query); err != nil {
		resp.Resp(c, 200, err.Error(), nil, 0)
		return
	}

	articles := model.GetAllPublishedArticles(query.Page.GetLimit(), query.Page.GetOffset())
	count := model.CountArticles(1)

	resp.Resp(c, 200, "", &GetAllResp{
		Total: int(count),
		Data:  articles,
	}, 0)
}
