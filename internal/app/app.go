package app

import (
	"log"
	"strconv"
	"strings"

	"github.com/AeronCyther/leet_tutor/internal/components"
	"github.com/AeronCyther/leet_tutor/internal/config"
	"github.com/AeronCyther/leet_tutor/internal/llm"
	"github.com/AeronCyther/leet_tutor/internal/problem"
	"github.com/AeronCyther/leet_tutor/internal/search"
	"github.com/AeronCyther/leet_tutor/internal/views"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const MAX_PROBLEMS_PER_PAGE = 10

func Init() *fiber.App {
	app := fiber.New()

	config.InitConfig()
	problem.InitProblemMap()
	llm.InitLLMAgent()

	if config.Config.Env == "dev" {
		sess_id, err := uuid.NewUUID()
		if err != nil {
			log.Fatal("Error generating session ID")
		}
		app.Head("*", func(c fiber.Ctx) error {
			c.Response().Header.Add("ETag", sess_id.String())
			c.Response().Header.Add("Content-Type", "text/html")
			return nil
		})
	}

	app.Get("/", func(c fiber.Ctx) error {
		return RenderComponent(c, views.Home())
	})

	app.Get("/fragment/home", func(c fiber.Ctx) error {
		return RenderComponent(c, views.HomeFragment())
	})

	app.Get("/about", func(c fiber.Ctx) error {
		return RenderComponent(c, views.About())
	})

	app.Get("/fragment/about", func(c fiber.Ctx) error {
		return RenderComponent(c, views.AboutFragment())
	})

	app.Get("/problem", func(c fiber.Ctx) error {
		totalPages, currentPage, searchParams, problems := determineProblemPaginationParams(c)
		return RenderComponent(c, views.ProblemList(
			problems,
			currentPage,
			totalPages,
			searchParams,
			string(c.Request().URI().Path()),
		))
	})

	app.Get("/fragment/problem", func(c fiber.Ctx) error {
		totalPages, currentPage, searchParams, problems := determineProblemPaginationParams(c)
		return RenderComponent(c, views.ProblemListFragment(
			problems,
			currentPage,
			totalPages,
			searchParams,
			strings.TrimPrefix(string(c.Request().URI().Path()), "/fragment"),
		))
	})

	app.Get("/problem/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		if p, ok := problem.ProblemMap[id]; ok {
			return RenderComponent(c, views.Problem(p))
		} else {
			err := RenderComponent(c, views.NotFound())
			if err != nil {
				return err
			}
			return c.SendStatus(404)
		}
	})

	app.Get("fragment/problem/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		if p, ok := problem.ProblemMap[id]; ok {
			return RenderComponent(c, views.ProblemFragment(p))
		} else {
			err := RenderComponent(c, views.NotFound())
			if err != nil {
				return err
			}
			return c.SendStatus(404)
		}
	})

	app.Get("fragment/problem/:id/explain/content", func(c fiber.Ctx) error {
		id := c.Params("id")
		if p, ok := problem.ProblemMap[id]; ok {
			explanation, err := llm.ExplainProblemContent(p)
			if err != nil {
				log.Printf("Error: %e", err)
				return RenderComponent(c, components.AIErrorContent(string(c.Request().URI().Path())))
			}

			return RenderComponent(c, components.AIGeneratedContent(explanation, string(c.Request().URI().Path())))

		} else {
			err := RenderComponent(c, views.NotFound())
			if err != nil {
				return err
			}
			return c.SendStatus(404)
		}
	})
	return app
}

func determineProblemPaginationParams(c fiber.Ctx) (int, int, *search.Params, []*problem.Problem) {
	query := strings.ToLower(strings.Trim(c.Query("q"), " "))
	difficulty := strings.Trim(c.Query("difficulty"), " ")
	difficultiesSet := make(map[string]struct{})

	if len(difficulty) > 0 {
		for _, d := range strings.Split(difficulty, ",") {
			difficultiesSet[d] = struct{}{}
		}
	}

	filteredProblems := make([]*problem.Problem, 0)
	if len(query) > 0 || len(difficultiesSet) > 0 {
		for _, problem := range problem.ProblemSlice {
			if len(difficultiesSet) != 0 {
				_, ok := difficultiesSet[problem.Difficulty]
				if !ok {
					continue
				}
			}
			if len(query) == 0 || strings.Contains(strings.ToLower(problem.Title), query) {
				filteredProblems = append(filteredProblems, problem)
			}
		}
	} else {
		filteredProblems = problem.ProblemSlice
	}

	totalPages := len(filteredProblems) / MAX_PROBLEMS_PER_PAGE
	if len(filteredProblems)%MAX_PROBLEMS_PER_PAGE != 0 {
		totalPages++
	}
	currentPage, err := strconv.Atoi(c.Query("p"))
	if err != nil || currentPage < 0 {
		currentPage = 0
	} else if currentPage >= totalPages && totalPages > 0 {
		currentPage = totalPages - 1
	}
	offset := currentPage * MAX_PROBLEMS_PER_PAGE
	numProblemsInPage := min(MAX_PROBLEMS_PER_PAGE,
		len(filteredProblems)-offset)
	problems := filteredProblems[offset : offset+numProblemsInPage]
	return totalPages, currentPage, &search.Params{
		Query:      query,
		Difficulty: difficulty,
	}, problems
}
