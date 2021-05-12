package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int64     `json:"author_id"`
	Author    Author    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt gorm.DeletedAt
}

func (b *Article) TableName() string {
	return "article"
}
