package main

import (
	_ "github.com/RajabovIlyas/golang-crud/docs"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/app"
	"log"
)

//	@title			GRUD API
//	@version		1.0v
//	@description	A Tag service API in golang using gin

//	@securityDefinitions.apikey ApiKeyAuth
//	@in header
//@name Authorization

// @host		localhost:3000
// @basePath	/api/v1
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
