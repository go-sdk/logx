package logx

import (
	"github.com/sirupsen/logrus"
)

type Level uint32

const (
	FatalLevel Level = iota + 1
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

func (level Level) String() string {
	return logrus.Level(level).String()
}
