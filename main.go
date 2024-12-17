package main

import (
	"github.com/DaHuangQwQ/web/ioc"

	"github.com/spf13/viper"
)

func main() {
	InitViper()
	app := ioc.InitApp()
	err := app.Server.Start(":8081")
	panic(err)
}

func InitViper() {
	viper.SetConfigFile("./config/dev.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
