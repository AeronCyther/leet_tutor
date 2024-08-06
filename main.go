package main

import (
	"fmt"
	"log"

	"github.com/AeronCyther/leet_tutor/internal/app"
	"github.com/AeronCyther/leet_tutor/internal/config"
)

func main() {
	app := app.Init()
	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Config.Port)))
}
