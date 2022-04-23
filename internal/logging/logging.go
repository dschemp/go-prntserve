package logging

import (
	"github.com/dschemp/go-prntserve/internal/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	FileNameFieldName       = "file"
	FolderPathFieldName     = "path"
	FileSizeFieldName       = "size"
	ListenAddressFieldName  = "listenAddress"
	ResultFieldName         = "result"
	StateFieldName          = "state"
	HTTPStatusCodeFieldName = "status_code"
	SystemCodeFieldName     = "code"
	RequestIDFieldName      = "req_id"
)

func SetupLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.TimestampFieldName = "ts" // "timestamp"
	if cmd.DebugLogging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	if !cmd.UseJSONLogging() {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}
