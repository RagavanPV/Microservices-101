package go_logger

import (
	"io"
	"os"
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/RagavanPV/Microservices-101/libraries/go_logger/model"
)
var (
	logger *zerolog.Logger
 )
 
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

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *zerolog.Logger) context.Context {
	return context.WithValue(ctx, logger, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *zerolog.Logger {
	if l, ok := ctx.Value(logger).(*zerolog.Logger); ok {
		return l
	}
	return InitDefault()
}