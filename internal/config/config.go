package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type LeetConfig struct {
	Port                         int    `validate:"required"`
	Env                          string `validate:"oneof=dev prod"`
	ProblemsDirectory            string `validate:"required,dirpath"`
	LLMProvider                  string `validate:"oneof=openai"`
	OpenAIAPIKey                 string `validate:"required_if=LLMProvider openai"`
	OpenAIModel                  string `validate:"required_if=LLMProvider openai"`
	OpenAIChatCompletionEndpoint string `validate:"required_if=LLMProvider openai"`
}

var Config *LeetConfig

func InitConfig() {
	Config = &LeetConfig{
		OpenAIChatCompletionEndpoint: "https://api.openai.com/v1/chat/completions",
		OpenAIModel:                  "gpt-3.5-turbo",
	}
	err := viper.BindEnv("port")
	if err != nil {
		log.Fatalf("unable to bind env \"PORT\",\n%v", err)
	}

	err = viper.BindEnv("env")
	if err != nil {
		log.Fatalf("unable to bind env \"ENV\",\n%v", err)
	}

	err = viper.BindEnv("ProblemsDirectory", "PROBLEMS_DIRECTORY")
	if err != nil {
		log.Fatalf("unable to bind env \"PROBLEMS_DIRECTORY\",\n%v", err)
	}

	err = viper.BindEnv("LLMProvider", "LLM_PROVIDER")
	if err != nil {
		log.Fatalf("unable to bind env \"LLM_PROVIDER\",\n%v", err)
	}

	err = viper.BindEnv("OpenAIAPIKey", "OPENAI_API_KEY")
	if err != nil {
		log.Fatalf("unable to bind env \"OPENAI_API_KEY\",\n%v", err)
	}

	err = viper.BindEnv("OpenAIModel", "OPENAI_MODEL")
	if err != nil {
		log.Fatalf("unable to bind env \"OPENAI_MODEL\",\n%v", err)
	}

	err = viper.BindEnv("OpenAIChatCompletionEndpoint", "OPENAI_CHAT_COMPLETION_ENDPOINT")
	if err != nil {
		log.Fatalf("unable to bind env \"OPENAI_CHAT_COMPLETION_ENDPOINT\",\n%v", err)
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
