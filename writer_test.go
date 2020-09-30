package logx

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConsoleWriter(t *testing.T) {
	w := NewConsoleWriter(ConsoleWriterConfig{})
	l := NewWithWriters(w)
	l.Info("info")
}

func TestNewConsoleWriter2(t *testing.T) {
	bb := &bytes.Buffer{}

	w := NewConsoleWriter(ConsoleWriterConfig{Out: bb, Level: InfoLevel})
	l := NewWithWriters(w)

	l.Debug("debug")
	assert.Empty(t, bb.String())

	l.Info("info")
	assert.Contains(t, bb.String(), "info")

	l.SetLevel(WarnLevel)
	bb.Reset()

	l.Info("info")
	assert.Empty(t, bb.String())

	l.Error("error")
	assert.Contains(t, bb.String(), "error")
}

func TestNewFileWriter1(t *testing.T) {
	path := t.Name() + ".log"

	w := NewFileWriter(FileWriterConfig{NoColor: true, Filename: path})
	l := NewWithWriters(w)

	l.Info("info1")

	_, err := os.Stat(path)
	assert.Nil(t, err)

	err = w.Rotate()
	assert.Nil(t, err)

	l.Info("info2")
}

func TestNewFileWriter2(t *testing.T) {
	path := t.Name() + ".log"

	w := NewFileWriter(FileWriterConfig{Type: JSONFileWriter, Filename: path})
	l := NewWithWriters(w)

	l.Info("info1")

	_, err := os.Stat(path)
	assert.Nil(t, err)

	err = w.Rotate()
	assert.Nil(t, err)

	l.Info("info2")
}
