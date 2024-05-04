package config

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Logger LoggerConfig
}

type ServerConfig struct {
	InternalPort      string
	RunMode           string
	StaticContentPath string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))

	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error in load config %v", err)
	}

	cfg, err := ParseConfig(v)

	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	if env == "production" {
		return "/config/config-production"
	} else {
		return "./config/config-development"
	}
}
