package service

import (
	"github.com/iamzhiyudong/xigua-blog/internal/model"
	"github.com/iamzhiyudong/xigua-blog/pkg/app"
)

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=1,max=100"`
	Desc          string `form:"desc" binding:"required,min=0,max=100"`
	CoverImageUrl string `form:"cover_image_url" binding:"omitempty,min=0,max=100"`
	Content       string `form:"content" binding:"-"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State         int    `form:"state,default=1" binding:"omitempty,oneof=0 1"`
}

type ListArticleRequest struct {
	Title   string `form:"title" binding:"min=1,max=100"`
	Desc    string `form:"desc" binding:"min=0,max=100"`
	Content string `form:"content" binding:"-"`

	State int `form:"state,default=-1" binding:"omitempty,oneof=-1 0 1"`
}

type CountArticleRequest struct {
	Title   string `form:"title" binding:"min=1,max=100"`
	Desc    string `form:"desc" binding:"min=0,max=100"`
	Content string `form:"content" binding:"-"`

	State int `form:"state,default=-1" binding:"omitempty,oneof=-1 0 1"`
}

func (s *Service) CreateArticle(param *CreateArticleRequest) error {
	return s.dao.CreateArticle(param.Title, param.Desc, param.CoverImageUrl, param.Content, param.CreatedBy, param.State)
}

func (s *Service) ListArticle(param *ListArticleRequest, pager *app.Pager) ([]*model.Article, error) {
	return s.dao.ListArticle(param.Title, param.Desc, param.Content, param.State, pager.Page, pager.PageSize)
}

func (s *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return s.dao.CountArticle(param.Title, param.Desc, param.Content, param.State)
}
