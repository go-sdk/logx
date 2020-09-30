package logx

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	l1 := New()
	l1.Debug("debug1")

	l2 := NewWithLevel(InfoLevel)
	l2.Debug("debug2")

	l3 := NewWithWriters()
	l3.Debug("debug3")

	bb := &bytes.Buffer{}

	w := NewConsoleWriter(ConsoleWriterConfig{Out: bb})
	l4 := NewWithWriters(w)
	SetLogger(l4)
	assert.Equal(t, GetLevel(), DebugLevel)

	Debug("debug")
	assert.Contains(t, bb.String(), "debug")

	Info("info")
	assert.Contains(t, bb.String(), "info")

	Warn("warn")
	assert.Contains(t, bb.String(), "warn")

	Error("error")
	assert.Contains(t, bb.String(), "error")

	SetLevel(InfoLevel)
	bb.Reset()
	assert.Equal(t, GetLevel(), InfoLevel)

	Debugf("debugf")
	assert.Empty(t, bb.String())

	Infof("infof")
	assert.Contains(t, bb.String(), "infof")

	Warnf("warnf")
	assert.Contains(t, bb.String(), "warnf")

	Errorf("errorf")
	assert.Contains(t, bb.String(), "errorf")

	bb.Reset()

	WithField("aaa", 1).Info("info")
	assert.Contains(t, bb.String(), "aaa")

	WithFields(map[string]interface{}{"bbb": 2}).Info("info")
	assert.Contains(t, bb.String(), "bbb")

	bb.Reset()

	l5 := Caller()
	l5.Info("info")
	assert.Contains(t, bb.String(), ">")

	bb.Reset()

	Info("info")
	assert.NotContains(t, bb.String(), ">")

	bb.Reset()

	Caller(6).Info("info")
	assert.Contains(t, bb.String(), ">")
}
