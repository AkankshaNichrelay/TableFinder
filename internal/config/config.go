package config

import (
	"fmt"

	"github.com/AkankshaNichrelay/TableFinder/internal/database"
	"github.com/AkankshaNichrelay/TableFinder/internal/redis"
	"github.com/spf13/viper"
)

// Config contains all configs of the application
type Config struct {
	Redis redis.Config
	DB    database.Config
}

// New loads all configs from specified yaml file
func New(filename string) (*Config, error) {
	config := Config{}

	viper.SetConfigName(filename)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config values, %v", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config values, %v", err)
	}

	fmt.Println("Config Values Loaded", config)
	return &config, nil
}

// NewRedisConfig return Redis client config
func NewRedisConfig(config *Config) *redis.Config {
	cfg := redis.Config{}
	return &cfg
}

// NewDBConfig returns Database config
func NewDBConfig(config *Config) *database.Config {
	cfg := database.Config{}
	return &cfg
}
