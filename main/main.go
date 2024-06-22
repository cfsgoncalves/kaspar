package main

import (
	"kaspar/api"
	"kaspar/configuration"

	"github.com/rs/zerolog/log"
)

func main() {

	log.Debug().Msg("Starting Kaspar& stock recommendation service")

	configuration.GetConfiguration()
	api.NewRouter()
}
