package configuration

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetConfiguration() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal().Msgf("Could not load env file. Error: %s}", err.Error())
	}

	LOG_LEVEL, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	if err != nil {
		log.Error().Msgf("configuration.GetConfiguration(): Error converting string to int. Error: %s", err)
	}

	zerolog.SetGlobalLevel(zerolog.Level(LOG_LEVEL))
}

func GetEnvAsString(env string, defaultVar string) string {
	envVar := os.Getenv(env)

	if envVar == "" {
		log.Warn().Msgf("Variable %s does not exists. Returning the default variable", env)
		return defaultVar
	}

	return envVar
}

func GetEnvAsInt(env string, defaultVar int) int {
	envVar := os.Getenv(env)

	if envVar == "" {
		log.Warn().Msgf("configuration.GetEnvAsInt():Variable %s does not exists. Returning the default variable", env)
		return defaultVar
	}

	intVar, err := strconv.Atoi(envVar)

	if err != nil {
		log.Error().Msgf("configuration.GetEnvAsInt(): Error converting string to int: %s. Using the default variable", env)
		return defaultVar
	}
	return intVar
}
