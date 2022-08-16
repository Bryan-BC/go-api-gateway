package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	AuthSvcURL string `mapstructure:"AUTH_SVC_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Panicf("Error unmarshalling config file, %s", err)
	}

	return
}
