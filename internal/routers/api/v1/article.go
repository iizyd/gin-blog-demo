package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iamzhiyudong/xigua-blog/global"
	"github.com/iamzhiyudong/xigua-blog/internal/service"
	"github.com/iamzhiyudong/xigua-blog/pkg/app"
	"github.com/iamzhiyudong/xigua-blog/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// @Summary 获取文章列表
// @Produce  json
// @Tags 文章
// @Param title query string false "文章名称" minlength(1) maxlength(100)
// @Param desc query string false "文章描述" minlength(0) maxlength(255)
// @Param content query string false "文章内容"
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (a Article) List(c *gin.Context) {
	params := service.ListArticleRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	total, err := svc.CountArticle(&service.CountArticleRequest{
		Title:   params.Title,
		Desc:    params.Desc,
		Content: params.Content,
		State:   params.State,
	})
	if err != nil {
		global.Logger.Errorf("svc.CountArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountArticleFail)
		return
	}

	articles, err := svc.ListArticle(&params, &pager)
	if err != nil {
		global.Logger.Errorf("svc.ListArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}

	response.ToResponseList(articles, total)
}

// @Summary 新增文章
// @Produce  json
// @Tags 文章
// @Param title body string true "文章名称" minlength(1) maxlength(100)
// @Param desc body string true "文章描述" minlength(0) maxlength(255)
// @Param cover_image_url body string false "封面 url" minlength(0) maxlength(255)
// @Param content body string false "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (a Article) Create(c *gin.Context) {
	params := service.CreateArticleRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.CreateArticle(&params)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
