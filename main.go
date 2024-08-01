package main

import (
	"log"

	"github.com/AeronCyther/leet_tutor/internal/app"
)

func main() {
	app := app.Init()
	log.Fatal(app.Listen(":3000"))
}
