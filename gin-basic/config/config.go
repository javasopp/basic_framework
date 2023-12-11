package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config AppConfig

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 将配置文件内容解析到结构体中
	marshalError := viper.Unmarshal(&Config)
	if marshalError != nil {
		panic(fmt.Errorf("Unable to unmarshal config: %s \n", marshalError))
	}
}
