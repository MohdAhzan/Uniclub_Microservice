package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	PORT          string `mapstructure:"PORT"`
	UserSvcUrl    string `mapstructure:"USER_SVC_URL"`
	InventorySvcUrl    string `mapstructure:"INV_SVC_URL"`
  ADMINSECRET   string `mapstructure:"ADMIN_SECRET"`
  USERACCESSSECRET   string `mapstructure:"USER_ACCESS_SECRET"`
  USERREFRESHSECRET   string `mapstructure:"USER_REFRESH_SECRET"`
 
  AWS_REGION            string `mapstructure:"AWS_REGION"`
  AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
  AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`

}

var envs = []string{
	"PORT", "USER_SVC_URL","INV_SVC_URL","ADMIN_SECRET","USER_SECRET","USER_ACCESS_SECRET","USER_REFRESH_SECRET","AWS_REGION","AWS_ACCESS_KEY_ID","AWS_SECRET_ACCESS_KEY",
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
