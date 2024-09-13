package domain

import "time"

type Post struct {
	ID        int       `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	Title     string    `gorm:"column:title" json:"title"`
	Content   string    `gorm:"column:content" json:"content"`
	ViewCount int       `gorm:"column:view_count" json:"view_count"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
}

func (Post) TableName() string {
	return "post"
}
