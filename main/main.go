package main

import (
	"fmt"
	"kaspar/api"
	"kaspar/configuration"

	"github.com/rs/zerolog/log"
)

func main() {

	log.Debug().Msg("Starting Kaspar& stock recommendation service")

	configuration.GetConfiguration()
	router := api.HTTPRouteEndpoints()

	go api.GrpcServe()

	router.Run(fmt.Sprintf(":%s", configuration.GetEnvAsString("SERVER_PORT", "8080")))

}
