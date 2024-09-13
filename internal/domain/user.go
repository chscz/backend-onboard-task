package domain

import (
	"time"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Email     string    `gorm:"column:email;unique"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
}

func (User) TableName() string {
	return "user"
}
