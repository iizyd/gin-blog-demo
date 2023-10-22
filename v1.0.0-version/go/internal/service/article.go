package service

import (
	"github.com/iamzhiyudong/xigua-blog/internal/model"
	"github.com/iamzhiyudong/xigua-blog/pkg/app"
)

type CreateArticleRequest struct {
	Title         string `form:"title" json:"title" binding:"required,min=1,max=100"`
	Desc          string `form:"desc" json:"desc" binding:"required,min=0,max=100"`
	CoverImageUrl string `form:"cover_image_url" json:"cover_image_url" binding:"omitempty,min=0,max=100"`
	Content       string `form:"content" json:"content" binding:"-"`
	Tag           []int  `form:"tag" json:"tag" binding:"-"`
	CreatedBy     string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	State         int    `form:"state,default=1" json:"state,default=1" binding:"omitempty,oneof=0 1"`
}

type ListArticleRequest struct {
	Title   string `form:"title" binding:"omitempty,min=1,max=100"`
	Desc    string `form:"desc" binding:"omitempty,min=0,max=100"`
	Content string `form:"content" binding:"-"`

	State int `form:"state,default=-1" binding:"omitempty,oneof=-1 0 1"`
}

type CountArticleRequest struct {
	Title   string `form:"title" binding:"omitempty,min=1,max=100"`
	Desc    string `form:"desc" binding:"omitempty,min=0,max=100"`
	Content string `form:"content" binding:"-"`

	State int `form:"state,default=-1" binding:"omitempty,oneof=-1 0 1"`
}

type GetArticleRequest struct {
	ID int `json:"id" form:"id" binding:"gt=0"`
}

type UpdateArticleRequest struct {
	ID            int    `json:"id" form:"id" binding:"gt=0"`
	Title         string `json:"title" form:"title" binding:"omitempty,min=1,max=100"`
	Desc          string `json:"desc" form:"desc" binding:"omitempty,min=0,max=100"`
	CoverImageUrl string `json:"cover_image_url" form:"cover_image_url" binding:"omitempty,min=0,max=100"`
	Tag           []int  `form:"tag" json:"tag" binding:"dive"`
	Content       string `json:"content" form:"content" binding:"-"`
	ModifiedBy    string `json:"modified_by" form:"modified_by" binding:"required,min=2,max=100"`
	State         int    `json:"state" form:"state" binding:"omitempty,oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID int `json:"id" form:"id" binding:"gt=0"`
}

func (s *Service) CreateArticle(param *CreateArticleRequest) error {
	return s.dao.CreateArticle(param.Title, param.Desc, param.CoverImageUrl, param.Content, param.CreatedBy, param.State, param.Tag)
}

func (s *Service) ListArticle(param *ListArticleRequest, pager *app.Pager) ([]*model.Article, error) {
	return s.dao.ListArticle(param.Title, param.Desc, param.Content, param.State, pager.Page, pager.PageSize)
}

func (s *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return s.dao.CountArticle(param.Title, param.Desc, param.Content, param.State)
}

func (s *Service) GetArticle(param *GetArticleRequest) ([]*model.Article, error) {
	return s.dao.GetArticle(param.ID)
}

func (s *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return s.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.CoverImageUrl, param.Content, param.ModifiedBy, param.State, param.Tag)
}

func (s *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return s.dao.DeleteArticle(param.ID)
}
