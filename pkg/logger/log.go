package logger

import (
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogr() logr.Logger {
	zl := New()
	return zerologr.New(&zl)
}

func New() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"

	zl := zerolog.New(os.Stderr)
	zl = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zl = zl.With().Caller().Timestamp().Logger()
	return zl
}
