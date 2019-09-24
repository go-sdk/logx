package logx

import (
	"time"

	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(v interface{})
	Info(v interface{})
	Warn(v interface{})
	Error(v interface{})

	Debugf(s string, v ...interface{})
	Infof(s string, v ...interface{})
	Warnf(s string, v ...interface{})
	Errorf(s string, v ...interface{})

	WithField(k string, v interface{}) Logger
	WithFields(kv map[string]interface{}) Logger

	Caller(skip ...int) Logger

	GetLevel() Level
	SetLevel(level Level)
}

// --------------------------------------------------------------------------------

const minSkipFrameCount = 5

type logger struct {
	log    zerolog.Logger
	skip   int
	fields map[string]interface{}
}

func New() Logger {
	return NewWithLevel(InfoLevel)
}

func NewWithLevel(level Level) Logger {
	o := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
	})
	return &logger{
		log:    zerolog.New(o).Level(zerolog.Level(level)),
		fields: make(map[string]interface{}, 8),
	}
}

func (l *logger) Debug(v interface{}) {
	l.print(DebugLevel, v)
}

func (l *logger) Info(v interface{}) {
	l.print(InfoLevel, v)
}

func (l *logger) Warn(v interface{}) {
	l.print(WarnLevel, v)
}

func (l *logger) Error(v interface{}) {
	l.print(ErrorLevel, v)
}

func (l *logger) Debugf(s string, v ...interface{}) {
	l.printf(DebugLevel, s, v...)
}

func (l *logger) Infof(s string, v ...interface{}) {
	l.printf(InfoLevel, s, v...)
}

func (l *logger) Warnf(s string, v ...interface{}) {
	l.printf(WarnLevel, s, v...)
}

func (l *logger) Errorf(s string, v ...interface{}) {
	l.printf(ErrorLevel, s, v...)
}

func (l *logger) WithField(k string, v interface{}) Logger {
	ll := l.new()
	ll.fields[k] = v
	return ll
}

func (l *logger) WithFields(kv map[string]interface{}) Logger {
	ll := l.new()
	for k, v := range kv {
		ll.fields[k] = v
	}
	return ll
}

func (l *logger) Caller(skip ...int) Logger {
	ll := l.new()
	if len(skip) > 0 && skip[0] > minSkipFrameCount {
		ll.skip = skip[0]
	} else {
		ll.skip = minSkipFrameCount
	}
	return ll
}

func (l *logger) GetLevel() Level {
	return Level(l.log.GetLevel())
}

func (l *logger) SetLevel(level Level) {
	l.log = l.log.Level(zerolog.Level(level))
}

func (l *logger) new() *logger {
	fields := make(map[string]interface{}, 8)
	for k, v := range l.fields {
		fields[k] = v
	}
	return &logger{
		log:    l.log,
		skip:   l.skip,
		fields: fields,
	}
}

func (l *logger) print(lvl Level, v interface{}) {
	l.printf(lvl, "%v", v)
}

func (l *logger) printf(lvl Level, s string, v ...interface{}) {
	ll := l.new().log
	if l.skip >= minSkipFrameCount {
		ll = ll.With().CallerWithSkipFrameCount(l.skip).Logger()
	}
	ll.WithLevel(zerolog.Level(lvl)).Timestamp().Fields(l.fields).Msgf(s, v...)
}

// --------------------------------------------------------------------------------

var log = New()

func Debug(v interface{}) {
	log.Debug(v)
}

func Info(v interface{}) {
	log.Info(v)
}

func Warn(v interface{}) {
	log.Warn(v)
}

func Error(v interface{}) {
	log.Error(v)
}

func Debugf(s string, v ...interface{}) {
	log.Debugf(s, v...)
}

func Infof(s string, v ...interface{}) {
	log.Infof(s, v...)
}

func Warnf(s string, v ...interface{}) {
	log.Warnf(s, v...)
}

func Errorf(s string, v ...interface{}) {
	log.Errorf(s, v...)
}

func WithField(k string, v interface{}) Logger {
	return log.WithField(k, v)
}

func WithFields(kv map[string]interface{}) Logger {
	return log.WithFields(kv)
}

func Caller(skip ...int) Logger {
	return log.Caller(skip...)
}

func GetLevel() Level {
	return log.GetLevel()
}

func SetLevel(level Level) {
	log.SetLevel(level)
}
