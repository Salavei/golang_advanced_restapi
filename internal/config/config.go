package config

import (
	"github.com/Salavei/golang_advanced_restapi/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-required:"port"`
		BindIP string `yaml:"bind_ip" env-required:"127.0.0.1"`
		Port   string `yaml:"port" env-required:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	PostgreSQl struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Database    string `yaml:"database"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		MaxAttempts int    `yaml:"max_attempts"`
	} `json:"postgresql"`

	MongoDB struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Database   string `yaml:"database"`
		AuthDB     string `yaml:"auth_db"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
}

var instance *Config

var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
