package model

import (
	"backend-go/internal/pkg/db"
)

type User struct {
	BaseModel
	UserName string `gorm:"column:username" json:"username"`
	Password string `gorm:"password" json:"password"`
}

func (User) TableName() string {
	return "user"
}

func GetUserById(name string) User {
	var user User
	db.DB.Where("username = ?", name).First(&user)
	return user
}
