package log

import (
	golog "log"
)

// Level Log level to use in `SetLevel()` function
type Level int

const (
	// LevelFatal to be used with `SetLevel(LevelFatal)`, to set log level to only Fatal calls
	LevelFatal Level = iota

	// LevelError to be used with `SetLevel(LevelError)`, to set log level to Error and Fatal calls
	LevelError Level = iota

	// LevelWarn to be used with `SetLevel(LevelWarn)`, to set log level to Error, Fatal and Warn calls
	LevelWarn Level = iota

	// LevelInfo to be used with `SetLevel(LevelInfo)`, to set log level to Error, Fatal, Warn and Info cals
	LevelInfo Level = iota

	// LevelDebug to be used with `SetLevel(LevelDebug)`, to set log level to Error, Fatal, Warn, Info and Debug calls.
	LevelDebug Level = iota

	// LevelTrace to be used with `SetLevel(LevelTrace)`, it enables all log levels.
	LevelTrace Level = iota
)

var (
	logLevel = LevelInfo
)

// SetLevel is used to set the logging level.  This library only support setting
// logging level globally.  The possible arguments to this function are:
// LevelFatal, LevelError, LevelWarn, LevelInfo, LevelDebug and LevelTrace.
func SetLevel(level Level) {
	logLevel = level
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	golog.Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	golog.Fatalf(format, v...)
}

// Error calls Golang's log.Print which in turn calls Output to print to the standard logger.
// It will call log.Printf if the log level is set Error or above.
// Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	if logLevel < LevelError {
		return
	}
	golog.Print(v...)
}

// Errorf calls Golang's log.Printf which in turn calls Output to print to the standard logger.
// It will call log.Printf if the log level is set Error or above.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	if logLevel < LevelError {
		return
	}
	golog.Printf(format, v...)
}

// Warn calls Golang's log.Print which in turn calls Output to print to the standard logger.
// It will call log.Printf if the log level is set Warn or above.
// Arguments are handled in the manner of fmt.Print.
func Warn(v ...interface{}) {
	if logLevel < LevelWarn {
		return
	}
	golog.Print(v...)
}

// Warnf calls Golang's log.Printf which in turn calls Output to print to the standard logger.
// It will call log.Printf if the log level is set Warn or above.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	if logLevel < LevelWarn {
		return
	}
	golog.Printf(format, v...)
}

// Info calls Golang's log.Print which in turn calls Output to print to the standard logger.
// It will only call log.Printf if the log level is set Info or above.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	if logLevel < LevelInfo {
		return
	}
	golog.Print(v...)
}

// Infof calls Golang's log.Printf which in turn calls Output to print to the standard logger.
// It will only call log.Printf if the log level is set Info or above.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	if logLevel < LevelInfo {
		return
	}
	golog.Printf(format, v...)
}

// Debug calls Golang's log.Print which in turn calls Output to print to the standard logger.
// It will only call log.Printf if the log level is set Debug or above.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {
	if logLevel < LevelDebug {
		return
	}
	golog.Print(v...)
}

// Debugf calls Golang's log.Printf which in turn calls Output to print to the standard logger.
// It will only call log.Printf if the log level is set Debug or above.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	if logLevel < LevelDebug {
		return
	}
	golog.Printf(format, v...)
}

// Trace calls Golang's log.Print which in turn calls Output to print to the standard logger.
// It will only call log.Printf if the log level is set Trace or above.
// Arguments are handled in the manner of fmt.Print.
func Trace(v ...interface{}) {
	if logLevel < LevelTrace {
		return
	}
	golog.Print(v...)
}

// Tracef calls Golang's log.Printf which in turn calls Output to print to the standard logger.
// It will only call log.Printf if the log level is set Trace or above.
// Arguments are handled in the manner of fmt.Printf.
func Tracef(format string, v ...interface{}) {
	if logLevel < LevelTrace {
		return
	}
	golog.Printf(format, v...)
}