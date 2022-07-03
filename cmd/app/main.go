package main

import (
	"github.com/cbr4yan/backend-template/cmd/app/loader"
)

func main() {
	app := loader.New("backend-template")
	app.Setup()
	app.Run()
}
