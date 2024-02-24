package config

import (
	"github.com/spf13/viper"
	"log"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading the config.")
	}

	err := viper.Unmarshal(&configs)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v	", err)
	}
}
