package main

import (
	"band-app-go/pkg/http/rest"
	"band-app-go/pkg/input"
	"band-app-go/pkg/insert"
	"band-app-go/pkg/storage/mongo"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := runBandApp()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func runBandApp() error {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("⌛ Starting up Band-app")

	envVars, err := input.GetEnvVariables()
	if err != nil {
		return err
	}
	log.Info("✅ Got environment variables")

	dbConn, err := mongo.ConnectToMongo(envVars)
	if err != nil {
		return err
	}
	log.Info("✅ Connected to storage")

	insertService := insert.NewService(dbConn)

	router := rest.Handler(insertService)
	log.Info("✅ Router created")

	log.Info("✅ Band-app is now serving")
	log.Fatal(http.ListenAndServe(":8080", router))

	return nil
}
