package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AeronCyther/leet_tutor/internal/config"
)

type ResponseFormat struct {
	Type string `json:"type"`
}

type ChatCompletionRequestBody struct {
	Model          string          `json:"model"`
	Messages       []*Message      `json:"messages"`
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
}

type ResponseChoice struct {
	Message *Message `json:"message"`
}

type ChatCompletionResponse struct {
	Choices []ResponseChoice
}

type OpenAIAgent struct{}

func (a *OpenAIAgent) invoke(messages []*Message, responseFormat *ResponseFormat) (*Message, error) {
	requestBodyData := ChatCompletionRequestBody{
		Model:          config.Config.OpenAIModel,
		Messages:       messages,
		ResponseFormat: responseFormat,
	}
	requestBodyBytes, err := json.Marshal(requestBodyData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, config.Config.OpenAIChatCompletionEndpoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.OpenAIAPIKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseDataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	responseData := &ChatCompletionResponse{}
	err = json.Unmarshal(responseDataBytes, responseData)
	if err != nil {
		return nil, err
	}

	return responseData.Choices[0].Message, nil
}

func (a *OpenAIAgent) Chat(messages []*Message) (*Message, error) {
	return a.invoke(messages, nil)
}

func (a *OpenAIAgent) StructuredChat(messages []*Message, response any) (*Message, error) {
	return a.invoke(messages, &ResponseFormat{
		Type: "json_object",
	})
}
