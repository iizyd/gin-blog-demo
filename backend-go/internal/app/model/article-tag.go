package model

type ArticleTag struct {
	TagID     int `gorm:"primaryKey" json:"tag_id"`
	ArticleID int `gorm:"primaryKey" json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "article_tag"
}
