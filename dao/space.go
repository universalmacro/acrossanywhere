package dao

import "gorm.io/gorm"

type Space struct {
	gorm.Model
	UserID uint   `gorm:"index"`
	Name   string `gorm:"type:varchar(100)"`
}
