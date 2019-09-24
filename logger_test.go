package logx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Log(GetLevel())
	assert.Equal(t, GetLevel(), InfoLevel)

	Debug("1")
	Info("1")
	Warn("1")
	Error("1")

	SetLevel(DebugLevel)
	t.Log(GetLevel())
	assert.Equal(t, GetLevel(), DebugLevel)

	Debugf("2")
	Infof("2")
	Warnf("2")
	Errorf("2")

	WithField("index", 1).Info("3")
	WithFields(map[string]interface{}{"index": 2}).Info("3")

	l2 := Caller()
	l2.Info("4")
	Info("4")
}
