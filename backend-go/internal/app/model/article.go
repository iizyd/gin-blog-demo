package model

import (
	"backend-go/internal/pkg/db"
	"backend-go/utils"
	"fmt"
	"time"
)

type Article struct {
	Model
	Title         string `gorm:"title" json:"title"`
	Description   string `gorm:"description" json:"description"`
	Content       string `gorm:"content" json:"content"`
	CoverImageUrl string `gorm:"column:cover_image_url" json:"cover_image_url"`
	Published     bool   `gorm:"published" json:"published"`

	Tag []Tag `gorm:"many2many:article_tag" json:"tags"`
}

func (Article) TableName() string {
	return "article"
}

func CountArticles(published int) int64 {
	var count int64
	if published == 0 {
		db.DB.Find(&Article{}).Count(&count)
	} else {
		db.DB.Find(&Article{}).Where("published = ?", published != 0).Count(&count)
	}
	return count
}

func GetAllArticles(limit int, offset int) []Article {
	var articles []Article
	db.DB.Model(&Article{}).Preload("Tag").Limit(limit).Offset(offset).Find(&articles)
	return articles
}

func GetAllPublishedArticles(limit int, offset int) []Article {
	var articles []Article
	db.DB.Model(&Article{}).Where("published = ?", true).Preload("Tag").Limit(limit).Offset(offset).Find(&articles)
	return articles
}

func GetArticleById(id int) Article {
	var article Article
	db.DB.Model(&Article{}).Preload("Tag").Where("id = ?", id).First(&article)
	return article
}

func CreateArticle(title string, description string, content string, cover_img_url string, published bool, tags []int) (Article, error) {
	article := Article{
		Title:         title,
		Description:   description,
		Content:       content,
		CoverImageUrl: cover_img_url,
		Published:     published,
	}

	var existingTags []Tag
	db.DB.Where("id IN (?)", tags).Find(&existingTags)

	if len(existingTags) > 0 {
		var list []Tag
		for _, v := range existingTags {
			list = append(list, Tag{ID: uint(v.ID)})
		}
		article.Tag = list
	}

	res := db.DB.Omit("created_at", "modified_at").Model(&Article{}).Preload("Tag").Create(&article)
	if res.Error != nil {
		return article, res.Error
	}
	return article, nil
}

func DeleteArticle(id int) error {
	article := Article{
		Model: Model{
			BaseModel: BaseModel{ID: uint(id)},
		},
	}
	res := db.DB.Model(&Article{}).Preload("Tag").Where("id = ?", id).Delete(&article)
	return res.Error
}

func UpdateArticle(id int, title string, description string, content string, cover_img_url string, published bool, tags []int) (Article, error) {
	currentTime := utils.LocalTime(time.Now())
	article := Article{
		Model: Model{
			BaseModel: BaseModel{
				ID: uint(id),
			},
			ModifiedAt: &currentTime,
		},
		Title:         title,
		Description:   description,
		Content:       content,
		CoverImageUrl: cover_img_url,
		Published:     published,
	}

	var existingTags []Tag
	db.DB.Where("id IN (?)", tags).Find(&existingTags)

	if len(existingTags) > 0 {
		var list []Tag
		for _, v := range existingTags {
			list = append(list, Tag{ID: uint(v.ID)})
		}
		article.Tag = list

		db.DB.Model(&article).Association("Tag").Replace(existingTags)
	}

	fmt.Println(article)

	res := db.DB.Model(&Article{}).Omit("created_at").Where("id = ?", article.ID).Preload("Tag").Save(&article)
	return article, res.Error
}
