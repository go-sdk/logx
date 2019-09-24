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

	assert.Equal(t, DebugLevel.String(), "debug")
	assert.Equal(t, InfoLevel.String(), "info")
	assert.Equal(t, WarnLevel.String(), "warn")
	assert.Equal(t, ErrorLevel.String(), "error")
	assert.Equal(t, OffLevel.String(), "off")
}
