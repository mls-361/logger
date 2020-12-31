/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import "github.com/mls-361/buffer"

const (
	_bufSize = 256
)

var (
	_bufPool = buffer.NewPool(_bufSize)
)

type (
	// Logger AFAIRE.
	Logger struct {
		id      string
		name    string
		level   Level
		output  Output
		loggers *loggers
	}
)

func (log *Logger) write(level Level, msg string, data ...interface{}) {
	if level < log.loadLevel() {
		return
	}

	buf := _bufPool.Get()

	_ = log.output.Write(buf, log, level, msg, data...) // AFINIR

	_bufPool.Put(buf)
}

// Trace AFAIRE.
func (log *Logger) Trace(msg string, data ...interface{}) {
	log.write(LevelTrace, msg, data...)
}

// Debug AFAIRE.
func (log *Logger) Debug(msg string, data ...interface{}) {
	log.write(LevelDebug, msg, data...)
}

// Info AFAIRE.
func (log *Logger) Info(msg string, data ...interface{}) {
	log.write(LevelInfo, msg, data...)
}

// Notice AFAIRE.
func (log *Logger) Notice(msg string, data ...interface{}) {
	log.write(LevelNotice, msg, data...)
}

// Warning AFAIRE.
func (log *Logger) Warning(msg string, data ...interface{}) {
	log.write(LevelWarning, msg, data...)
}

// Error AFAIRE.
func (log *Logger) Error(msg string, data ...interface{}) {
	log.write(LevelError, msg, data...)
}

// Fatal AFAIRE.
func (log *Logger) Fatal(msg string, data ...interface{}) {
	log.write(LevelFatal, msg, data...)
}

// SetLevel AFAIRE.
func (log *Logger) SetLevel(level string) {
	log.storeLevel(StringToLevel(level))
}

// CreateChild AFAIRE.
func (log *Logger) CreateLogger(id, name string) *Logger {
	logger := &Logger{
		id:      id,
		name:    name,
		level:   log.loadLevel(),
		output:  log.output,
		loggers: log.loggers,
	}

	log.loggers.add(logger)

	return logger
}

// RemoveChild AFAIRE.
func (log *Logger) RemoveLogger(id string) {
	if id == "" {
		id = log.id
	}

	log.loggers.remove(id)
}

/*
######################################################################################################## @(°_°)@ #######
*/
