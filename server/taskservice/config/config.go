package config

import "github.com/spf13/viper"

type Config struct {
	Port string
}

func Get() *Config {
	initConfig()
	return &Config{
		Port: viper.GetString("server.port"),
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("error while reading configs")
	}
}
