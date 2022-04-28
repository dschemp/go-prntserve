package main

import (
	"fmt"
	internal "github.com/dschemp/go-prntserve/internal"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"github.com/dschemp/go-prntserve/internal/handler"
	"github.com/dschemp/go-prntserve/internal/logging"
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
	logging.SetupLogging()

	log.Info().
		Str(logging.FolderPathFieldName, cmd.FullStoragePath()).
		Str(logging.StateFieldName, "starting").
		Msg("Probing storage path")
	err := handler.ProbeStoragePathOnFS()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error probing storage path")
	}
	log.Info().
		Str(logging.ResultFieldName, "success").
		Str(logging.StateFieldName, "done").
		Msg("Probing storage path")

	// Start server
	internal.StartServer()
}