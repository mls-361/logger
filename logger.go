/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import (
	"log"

	"github.com/mls-361/buffer"
)

const (
	_bufSize = 256
)

var (
	_bufPool = buffer.NewPool(_bufSize)
)

type (
	// Logger AFAIRE.
	Logger interface {
		Trace(msg string, data ...interface{})
		Debug(msg string, data ...interface{})
		Info(msg string, data ...interface{})
		Notice(msg string, data ...interface{})
		Warning(msg string, data ...interface{})
		Error(msg string, data ...interface{})
		Fatal(msg string, data ...interface{})
		SetLevel(level string)
		CreateLogger(id, name string) Logger
		Remove()
		NewStdLogger(level, prefix string, flag int) *log.Logger
	}

	logger struct {
		id      string
		name    string
		level   Level
		output  Output
		loggers *loggers
	}
)

func (log *logger) write(level Level, msg string, data ...interface{}) {
	if level < log.loadLevel() {
		return
	}

	buf := _bufPool.Get()

	_ = log.output.Write(buf, log, level, msg, data...) // AFINIR

	_bufPool.Put(buf)
}

// Trace AFAIRE.
func (log *logger) Trace(msg string, data ...interface{}) {
	log.write(LevelTrace, msg, data...)
}

// Debug AFAIRE.
func (log *logger) Debug(msg string, data ...interface{}) {
	log.write(LevelDebug, msg, data...)
}

// Info AFAIRE.
func (log *logger) Info(msg string, data ...interface{}) {
	log.write(LevelInfo, msg, data...)
}

// Notice AFAIRE.
func (log *logger) Notice(msg string, data ...interface{}) {
	log.write(LevelNotice, msg, data...)
}

// Warning AFAIRE.
func (log *logger) Warning(msg string, data ...interface{}) {
	log.write(LevelWarning, msg, data...)
}

// Error AFAIRE.
func (log *logger) Error(msg string, data ...interface{}) {
	log.write(LevelError, msg, data...)
}

// Fatal AFAIRE.
func (log *logger) Fatal(msg string, data ...interface{}) {
	log.write(LevelFatal, msg, data...)
}

// SetLevel AFAIRE.
func (log *logger) SetLevel(level string) {
	log.storeLevel(StringToLevel(level))
}

// CreateLogger AFAIRE.
func (log *logger) CreateLogger(id, name string) *logger {
	logger := &logger{
		id:      id,
		name:    name,
		level:   log.loadLevel(),
		output:  log.output,
		loggers: log.loggers,
	}

	log.loggers.add(logger)

	return logger
}

// Remove AFAIRE.
func (log *logger) Remove() {
	log.loggers.remove(log.id)
}

/*
######################################################################################################## @(°_°)@ #######
*/
