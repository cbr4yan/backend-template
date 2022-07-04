package main

import (
	"github.com/cbr4yan/backend-template/cmd/app/loader"
)

func main() {
	app := loader.New("example")
	app.Setup()
	app.Run()
}
