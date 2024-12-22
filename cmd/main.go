package main

import (
	"github.com/MiklilkiM/calculator-on-go/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}