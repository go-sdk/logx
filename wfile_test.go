package logx

import (
	"io/ioutil"
	"testing"
)

func TestAddFileWriter(t *testing.T) {
	logPath := "log.log"

	AddFileWriter(&FileWriterConfig{
		Level:    DebugLevel,
		Filename: logPath,
	})
	defer Close()

	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")

	bs, _ := ioutil.ReadFile(logPath)
	t.Logf("%s", bs)
}
