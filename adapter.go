/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import (
	_log "log"
)

type (
	adapter struct {
		level  Level
		logger *Logger
	}
)

func (a *adapter) Write(p []byte) (int, error) {
	a.logger.write(a.level, string(p))
	return len(p), nil
}

// NewStdLogger AFAIRE.
func (log *Logger) NewStdLogger(level, prefix string, flag int) *_log.Logger {
	a := &adapter{
		level:  StringToLevel(level),
		logger: log,
	}

	return _log.New(a, prefix, flag)
}

/*
######################################################################################################## @(°_°)@ #######
*/
