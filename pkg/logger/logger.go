package logger

import (
	"github.com/cbr4yan/backend-template/config"
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger

func init() {
	Log = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

func Setup(config *config.Config) {
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
