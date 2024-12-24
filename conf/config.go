package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	// ===========================================
	// viper设置配置文件相关信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("Load Config Error %v", err.Error()))
	}

}
