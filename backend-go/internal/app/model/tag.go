package model

import "backend-go/internal/pkg/db"

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"name" json:"name"`
}

func (Tag) TableName() string {
	return "tag"
}

func GetTags() []Tag {
	var tag []Tag
	db.DB.Find(&tag)
	return tag
}

func GetTagById(id int) Tag {
	var tag Tag
	db.DB.Where("id = ?", id).First(&tag)
	return tag
}

func CreateTag(name string) (Tag, error) {
	tag := Tag{Name: name}

	res := db.DB.Create(&tag)
	if res.Error != nil {
		return tag, res.Error
	}
	return tag, nil
}

func UpdateTag(id int, name string) error {
	tag := Tag{ID: uint(id), Name: name}

	res := db.DB.Save(&tag)
	return res.Error
}

func DeleteTag(id int) error {
	tag := Tag{ID: uint(id)}

	res := db.DB.Delete(&tag)
	return res.Error
}
