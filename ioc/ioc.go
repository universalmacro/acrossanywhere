package ioc

import (
	"github.com/spf13/viper"
	"github.com/universalmacro/common/singleton"
)

var GetConfig = singleton.EagerSingleton(func() *viper.Viper {
	return viper.New()
})
