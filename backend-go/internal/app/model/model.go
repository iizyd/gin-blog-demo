package model

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"create_at"`
}

type Model struct {
	BaseModel
	ModifiedAt time.Time `json:"modified_at"`
	// DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
