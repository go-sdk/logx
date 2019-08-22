package logx

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	l *logrus.Logger
	*Entry
}

var (
	log = NewLogger()

	closes []func()
)

func NewLogger() *Logger {
	log := &Logger{}
	log.l = logrus.New()
	log.l.SetLevel(logrus.DebugLevel)
	log.l.SetOutput(os.Stdout)
	log.l.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		TimestampFormat:  time.RFC3339,
		QuoteEmptyFields: true,
	})
	log.Entry = &Entry{logrus.NewEntry(log.l)}
	return log
}

func DefaultLogger() *Logger {
	return log
}

func Close() {
	for i := range closes {
		closes[i]()
	}
}

func (l *Logger) SetLevel(level Level) {
	l.l.SetLevel(logrus.Level(level))
}

func (l *Logger) GetLevel() Level {
	return Level(l.l.GetLevel())
}

func (l *Logger) WithError(err error) *Entry {
	return &Entry{e: l.l.WithError(err)}
}

func (l *Logger) WithContext(ctx context.Context) *Entry {
	return &Entry{e: l.l.WithContext(ctx)}
}

func (l *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{e: l.l.WithField(key, value)}
}

func (l *Logger) WithFields(fields map[string]interface{}) *Entry {
	return &Entry{e: l.l.WithFields(fields)}
}

func (l *Logger) WithTime(t time.Time) *Entry {
	return &Entry{e: l.l.WithTime(t)}
}

type Entry struct {
	e *logrus.Entry
}

func (e *Entry) Debug(args ...interface{}) {
	e.e.Debug(args...)
}

func (e *Entry) Info(args ...interface{}) {
	e.e.Info(args...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.e.Warn(args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.e.Error(args...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.e.Fatal(args...)
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.e.Debugf(format, args...)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.e.Infof(format, args...)
}

func (e *Entry) Warnf(format string, args ...interface{}) {
	e.e.Warnf(format, args...)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	e.e.Errorf(format, args...)
}

func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.e.Fatalf(format, args...)
}

func (e *Entry) Debugln(args ...interface{}) {
	e.e.Debugln(args...)
}

func (e *Entry) Infoln(args ...interface{}) {
	e.e.Infoln(args...)
}

func (e *Entry) Warnln(args ...interface{}) {
	e.e.Warnln(args...)
}

func (e *Entry) Errorln(args ...interface{}) {
	e.e.Errorln(args...)
}

func (e *Entry) Fatalln(args ...interface{}) {
	e.e.Fatalln(args...)
}

func SetLevel(level Level) {
	log.SetLevel(level)
}

func GetLevel() Level {
	return log.GetLevel()
}

func WithError(err error) *Entry {
	return log.WithError(err)
}

func WithContext(ctx context.Context) *Entry {
	return log.WithContext(ctx)
}

func WithField(key string, value interface{}) *Entry {
	return log.WithField(key, value)
}

func WithFields(fields map[string]interface{}) *Entry {
	return log.WithFields(fields)
}

func WithTime(t time.Time) *Entry {
	return log.WithTime(t)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Debugln(args ...interface{}) {
	log.Debugln(args...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

func Warnln(args ...interface{}) {
	log.Warnln(args...)
}

func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}
