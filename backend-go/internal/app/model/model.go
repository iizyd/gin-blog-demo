package model

import (
	"backend-go/utils"
)

type BaseModel struct {
	ID        uint             `gorm:"primaryKey" json:"id"`
	CreatedAt *utils.LocalTime `json:"created_at"`
}

type Model struct {
	BaseModel
	ModifiedAt *utils.LocalTime `json:"modified_at"`
	// DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
