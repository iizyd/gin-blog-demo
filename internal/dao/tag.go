package dao

import (
	"github.com/iamzhiyudong/xigua-blog/internal/model"
	"github.com/iamzhiyudong/xigua-blog/pkg/app"
)

func (d *Dao) CountTag(name string, state int8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state int8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state int8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state int8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
