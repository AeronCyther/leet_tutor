package llm

import (
	"fmt"
	"strings"

	"github.com/AeronCyther/leet_tutor/internal/problem"
)

func ExplainProblemContent(p *problem.Problem) (string, error) {
	response, err := llmAgent.Chat([]*Message{
		{
			Role:    "system",
			Content: "You are an expert problem solver capable of solving the toughest of coding problems. Given a problem, explain in detail the problem statement to the user. Give examples to go along with your explanation. Do not refuse to explain or say that no further explanation is needed. Your response should not contain anything other than the explanation. Use markdown syntax. Do not wrap your response with any symbol. Do not include a title or any text preceding the explanation. Do not repeat the problem statement or the constraints.",
		},
		{
			Role: "user",
			Content: fmt.Sprintf("```problem statement\n%s\n```\n\n```constraints\n%s\n```",
				strings.Join(p.Body.Content, "\n"), strings.Join(p.Body.Constraints, "\n")),
		},
	})
	if err != nil {
		return "", err
	}

	return response.Content, nil
}
