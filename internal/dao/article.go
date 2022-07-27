package dao

import (
	"github.com/iamzhiyudong/xigua-blog/internal/model"
	"github.com/iamzhiyudong/xigua-blog/pkg/app"
)

func (dao *Dao) CreateArticle(title string, desc string, cover_image_url string, content string, createdBy string, state int) error {
	article := model.Article{
		Title:         title,
		Desc:          desc,
		CoverImageUrl: cover_image_url,
		Content:       content,
		State:         state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}
	return article.Create(dao.engine)
}

func (dao *Dao) ListArticle(title string, desc string, content string, state int, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{
		Title:   title,
		Desc:    desc,
		Content: content,
		State:   state,
	}

	pageOffset := app.GetPageOffset(page, pageSize)

	return article.List(dao.engine, pageOffset, pageSize)
}

func (dao *Dao) CountArticle(title string, desc string, content string, state int) (int, error) {
	article := model.Article{
		Title:   title,
		Desc:    desc,
		Content: content,
		State:   state,
	}

	return article.Count(dao.engine)
}
