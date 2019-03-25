# Log

[![GoDoc](https://godoc.org/github.com/libgolang/log?status.svg)](https://godoc.org/github.com/libgolang/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/libgolang/log)](https://goreportcard.com/report/github.com/libgolang/log)
[![Build Status](https://travis-ci.org/libgolang/log.svg?branch=master)](https://travis-ci.org/libgolang/log)


- It ehnances Golang's logger with levels.
- It wraps Golang's `log` package


## Usage

Download it

```bash
go get -u github.com/libgolang/log

```


In Code

```go
import "github.com/libgolang/log"


func main() {

	log.SetLevel(log.LevelDebug) // defaults to log.LevelInfo

	someVar := "example"

	log.Debug("message", someVar)
	log.Debugf("message: %s", someVar)

	log.Info("message", someVar)
	log.Infof("message: %s", someVar)

	log.Warn("message", someVar)
	log.Warnf("message: %s", someVar)

	log.Error("message", someVar)
	log.Errorf("message: %s", someVar)

	log.Fatal("message", someVar)
	log.Fatalf("message: %s", someVar)
}

```

## Why
- Does no require 3rd party libraries

