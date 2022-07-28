package model

import (
	"github.com/iamzhiyudong/xigua-blog/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         int    `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error

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

	if err := db.Model(&a).Where("'is_del' = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (a Article) Get(db *gorm.DB) ([]*Article, error) {
	var article []*Article

	if err := db.Where("`id` = ? AND is_del = ?", a.ID, 0).First(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (a Article) Update(db *gorm.DB, values any) error {
	if err := db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
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
