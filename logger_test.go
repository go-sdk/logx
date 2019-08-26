package logx

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	l := NewLogger()
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
}

func TestDefaultLogger(t *testing.T) {
	SetLevel(WarnLevel)
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}

func TestDiscardLogger(t *testing.T) {
	l := DiscardLogger()
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
}
