package log

import (
	golog "log"
)

type Level int

const (
	LevelFatal Level = iota
	LevelError Level = iota
	LevelWarn  Level = iota
	LevelInfo  Level = iota
	LevelDebug Level = iota
	LevelTrace Level = iota
)

var (
	logLevel = LevelInfo
)

func SetLevel(level Level) {
	logLevel = level
}

func Fatal(v ...interface{}) {
	golog.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	golog.Fatalf(format, v...)
}

func Error(v ...interface{}) {
	if logLevel < LevelError {
		return
	}
	golog.Print(v...)
}

func Errorf(format string, v ...interface{}) {
	if logLevel < LevelError {
		return
	}
	golog.Printf(format, v...)
}

func Warn(v ...interface{}) {
	if logLevel < LevelWarn {
		return
	}
	golog.Print(v...)
}

func Warnf(format string, v ...interface{}) {
	if logLevel < LevelWarn {
		return
	}
	golog.Printf(format, v...)
}

func Info(v ...interface{}) {
	if logLevel < LevelInfo {
		return
	}
	golog.Print(v...)
}

func Infof(format string, v ...interface{}) {
	if logLevel < LevelInfo {
		return
	}
	golog.Printf(format, v...)
}

func Debug(v ...interface{}) {
	if logLevel < LevelDebug {
		return
	}
	golog.Print(v...)
}

func Debugf(format string, v ...interface{}) {
	if logLevel < LevelDebug {
		return
	}
	golog.Printf(format, v...)
}

func Trace(v ...interface{}) {
	if logLevel < LevelTrace {
		return
	}
	golog.Print(v...)
}

func Tracef(format string, v ...interface{}) {
	if logLevel < LevelTrace {
		return
	}
	golog.Printf(format, v...)
}
