package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type LeetConfig struct {
	Port              int    `validate:"required"`
	Env               string `validate:"oneof=dev prod"`
	ProblemsDirectory string `validate:"required,dirpath"`
}

var Config *LeetConfig

func InitConfig() {
	Config = &LeetConfig{}
	err := viper.BindEnv("port")
	if err != nil {
		log.Fatalf("unable to bind env \"port\",\n%v", err)
	}

	err = viper.BindEnv("env")
	if err != nil {
		log.Fatalf("unable to bind env \"env\",\n%v", err)
	}

	err = viper.BindEnv("ProblemsDirectory", "PROBLEMS_DIRECTORY")
	if err != nil {
		log.Fatalf("unable to bind env \"PROBLEMS_DIRECTORY\",\n%v", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct,\n%v", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(Config)
	if err != nil {
		log.Fatalf("unable to validate config,\n%v", err)
	}
}
