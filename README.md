# logx

[![CircleCI](https://img.shields.io/circleci/build/github/go-sdk/logx)](https://circleci.com/gh/go-sdk/logx)
[![Travis](https://img.shields.io/travis/go-sdk/logx/master)](https://travis-ci.org/go-sdk/logx)
[![License](https://img.shields.io/badge/license-Apache%20License%202.0-blue)](./LICENSE)

## Install

```bash
go get -u github.com/go-sdk/logx
```

## Usage

```go
package main

import (
	"github.com/go-sdk/logx"
)

func main() {
	logx.SetLevel(logx.DebugLevel)

	logx.Debug("debug")
	logx.Info("info")
	logx.Warn("warn")
	logx.Error("error")
	logx.Fatal("fatal")
	// DEBU[2019-08-27T15:00:00+08:00] debug                                        
	// INFO[2019-08-27T15:00:00+08:00] info                                         
	// WARN[2019-08-27T15:00:00+08:00] warn                                         
	// ERRO[2019-08-27T15:00:00+08:00] error  
	// FATA[2019-08-27T15:00:00+08:00] fatal  

	logx.WithField("k", "v").Info("field")
	// INFO[2019-08-27T15:04:24+08:00] field  k=v

	logx.AddFileWriter(&logx.FileWriterConfig{
		Level:    logx.DebugLevel,
		Filename: "log.log",
	})

	logx.Debug("debug")
	// {"level":"debug","msg":"debug","time":"2019-08-27T15:00:00+08:00"}
}
```

## License

[Apache License 2.0](./LICENSE)
