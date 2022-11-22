package main

import (
	"github.com/NYTimes/gziphandler"
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/rugby-data-api/internal/routes"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	configPath = "config.yaml"
)

const Port = "6080"

func main() {
	defer panicQuit()
	log.Fatal(godotenv.Load())
	appConfig := config.NewFromFile(configPath)
	facade, err := initializeDAO(*appConfig)
	if err != nil {
		log.Error(err)
		panicQuit()
	}
	handler := routes.Handler{Service: facade}

	router := handler.InitializeRoutes()
	c := corsHandler()

	log.Fatal(listenAndServe(Port, gziphandler.GzipHandler(c.Handler(router))))
}

func panicQuit() {
	if r := recover(); r != nil {
		log.Errorf("I panicked and am quitting: %v", r)
		log.Error("I should be alerting someone...")
	}
}
