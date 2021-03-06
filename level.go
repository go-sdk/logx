package logx

import (
	"strings"

	"github.com/rs/zerolog"
)

type Level uint8

const (
	DebugLevel = Level(zerolog.DebugLevel)
	InfoLevel  = Level(zerolog.InfoLevel)
	WarnLevel  = Level(zerolog.WarnLevel)
	ErrorLevel = Level(zerolog.ErrorLevel)
	FatalLevel = Level(zerolog.FatalLevel)
	OffLevel   = Level(zerolog.Disabled)
)

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}
	return "off"
}

func ParseLevel(level string) Level {
	switch strings.ToLower(level) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	}
	return OffLevel
}
