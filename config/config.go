package config

import (
	"log"

	"github.com/spf13/viper"
)

// Viper config
func ViperConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
}
