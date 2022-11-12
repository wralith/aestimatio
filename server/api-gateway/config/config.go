package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Services ServicesConfig
}

type ServerConfig struct {
	Port string
}

type LoggerConfig struct {
	Pretty bool
	Level  string
}

type ServicesConfig struct {
	Auth string
}

func Get() *Config {
	if err := initConfig(); err != nil {
		log.Panic().Err(err)
	}
	return &Config{
		Server: ServerConfig{
			Port: viper.GetString("server.port"),
		},
		Logger: LoggerConfig{
			Pretty: viper.GetBool("logger.pretty"),
			Level:  viper.GetString("logger.level"),
		},
		Services: ServicesConfig{
			Auth: viper.GetString("services.auth"),
		},
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	return viper.ReadInConfig()
}
