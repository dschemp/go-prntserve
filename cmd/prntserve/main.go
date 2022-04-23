package main

import (
	"fmt"
	internal "github.com/dschemp/go-prntserve/internal"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"github.com/dschemp/go-prntserve/internal/handler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	cmd.ParseArguments()
	fmt.Println("Welcome to prntserve!")

	if cmd.ShowVersion() {
		// TODO: Show version
		fmt.Printf("Version: %s\n", "unknown")
		os.Exit(0)
	}

	// Setup logging
	setupLogging()

	log.Info().
		Str("path", cmd.FullStoragePath()).
		Str("state", "starting").
		Msg("Probing storage path")
	err := handler.ProbeStoragePathOnFS()
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Info().
		Str("result", "success").
		Str("state", "done").
		Msg("Probing storage path")

	// Start server
	internal.StartServer()
}

func setupLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.TimestampFieldName = "ts" // "timestamp"
	if cmd.DebugLogging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	if !cmd.UseJSONLogging() {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out: os.Stdout,
		})
	}
}
