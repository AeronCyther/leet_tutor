package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type LeetConfig struct {
	Port int    `validate:"required"`
	Env  string `validate:"oneof=dev prod"`
}

var conf *LeetConfig

func GetOrInitConfig() *LeetConfig {
	if conf == nil {
		err := viper.BindEnv("port")
		if err != nil {
			log.Fatalf("unable to bind env \"port\",\n%v", err)
		}

		err = viper.BindEnv("env")
		if err != nil {
			log.Fatalf("unable to bind env \"env\",\n%v", err)
		}

		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatalf("unable to decode into struct,\n%v", err)
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(conf)
		if err != nil {
			log.Fatalf("unable to validate config,\n%v", err)
		}

	}
	log.Println(conf)
	return conf
}
