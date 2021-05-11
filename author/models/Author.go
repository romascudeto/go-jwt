package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password,omitempty" gorm:"-"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (b *Author) TableName() string {
	return "author"
}
