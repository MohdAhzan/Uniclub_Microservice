package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PORT          string `mapstructure:"PORT"`
	UserSvcUrl    string `mapstructure:"USER_SVC_URL"`
  ADMINSECRET   string `mapstructure:"ADMIN_SECRET"`
  USERACCESSSECRET   string `mapstructure:"USER_ACCESS_SECRET"`
  USERREFRESHSECRET   string `mapstructure:"USER_REFRESH_SECRET"`
  
}

var envs = []string{
	"PORT", "USER_SVC_URL","ADMIN_SECRET","USER_SECRET","USER_ACCESS_SECRET","USER_REFRESH_SECRET",
}

func LoadConfig() (Config, error) {
	var config Config
   
	viper.AddConfigPath(".")
	viper.SetConfigFile(".apigateway.env")
	viper.ReadInConfig()


	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
  fmt.Println("onnnnnn",config)

	return config, nil

} 
