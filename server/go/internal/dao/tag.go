package dao

import (
	"github.com/iamzhiyudong/xigua-blog/internal/model"
	"github.com/iamzhiyudong/xigua-blog/pkg/app"
)

func (d *Dao) CountTag(name string, state int) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state int, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state int, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id int, name string, state int, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	if state != -1 {
		values["state"] = state
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id int) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
