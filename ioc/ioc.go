package ioc

import (
	"github.com/spf13/viper"
	"github.com/universalmacro/common/singleton"
	"gorm.io/gorm"
)

var GetConfig = singleton.EagerSingleton(func() *viper.Viper {
	return viper.New()
})

var GetDB = singleton.EagerSingleton(func() *gorm.DB {
	return nil
	// return dao.NewConnection()
})
