package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var config Config

type Config struct {
  DBHost           string `mapstructure:"DB_HOST"`
  DBName           string `mapstructure:"DB_NAME"`
  DBUser           string `mapstructure:"DB_USER"`
  DBPassword       string `mapstructure:"DB_PASSWORD"`
  DBPort           string `mapstructure:"DB_PORT"`
  PORT             string `mapstructure:"PORT"`

  AWS_REGION            string `mapstructure:"AWS_REGION"`
  AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
  AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`



  SMTP_USERNAME string `mapstructure:"SMTP_USERNAME"`
  SMTP_PASSWORD string `mapstructure:"SMTP_PASSWORD"`
  SMTP_HOST     string `mapstructure:"SMTP_HOST"`
  SMTP_PORT     string `mapstructure:"SMTP_PORT"`

}

var envs = []string{
  "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SMTP_USERNAME", "SMTP_PASSWORD", "SMTP_HOST", "SMTP_PORT",
}

func LoadConfig() (Config, error) {

  viper.AddConfigPath(".")
  viper.SetConfigFile(".inv.env")

  if err := viper.ReadInConfig(); err != nil{
    return config, fmt.Errorf("error reading config file: %v", err)
  }

  fmt.Println("env in LOad Config")

  for _, env := range envs {
    if err := viper.BindEnv(env); err != nil {
      return config, err
    }
  }

  if err := viper.Unmarshal(&config); err != nil {

    return config, err
  }
  if err := validator.New().Struct(&config); err != nil {
    return config, err
  }

  return config, nil

}
