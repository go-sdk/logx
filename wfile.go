package logx

import (
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type FileWriterConfig struct {
	Level      Level
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
}

type FileWriter struct {
	c *FileWriterConfig
	f logrus.Formatter
	w lumberjack.Logger
}

func (l *Logger) AddFileWriter(config *FileWriterConfig) *FileWriter {
	if config.Level == 0 {
		config.Level = InfoLevel
	}
	w := &FileWriter{
		c: config,
		f: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		},
		w: lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize,
			MaxAge:     config.MaxAge,
			MaxBackups: config.MaxBackups,
			LocalTime:  config.LocalTime,
			Compress:   config.Compress,
		},
	}
	l.l.AddHook(w)
	closes = append(closes, w.Close)
	return w
}

func AddFileWriter(config *FileWriterConfig) *FileWriter {
	return log.AddFileWriter(config)
}

func (w *FileWriter) Levels() []logrus.Level {
	return logrus.AllLevels[1 : w.c.Level+1]
}

func (w *FileWriter) Fire(entry *logrus.Entry) (err error) {
	b, err := w.f.Format(entry)
	if err != nil {
		return err
	}
	_, err = w.w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (w *FileWriter) Rotate() error {
	return w.w.Rotate()
}

func (w *FileWriter) Close() {
	_ = w.w.Close()
}
