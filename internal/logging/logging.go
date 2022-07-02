package logging

import (
	"fmt"
	"github.com/dschemp/go-prntserve/internal/cmd"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
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

	MethodRequestFieldName        = "method"
	SchemeRequestFieldName        = "scheme"
	HostRequestFieldName          = "host"
	PathRequestFieldName          = "path"
	ProtocolRequestFieldName      = "protocol"
	RemoteAddressRequestFieldName = "remote_address"
	StatusCodeRequestFieldName    = "status_code"
	BytesWrittenRequestFieldName  = "bytes_written"
	DurationMSRequestFieldName    = "duration_ms"
)

// SetupLogging sets up some global logging parameters.
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

// ChiLogger returns a logger for Chi that uses the app's logger.
func ChiLogger(next http.Handler) http.Handler {
	return requestLogger()(next)
}

func requestLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}

				fullRequestUri := fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

				log.Info().
					Str(MethodRequestFieldName, r.Method).
					Str(SchemeRequestFieldName, scheme).
					Str(HostRequestFieldName, r.Host).
					Str(PathRequestFieldName, r.RequestURI).
					Str(ProtocolRequestFieldName, r.Proto).
					Str(RemoteAddressRequestFieldName, r.RemoteAddr).
					Int(StatusCodeRequestFieldName, ww.Status()).
					Int(BytesWrittenRequestFieldName, ww.BytesWritten()).
					Dur(DurationMSRequestFieldName, time.Since(t1)).
					Msg(fullRequestUri)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
