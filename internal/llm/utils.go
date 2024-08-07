package llm

import (
	"fmt"
	"strings"

	"github.com/AeronCyther/leet_tutor/internal/problem"
)

const EXPLAIN_PROBLEM_SYSTEM_PROMPT = `You are an expert problem solver capable of solving the toughest of coding problems.
Given a problem, explain in detail the problem statement to the user.
Give examples to go along with your explanation.
Do not refuse to explain or say that no further explanation is needed.
Your response should not contain anything other than the explanation.
Use markdown syntax.
Do not wrap your response with any symbol.
Do not include a title or any text preceding the explanation.
Do not repeat the problem statement or the constraints.`

const EXPLAIN_PROBLEM_EXAMPLE_SYSTEM_PROMPT = `You are an expert problem solver capable of solving the toughest of coding problems.
Given a problem and its example, explain in detail the example to the user.
Explain why the output is a result of the given input.
Do not refuse to explain or say that no further explanation is needed.
Your response should not contain anything other than the explanation.
Use markdown syntax.
Do not wrap your response with any symbol.
Do not include a title or any text preceding the explanation.
Do not repeat the problem statement or the constraints.`

func ExplainProblemContent(p *problem.Problem) (string, error) {
	response, err := llmAgent.Chat([]*Message{
		{
			Role:    "system",
			Content: EXPLAIN_PROBLEM_SYSTEM_PROMPT,
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

func ExplainProblemExample(p *problem.Problem, exampleIndex int) (string, error) {
	example := p.Body.Examples[exampleIndex]
	response, err := llmAgent.Chat([]*Message{
		{
			Role:    "system",
			Content: EXPLAIN_PROBLEM_SYSTEM_PROMPT,
		},
		{
			Role: "user",
			Content: fmt.Sprintf("```problem statement\n%s\n```\n\n```constraints\n%s\n```\n\n```example\ninputs: %s\noutputs: %s\n```",
				strings.Join(p.Body.Content, "\n"),
				strings.Join(p.Body.Constraints, "\n"),
				strings.Join(example.Inputs, "\n"),
				strings.Join(example.Outputs, "\n")),
		},
	})
	if err != nil {
		return "", err
	}

	return response.Content, nil
}
