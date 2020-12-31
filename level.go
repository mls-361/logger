/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import "sync/atomic"

type (
	Level int64
)

const (
	// LevelTrace AFAIRE.
	LevelTrace Level = iota
	// LevelDebug AFAIRE.
	LevelDebug
	// LevelInfo AFAIRE.
	LevelInfo
	// LevelNotice AFAIRE.
	LevelNotice
	// LevelWarning AFAIRE.
	LevelWarning
	// LevelError AFAIRE.
	LevelError
	// LevelFatal AFAIRE.
	LevelFatal
)

// StringToLevel AFAIRE.
func StringToLevel(level string) Level {
	switch level {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "notice":
		return LevelNotice
	case "warning":
		return LevelWarning
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	default:
		return LevelTrace
	}
}

func (log *Logger) loadLevel() Level {
	return Level(atomic.LoadInt64((*int64)(&log.level)))
}

func (log *Logger) storeLevel(level Level) {
	atomic.StoreInt64((*int64)(&log.level), int64(level))
}

/*
######################################################################################################## @(°_°)@ #######
*/
