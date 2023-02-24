package main

import (
	"github.com/scitax/testGo/handlers"
	"github.com/scitax/testGo/routes"
)

var Global string = "sad"

func main() {
	go handlers.Preproces()
	app := routes.NewRouter()
	app.Listen(":3000")
}
