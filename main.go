package main

import (
	"fmt"
	"log"

	"github.com/AeronCyther/leet_tutor/internal/app"
	"github.com/AeronCyther/leet_tutor/internal/config"
)

func main() {
	app := app.Init()
	conf := config.GetOrInitConfig()
	log.Fatal(app.Listen(fmt.Sprintf(":%d", conf.Port)))
}
