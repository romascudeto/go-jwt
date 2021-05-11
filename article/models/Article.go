package models

import "time"

type Article struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Article) TableName() string {
	return "article"
}
