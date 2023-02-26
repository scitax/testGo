package main

import (
	"github.com/scitax/testGo/models/db_tests"
)

var Global string = "sad"

func main() {
	// go handlers.Preproces()
	// app := routes.NewRouter()
	// app.Listen(":3000")
	db_tests.SearchIndex()
}
