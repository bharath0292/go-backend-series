package main

import (
	"os"

	"github.com/go-api/pkg/utils"
	server "github.com/go-api/servers"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading .env file: %d", err)
	}

	logLevel := zerolog.InfoLevel
	if os.Getenv("GIN_MODE") == "debug" {
		logLevel = zerolog.DebugLevel
	}

	utils.InitLogger(logLevel)
}

func main() {
	const (
		certPath = "cert/cert.pem"
		keyPath  = "cert/key.pem"
		addr     = "0.0.0.0:8080"
	)

	restEngine := server.NewRestEngine()
	restEngine.Setup()

	if err := utils.ValidateOrCreateSelfSignedCert(certPath, keyPath); err != nil {
		log.Error().Err(err).Msg("Failed to create or validate certs, falling back to HTTP")
		log.Info().Msgf("Starting HTTPS server at http://%s", addr)
		if err := restEngine.Run(addr); err != nil {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
		log.Info().Msgf("App running at http://%s", addr)
	} else {
		log.Info().Msgf("Starting HTTPS server at https://%s", addr)
		if err := restEngine.RunTLS(addr, certPath, keyPath); err != nil {
			log.Fatal().Err(err).Msg("Failed to start HTTPS server")
		}
	}
}
