package settings

import (
	"log"

	"github.com/spf13/viper"
)

var (
	AppSettings Settings
)

func ReadSettings() Settings {
	var option Settings
	viper.AddConfigPath(".")
	viper.SetConfigName("settings")
	viper.SetConfigType("json")

	if err:=viper.ReadInConfig(); err!=nil{
		log.Fatalln(err)
	}

	if err:=viper.Unmarshal(&option); err!=nil{
		log.Fatalln(err)
	}

	return option
}