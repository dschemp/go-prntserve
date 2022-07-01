package main

import (
	"fmt"
	internal "github.com/dschemp/go-prntserve/internal"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"github.com/dschemp/go-prntserve/internal/handler"
	"github.com/dschemp/go-prntserve/internal/logging"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	version      = "v0.0.0"  // will be set by the build pipeline (see Makefile)
	buildNumber  = "unknown" // ^
	distribution = "custom"  // ^
)

var versionString = fmt.Sprintf("%s-%s (%s)\n", version, distribution, buildNumber)

func main() {
	// cmd.ParseArguments()
	app := &cli.App{
		Name:     "go-prntserve",
		Usage:    "A small and simple web app that allows for simple file up- and download",
		Version:  versionString,
		Flags:    appFlags(),
		Commands: appCommands(),
		Authors: []*cli.Author{
			{
				Name:  "Daniel Schemp",
				Email: "dschemp@mailbox.org",
			},
		},
		Action: runApp,
	}

	cli.VersionFlag = &cli.BoolFlag{}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error occurred when running the application")
	}
}

func appCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "version",
			Usage: "Show the version",
			Action: func(context *cli.Context) error {
				fmt.Println(versionString)
				return nil
			},
		},
	}
}

func appFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "web.address",
			Usage:       "Address on which to listen",
			Value:       ":8080",
			Destination: &cmd.SETTINGS.ListenAddress,
		},
		&cli.StringFlag{
			Name:        "jwt.secret",
			Usage:       "Instance-wide secret used for JWT authorization",
			Value:       "CHANGE_ME",
			Destination: &cmd.SETTINGS.JWTSecret,
		},
		&cli.BoolFlag{
			Name:        "web.rproxy",
			Usage:       "Is instance behind a reverse proxy which passes IP information?",
			Value:       false,
			Destination: &cmd.SETTINGS.IsBehindReverseProxy,
		},
		&cli.StringFlag{
			Name:        "storage.path",
			Usage:       "Folder path to the directory in which all files will be stored",
			Value:       "files",
			Destination: &cmd.SETTINGS.StoragePath,
		},
		&cli.BoolFlag{
			Name:        "log.verbose",
			Aliases:     []string{"v"},
			Usage:       "Enable debug / verbose logging",
			Value:       false,
			Destination: &cmd.SETTINGS.UseDebugLogging,
		},
		&cli.BoolFlag{
			Name:        "log.useJson",
			Usage:       "Use JSON logging format instead of logfmt",
			Value:       false,
			Destination: &cmd.SETTINGS.UseJSONLogging,
		},
	}
}

func runApp(context *cli.Context) error {
	// Setup logging
	logging.SetupLogging()

	log.Info().
		Str(logging.FolderPathFieldName, cmd.FullStoragePath()).
		Str(logging.StateFieldName, "starting").
		Msg("Probing storage path")

	err := handler.ProbeStoragePathOnFS()
	if err != nil {
		log.Err(err).
			Msg("Error probing storage path")
		return err
	}

	log.Info().
		Str(logging.ResultFieldName, "success").
		Str(logging.StateFieldName, "done").
		Msg("Probing storage path")

	// Start server
	internal.StartServer()

	return nil
}
