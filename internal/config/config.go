package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")    // name of file (without extension)
	viper.SetConfigType("yaml")      // type of the config file
	viper.AddConfigPath("./configs") // path to look for the config file

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &cfg
}
