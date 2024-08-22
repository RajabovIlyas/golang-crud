package main

import (
	"github.com/RajabovIlyas/golang-crud/internal/pkg/app"
	"log"
)

func main() {

	a, err := app.New()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = a.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}
