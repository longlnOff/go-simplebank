package initialize

import (
	"github.com/longln/go-simplebank/global"
	"github.com/spf13/viper"
)


func LoadConfig() {
	viper := viper.New()

	viper.AddConfigPath("/home/longln/SourceCode/github.com/longln/go-simplebank/local")

	viper.SetConfigName("config")

	viper.SetConfigType("yaml")


	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}