package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	conf AppConfig
	one  sync.Once
)

type AppConfig struct {
	Env      string         `mapstructure:"env"`
	Port     int            `mapstructure:"port"`
	Database DatabaseConfig `mapstructure:"database"`
	Swagger  SwaggerConfig  `mapstructure:"swagger"`
}

type DatabaseConfig struct {
	Url string `mapstructure:"url"`
}

type SwaggerConfig struct {
	HostUrl string `mapstructure:"host_url"`
}

func Load(path ...string) *AppConfig {
	one.Do(func() {
		appConfig := AppConfig{}

		filePath := os.Getenv("CONFIG_FILE_PATH")

		if filePath != "" {
			viper.SetConfigFile(filePath)
		} else if len(path) > 0 {
			p := path[0]
			viper.SetConfigFile(p)
		} else {
			viper.AddConfigPath("./config")
			viper.SetConfigName("config")
		}

		viper.SetConfigType("yaml")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("error occurs while reading the config. ", err)
		}

		if err := viper.Unmarshal(&appConfig); err != nil {
			log.Fatal("error occurs while unmarshalling the config. ", err)
		}
		conf = appConfig
	})
	return &conf
}
