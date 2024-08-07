package llm

import (
	"log"

	"github.com/AeronCyther/leet_tutor/internal/config"
)

type Agent interface {
	Chat(messages []*Message) (*Message, error)
	StructuredChat(messages []*Message, response any) (*Message, error)
}

var llmAgent Agent

func InitLLMAgent() {
	switch config.Config.LLMProvider {
	case "openai":
		llmAgent = &OpenAIAgent{}
	default:
		log.Fatalf("Invalid LLMProvider %s", config.Config.LLMProvider)
	}
}
