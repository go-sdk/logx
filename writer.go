package logx

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// --------------------------------------------------------------------------------

type ConsoleWriterConfig struct {
	Level      Level
	Out        io.Writer
	NoColor    bool
	TimeFormat string
	w          zerolog.ConsoleWriter
}

func NewConsoleWriter(conf ConsoleWriterConfig) ConsoleWriterConfig {
	if conf.Out == nil {
		conf.Out = os.Stdout
	}
	if conf.TimeFormat == "" {
		conf.TimeFormat = time.RFC3339
	}
	conf.w = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = conf.Out
		w.NoColor = conf.NoColor
		w.TimeFormat = conf.TimeFormat
	})
	return conf
}

func (w ConsoleWriterConfig) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w ConsoleWriterConfig) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level < zerolog.Level(w.Level) {
		return len(p), nil
	}
	return w.Write(p)
}

func (w ConsoleWriterConfig) Close() error {
	return nil
}

// --------------------------------------------------------------------------------

type FileWriterConfig struct {
	Level      Level
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
	w          *lumberjack.Logger
}

func NewFileWriter(conf FileWriterConfig) FileWriterConfig {
	conf.w = &lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,
		MaxAge:     conf.MaxAge,
		MaxBackups: conf.MaxBackups,
		LocalTime:  conf.LocalTime,
		Compress:   conf.Compress,
	}
	return conf
}

func (w FileWriterConfig) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w FileWriterConfig) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level < zerolog.Level(w.Level) {
		return len(p), nil
	}
	return w.Write(p)
}

func (w FileWriterConfig) Rotate() error {
	return w.w.Rotate()
}

func (w FileWriterConfig) Close() error {
	return w.w.Close()
}
