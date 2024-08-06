package problem

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/AeronCyther/leet_tutor/internal/config"
	"github.com/gofiber/fiber/v3/log"
)

var (
	ProblemMap   = map[string]*Problem{}
	ProblemSlice = []*Problem{}
)

type ProblemExample struct {
	Inputs      []string `json:"inputs"`
	Outputs     []string `json:"outputs"`
	Explanation []string `json:"explanation"`
}

type ProblemBody struct {
	Content     []string          `json:"content"`
	Examples    []*ProblemExample `json:"examples"`
	Constraints []string          `json:"constraints"`
	FollowUp    []string          `json:"followUp"`
}

type Problem struct {
	ID         string      `json:"id"`
	Title      string      `json:"title"`
	Difficulty string      `json:"difficulty"`
	Body       ProblemBody `json:"body"`
}

func InitProblemMap() {
	ProblemsDirectoryPath := config.Config.ProblemsDirectory
	ProblemsDirectory, err := os.ReadDir(ProblemsDirectoryPath)
	if err != nil {
		log.Fatal("Could not read the problems directory")
	}

	numProblems := 0

	for _, pf := range ProblemsDirectory {
		if strings.HasSuffix(pf.Name(), ".json") {
			numProblems++
		}
	}

	ProblemSlice = make([]*Problem, numProblems)
	problemSliceIndex := 0

	for _, pf := range ProblemsDirectory {
		if strings.HasSuffix(pf.Name(), ".json") {
			data, err := os.ReadFile(path.Join(ProblemsDirectoryPath, pf.Name()))
			if err != nil {
				log.Fatal("Could not read problem data, reading", pf.Name())
			}

			err = json.Unmarshal(data, &ProblemSlice[problemSliceIndex])
			if err != nil {
				log.Fatal("Could not unmarshal data, reading", pf.Name())
			}
			problemSliceIndex += 1
		}
	}

	for p := range ProblemSlice {
		ProblemMap[ProblemSlice[p].ID] = ProblemSlice[p]
	}
}
