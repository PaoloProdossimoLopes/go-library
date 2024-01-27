package main

import (
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/database"
	"github.com/PaoloProdossimoLopes/go-library/enviroment"
	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/route"
)

func main() {
	prepareEnviromentVariables()

	database.Init()
	route.Init()

	http.ListenAndServe(enviroment.Enviroment.GetPort(), nil)
}

func prepareEnviromentVariables() {
	if initEnvError := enviroment.Init(); initEnvError != nil {
		logger.Fatal(initEnvError.Error())
		panic(initEnvError)
	}
}
