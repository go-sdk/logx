package logx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevel(t *testing.T) {
	assert.Equal(t, DebugLevel, ParseLevel("debug"))
	assert.Equal(t, InfoLevel, ParseLevel("info"))
	assert.Equal(t, WarnLevel, ParseLevel("warn"))
	assert.Equal(t, ErrorLevel, ParseLevel("error"))
	assert.Equal(t, OffLevel, ParseLevel("fatal"))
	assert.Equal(t, OffLevel, ParseLevel(""))

	assert.Equal(t, "debug", DebugLevel.String())
	assert.Equal(t, "info", InfoLevel.String())
	assert.Equal(t, "warn", WarnLevel.String())
	assert.Equal(t, "error", ErrorLevel.String())
	assert.Equal(t, "off", OffLevel.String())
}
