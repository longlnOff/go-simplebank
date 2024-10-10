package initialize

import (
	"github.com/longln/go-simplebank/global"
	"github.com/spf13/viper"
)


func LoadConfig() {
	config := viper.New()

	config.AddConfigPath("local")

	config.SetConfigName("config")

	config.SetConfigType("yaml")


	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := config.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}