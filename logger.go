package logx

import (
	"io"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(v interface{})
	Info(v interface{})
	Warn(v interface{})
	Error(v interface{})
	Fatal(v interface{})

	Debugf(s string, v ...interface{})
	Infof(s string, v ...interface{})
	Warnf(s string, v ...interface{})
	Errorf(s string, v ...interface{})
	Fatalf(s string, v ...interface{})

	WithField(k string, v interface{}) Logger
	WithFields(kv map[string]interface{}) Logger

	Caller(skip ...int) Logger

	GetLevel() Level
	SetLevel(level Level)
}

// --------------------------------------------------------------------------------

const (
	defaultLevel = DebugLevel

	minSkipFrameCount = 5
)

type logger struct {
	log    zerolog.Logger
	level  Level
	skip   int
	fields map[string]interface{}
}

func New() Logger {
	return NewWithLevel(defaultLevel)
}

func NewWithLevel(level Level) Logger {
	return NewWithWriters(NewConsoleWriter(ConsoleWriterConfig{Level: level}))
}

func NewWithWriters(writers ...io.Writer) Logger {
	if len(writers) == 0 {
		writers = []io.Writer{NewConsoleWriter(ConsoleWriterConfig{Level: defaultLevel})}
	}
	return &logger{
		log:    zerolog.New(zerolog.MultiLevelWriter(writers...)).Level(zerolog.Level(defaultLevel)),
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

func (l *logger) Fatal(v interface{}) {
	l.print(FatalLevel, v)
	os.Exit(1)
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

func (l *logger) Fatalf(s string, v ...interface{}) {
	l.printf(FatalLevel, s, v...)
	os.Exit(1)
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
	return l.level
}

func (l *logger) SetLevel(level Level) {
	l.level = level
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
	if lvl < l.level {
		return
	}
	ll := l.new().log
	if l.skip >= minSkipFrameCount {
		ll = ll.With().CallerWithSkipFrameCount(l.skip).Logger()
	}
	ll.WithLevel(zerolog.Level(lvl)).Timestamp().Fields(l.fields).Msgf(s, v...)
}

// --------------------------------------------------------------------------------

var (
	log Logger

	logMu = sync.Mutex{}
)

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

func Fatal(v interface{}) {
	log.Fatal(v)
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

func Fatalf(s string, v ...interface{}) {
	log.Fatalf(s, v...)
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
	logMu.Lock()
	defer logMu.Unlock()
	return log.GetLevel()
}

func SetLevel(level Level) {
	logMu.Lock()
	defer logMu.Unlock()
	log.SetLevel(level)
}

func SetLogger(logger Logger) {
	logMu.Lock()
	defer logMu.Unlock()
	log = logger
}
