package logx

import (
	"io/ioutil"
	"testing"
)

func TestAddFileWriter(t *testing.T) {
	logPath := "log.log"

	SetLevel(DebugLevel)

	l := AddFileWriter(&FileWriterConfig{
		Level:    InfoLevel,
		Filename: logPath,
	})
	defer Close()

	Debug("Debug")
	Info("Info")
	Warn("Warn")
	Error("Error")

	bs, _ := ioutil.ReadFile(logPath)
	t.Logf("%s", bs)

	t.Logf("%v", l.Rotate())
}
