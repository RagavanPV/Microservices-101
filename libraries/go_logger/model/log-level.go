package model

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

type LogLevel string

// Represents all the Log log levels accepted
// Note: the Panic level has been added to cover this log level not available in Javascript
const (
	Trace LogLevel = "10"
	Debug LogLevel = "20"
	Info  LogLevel = "30"
	Warn  LogLevel = "40"
	Error LogLevel = "50"
	Fatal LogLevel = "60"
	Panic LogLevel = "70"
)

// ConvertLevel Convert a zerolog log level into the corresponding Log level
func ConvertLevel(level zerolog.Level) string {
	var LogLevel LogLevel
	switch level {
	case zerolog.TraceLevel:
		LogLevel = Trace
	case zerolog.DebugLevel:
		LogLevel = Debug
	case zerolog.InfoLevel:
		LogLevel = Info
	case zerolog.WarnLevel:
		LogLevel = Warn
	case zerolog.ErrorLevel:
		LogLevel = Error
	case zerolog.FatalLevel:
		LogLevel = Fatal
	case zerolog.PanicLevel:
		LogLevel = Panic
	case zerolog.Disabled:
		fallthrough
	case zerolog.NoLevel:
		fallthrough
	default:
		return ""
	}

	return string(LogLevel)
}

// ParseLevel Parse a string name of the log level and return the corresponding zerolog level
func ParseLevel(level string) (zerolog.Level, error) {
	if len(level) > 0 {
		switch strings.ToLower(level) {
		case "trace":
			return zerolog.TraceLevel, nil
		case "debug":
			return zerolog.DebugLevel, nil
		case "info":
			return zerolog.InfoLevel, nil
		case "warn":
			return zerolog.WarnLevel, nil
		case "error":
			return zerolog.ErrorLevel, nil
		case "fatal":
			return zerolog.FatalLevel, nil
		case "panic":
			return zerolog.PanicLevel, nil
		case "silent":
			return zerolog.Disabled, nil
		default:
			return zerolog.NoLevel, fmt.Errorf("level %s is not recognized", level)
		}
	}

	return zerolog.InfoLevel, nil
}