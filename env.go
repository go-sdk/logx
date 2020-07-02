package logx

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func init() {
	B := func(key string) bool {
		return strings.ToLower(os.Getenv(key)) == "true"
	}

	I := func(key string) int {
		i, _ := strconv.ParseInt(os.Getenv(key), 10, 64)
		return int(i)
	}

	S := func(key string, defaults ...string) string {
		v, b := os.LookupEnv(key)
		if !b && len(defaults) > 0 {
			return defaults[0]
		}
		return v
	}

	FileType := func() FileWriterType {
		v := strings.ToLower(os.Getenv("LOGX_FILE_TYPE"))
		if v == "json" {
			return JSONFileWriter
		}
		return TextFileWriter
	}

	writers := []io.Writer{
		NewConsoleWriter(ConsoleWriterConfig{Level: defaultLevel}),
	}

	if level := ParseLevel(S("LOGX_FILE_LEVEL")); level != OffLevel {
		writers = append(writers, NewFileWriter(FileWriterConfig{
			Level:      ParseLevel(S("LOGX_FILE_LEVEL")),
			Type:       FileType(),
			NoColor:    B("LOGX_FILE_NO_COLOR"),
			Filename:   S("LOGX_FILE_NAME", "app.log"),
			MaxSize:    I("LOGX_FILE_MAX_SIZE"),
			MaxAge:     I("LOGX_FILE_MAX_AGE"),
			MaxBackups: I("LOGX_FILE_MAX_BACKUPS"),
			LocalTime:  B("LOGX_FILE_LOCAL_TIME"),
			Compress:   B("LOGX_FILE_COMPRESS"),
		}))
	}

	log = NewWithWriters(writers...)
}
