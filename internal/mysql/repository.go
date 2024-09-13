package mysql

import "gorm.io/gorm"

type UserRepo struct {
	DB *gorm.DB
}

type PostRepo struct {
	DB *gorm.DB
}
