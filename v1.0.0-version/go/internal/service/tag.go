package service

import (
	"github.com/iizyd/xigua-blog/internal/model"
	"github.com/iizyd/xigua-blog/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int    `form:"state" binding:"omitempty,oneof=-1 0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int    `form:"state,default=-1" binding:"omitempty,oneof=-1 0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" json:"name" binding:"required,min=2,max=100"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=2,max=100"`
	State     int    `form:"state,default=1" json:"state,default=1" binding:"omitempty,oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         int    `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      int    `form:"state,default=-1" binding:"omitempty,oneof=-1 0 1"`
	ModifiedBy string `json:"modified_by" form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteTagRequest struct {
	ID int `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
