package test

import (
	"testing"

	"github.com/go-sdk/logx"
)

func TestNewLogger(t *testing.T) {
	l := logx.NewLogger()
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
}

func TestDefaultLogger(t *testing.T) {
	logx.SetLevel(logx.WarnLevel)
	logx.Debug("debug")
	logx.Info("info")
	logx.Warn("warn")
	logx.Error("error")
}

func TestDiscardLogger(t *testing.T) {
	l := logx.DiscardLogger()
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
	l.Error("error")
}
