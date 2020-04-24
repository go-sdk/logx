# LogX

![Go](https://github.com/go-sdk/logx/workflows/Go/badge.svg)
![Codecov](https://img.shields.io/codecov/c/github/go-sdk/logx)
![License](https://img.shields.io/badge/License-Apache%20License%202.0-blue)

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
	logx.Debug("1")
	logx.Info("1")
	logx.Warn("1")
	logx.Error("1")

	logx.SetLevel(logx.DebugLevel)

	logx.Debugf("2")
	logx.Infof("2")
	logx.Warnf("2")
	logx.Errorf("2")

	logx.WithField("index", 1).Info("3")
	logx.WithFields(map[string]interface{}{"index": 2}).Info("3")

	l2 := logx.Caller()
	l2.Info("4")
	logx.Info("4")

	logx.Caller(6).Info("5")
}
```

## ScreenShot

![ScreenShot](./screenshot/1.png)

## License

[Apache License 2.0](./LICENSE)
