package model

import (
	"github.com/iizyd/xigua-blog/pkg/app"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name    string    `json:"name"`
	State   int       `json:"state"`
	Article []Article `gorm:"many2many:blog_article_tag" json:"article"`
}

type TagSwagger struct {
	List []*Tag
	*app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if t.State != -1 {
		db = db.Where("state = ?", t.State)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	if t.State != -1 {
		db = db.Where("state = ?", t.State)
	}
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	var err error

	db = db.Where("id = ? AND is_del = ?", t.Model.ID, 0)
	if err = db.Delete(&t).Error; err != nil {
		return err
	}

	return db.Model(&t).Association("Article").Clear().Error
}
