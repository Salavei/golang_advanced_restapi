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
