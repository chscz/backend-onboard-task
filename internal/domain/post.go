package domain

import "time"

type Post struct {
	ID        int       `gorm:"column:id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UserID    string    `gorm:"column:user_id"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	ViewCount int       `gorm:"column:view_count"`
	User      User      `gorm:"foreignKey:UserID"`
}

func (Post) TableName() string {
	return "post"
}
