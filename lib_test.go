package log

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	// given
	levels := []Level{LevelTrace, LevelDebug, LevelInfo, LevelWarn, LevelError, LevelFatal}

	for _, level := range levels {
		// when
		SetLevel(level)

		// then
		if LogLevel != level {
			t.Errorf("Expected %s, bug got %s", level, LogLevel)
		}
	}
}

func TestParseLevelLowerCase(t *testing.T) {
	// given
	levels := map[string]Level{
		"trace":     LevelTrace,
		"tracei":    LevelTrace,
		"tracein":   LevelTrace,
		"traceing":  LevelTrace,
		"debug":     LevelDebug,
		"debugg":    LevelDebug,
		"debuggi":   LevelDebug,
		"debuggin":  LevelDebug,
		"debugging": LevelDebug,
		"info":      LevelInfo,
		"warn":      LevelWarn,
		"warni":     LevelWarn,
		"warnin":    LevelWarn,
		"warning":   LevelWarn,
		"error":     LevelError,
		"fatal":     LevelFatal,
	}

	for str, level := range levels {
		// when
		result := ParseLevel(str)

		// then
		if result != level {
			t.Errorf("Expected %s, bug got %s", level, result)
		}
	}
}

func TestParseLevelUpperCase(t *testing.T) {
	// given
	levels := map[string]Level{
		"TRACE":     LevelTrace,
		"TRACEI":    LevelTrace,
		"TRACEIN":   LevelTrace,
		"TRACEING":  LevelTrace,
		"DEBUG":     LevelDebug,
		"DEBUGG":    LevelDebug,
		"DEBUGGI":   LevelDebug,
		"DEBUGGIN":  LevelDebug,
		"DEBUGGING": LevelDebug,
		"INFO":      LevelInfo,
		"WARN":      LevelWarn,
		"WARNI":     LevelWarn,
		"WARNIN":    LevelWarn,
		"WARNING":   LevelWarn,
		"ERROR":     LevelError,
		"FATAL":     LevelFatal,
	}

	for str, level := range levels {
		// when
		result := ParseLevel(str)

		// then
		if result != level {
			t.Errorf("Expected %s, bug got %s", level, result)
		}
	}
}
