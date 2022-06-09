package main

import (
	"band-app-go/pkg/http/rest"
	"band-app-go/pkg/insert"
	"band-app-go/pkg/storage/mongo"
	"band-app-go/pkg/util"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	err := runBandApp()
	if err != nil {
		log.Fatal().Msgf("%+v", err)
	}
}

func runBandApp() error {
	log.Info().Msg("⌛ Starting up Band-app")

	config, err := util.LoadConfig()
	if err != nil {
		return err
	}
	log.Info().Msg("✅ Loaded config - Got enviroment variables")

	log.Debug().Msg("user: " + config.MongoUsername)
	log.Debug().Msg("password: " + config.MongoPassword)
	log.Debug().Msg("host: " + config.MongoHost)
	log.Debug().Msg("port: " + config.MongoPort)
	log.Debug().Msg("scheme: " + config.MongoScheme)

	dbConn, err := mongo.ConnectToMongo(&config)
	if err != nil {
		return err
	}
	log.Info().Msg("✅ Connected to storage")

	insertService := insert.NewService(dbConn)

	router := rest.Handler(insertService)
	log.Info().Msg("✅ Router created")

	// TODO add functionality for graceful shutdown
	log.Info().Msg("✅ Band-app is now serving")
	log.Fatal().Err(http.ListenAndServe(":8080", router))

	return nil
}
