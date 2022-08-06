package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("user_name = ? AND password = ? AND is_del = ?", a.UserName, a.Password, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
