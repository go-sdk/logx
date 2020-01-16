package logx

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// --------------------------------------------------------------------------------

type ConsoleWriterConfig struct {
	Level   Level
	NoColor bool
	w       zerolog.ConsoleWriter
}

func NewConsoleWriter(conf ConsoleWriterConfig) ConsoleWriterConfig {
	conf.w = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stdout
		w.NoColor = conf.NoColor
		w.TimeFormat = time.RFC3339
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
