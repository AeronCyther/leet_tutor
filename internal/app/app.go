package app

import (
	"log"
	"strconv"
	"strings"

	"github.com/AeronCyther/leet_tutor/internal/config"
	"github.com/AeronCyther/leet_tutor/internal/problem"
	"github.com/AeronCyther/leet_tutor/internal/views"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

const MAX_PROBLEMS_PER_PAGE = 10

func Init() *fiber.App {
	app := fiber.New()

	conf := config.GetOrInitConfig()
	problem.InitProblemMap()

	if conf.Env == "dev" {
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
		totalPages, currentPage, offset, numProblemsInPage := determineProblemPaginationParams(c)
		problems := problem.ProblemSlice[offset : offset+numProblemsInPage]
		return RenderComponent(c, views.ProblemList(
			problems,
			currentPage,
			totalPages,
			string(c.Request().URI().Path()),
		))
	})

	app.Get("/fragment/problem", func(c fiber.Ctx) error {
		totalPages, currentPage, offset, numProblemsInPage := determineProblemPaginationParams(c)
		problems := problem.ProblemSlice[offset : offset+numProblemsInPage]
		return RenderComponent(c, views.ProblemListFragment(
			problems,
			currentPage,
			totalPages,
			strings.TrimPrefix(string(c.Request().URI().Path()), "/fragment"),
		))
	})

	app.Get("/problem/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		if _, ok := problem.ProblemMap[id]; ok {
			return RenderComponent(c, views.Problem(problem.ProblemMap[id]))
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
		if _, ok := problem.ProblemMap[id]; ok {
			return RenderComponent(c, views.ProblemFragment(problem.ProblemMap[id]))
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

func determineProblemPaginationParams(c fiber.Ctx) (int, int, int, int) {
	totalPages := len(problem.ProblemSlice) / MAX_PROBLEMS_PER_PAGE
	if len(problem.ProblemSlice)%MAX_PROBLEMS_PER_PAGE != 0 {
		totalPages++
	}
	currentPage, err := strconv.Atoi(c.Query("p"))
	if err != nil || currentPage < 0 {
		currentPage = 0
	} else if currentPage >= totalPages {
		currentPage = totalPages - 1
	}
	offset := currentPage * MAX_PROBLEMS_PER_PAGE
	numProblemsInPage := min(MAX_PROBLEMS_PER_PAGE,
		len(problem.ProblemSlice)-offset)
	return totalPages, currentPage, offset, numProblemsInPage
}
