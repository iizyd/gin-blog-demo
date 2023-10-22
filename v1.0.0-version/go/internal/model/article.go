package model

import (
	"github.com/iizyd/xigua-blog/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `form:"title" json:"title"`
	Desc          string `form:"desc" json:"desc"`
	Content       string `form:"content" json:"content"`
	CoverImageUrl string `form:"cover_image_url" json:"cover_image_url"`
	State         int    `form:"state" json:"state"`
	Tag           []Tag  `gorm:"many2many:blog_article_tag" json:"tag"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	err := db.Omit("Tag").Create(&a).Error
	if err != nil {
		return err
	}
	return db.Model(&a).Preload("Tag").Association("Tag").Replace(a.Tag).Error
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error

	db = db.Preload("Tag")

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if a.Title != "" {
		db = db.Where("`title` LIKE ?", "%"+a.Title+"%")
	}
	if a.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+a.Desc+"%")
	}
	if a.Content != "" {
		db = db.Where("`content` LIKE ?", "%"+a.Content+"%")
	}

	if a.State != -1 {
		db = db.Where("state = ?", a.State)
	}

	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int

	if a.Title != "" {
		db = db.Where("`title` LIKE ?", "%"+a.Title+"%")
	}
	if a.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+a.Desc+"%")
	}
	if a.Content != "" {
		db = db.Where("`content` LIKE ?", "%"+a.Content+"%")
	}

	if a.State != -1 {
		db = db.Where("'state' = ?", a.State)
	}

	if err := db.Model(&a).Where("`is_del` = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) Get(db *gorm.DB) ([]*Article, error) {
	var article []*Article

	if err := db.Preload("Tag").Where("`id` = ? AND is_del = ?", a.ID, 0).First(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (a Article) Update(db *gorm.DB, values any) error {
	db = db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0)
	if err := db.Omit("Tag").Updates(values).Error; err != nil {
		return err
	}
	if err := db.Model(&a).Preload("Tag").Association("Tag").Replace(a.Tag).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}
