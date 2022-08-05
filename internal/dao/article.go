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

func (dao *Dao) GetArticle(id int) ([]*model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}}

	return article.Get(dao.engine)
}

func (dao *Dao) UpdateArticle(id int, title, desc, cover_image_url, content, modified_by string, state int, tag []int) error {
	article := model.Article{
		Model: &model.Model{
			ID:         id,
			ModifiedBy: modified_by,
		},
		Title:         title,
		Desc:          desc,
		CoverImageUrl: cover_image_url,
		Content:       content,
		State:         state,
	}

	values := map[string]interface{}{
		"modified_by":     modified_by,
		"cover_image_url": cover_image_url,
	}
	if title != "" {
		values["title"] = title
	}
	if desc != "" {
		values["desc"] = desc
	}
	if content != "" {
		values["content"] = content
	}
	if state != -1 {
		values["state"] = state
	}
	if len(tag) != 0 {
		var tags []model.Tag
		for _, val := range tag {
			tags = append(tags, model.Tag{Model: &model.Model{ID: val}})
		}
		article.Tag = tags
	}

	return article.Update(dao.engine, values)
}

func (dao *Dao) DeleteArticle(id int) error {
	article := model.Article{
		Model: &model.Model{ID: id},
	}

	return article.Delete(dao.engine)
}
