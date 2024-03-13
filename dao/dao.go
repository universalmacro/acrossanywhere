package dao

import (
	"github.com/universalmacro/common/auth"
	"github.com/universalmacro/common/dao/data"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	*data.PhoneNumber
	*auth.Password
}

func (u *User) Spaces() []Space {
	return nil
}
