package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AMQP_Driver string `mapstructure:"AMQP_DRIVER"`
	AMQP_User   string `mapstructure:"AMQP_USER"`
	AMQP_Pass   string `mapstructure:"AMQP_PASSWORD"`
	AMQP_Host   string `mapstructure:"AMQP_HOST"`
	AMQP_Port   string `mapstructure:"AMQP_PORT"`

	EMAIL_Host     string `mapstructure:"CONFIG_SMTP_HOST"`
	EMAIL_Port     string `mapstructure:"CONFIG_SMTP_PORT"`
	EMAIL_Email    string `mapstructure:"CONFIG_SMTP_EMAIL"`
	EMAIL_Password string `mapstructure:"CONFIG_SMTP_PASSWORD"`
}

func LoadConfig(Dir string) (cfg *Config) {
	if Dir == "" {
		Dir = ".env"
	}
	viper.SetConfigFile(Dir)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Failed Load File ENV : " + err.Error())
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicf("Cannot Set Data Config : %s", err)
	}
	fmt.Println(os.Getenv("AMQP_DRIVER"))
	return
}
