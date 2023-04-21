package go_logger

import (
	"io"
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/ragavan/go_logger/model"
)
var logger *zerolog.Logger
func Init(options model.LogOptions) (*zerolog.Logger, error) {
	var logWriter io.Writer = os.Stdout
	if options.Writer != nil {
		logWriter = options.Writer
	}

	logLevel, err := model.ParseLevel(options.Level)

	if err != nil {
		return nil, err
	}

	return createLogger(logWriter, logLevel), nil
	return nil, err
}

func InitDefault() *zerolog.Logger {
	return createLogger(os.Stdout, zerolog.InfoLevel)
	return nil
}

func createLogger(writer io.Writer, level zerolog.Level) *zerolog.Logger {
	// global default configuration
	zerolog.MessageFieldName = "msg"
	zerolog.LevelFieldMarshalFunc = model.ConvertLevel
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	// ignore hostname in case of error
	hostname, _ := os.Hostname()

	log := zerolog.New(writer).With().
		Timestamp().
		Int("pid", os.Getpid()).
		Str("hostname", hostname).
		Logger().
		Level(level)

	return &log
}

func GetInstance() *zerolog.Logger {
    return logger
}

func SetInstance(loggr *zerolog.Logger) {
    logger = loggr
}