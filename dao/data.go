package dao

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	SpaceID uint `gorm:"index"`
}
