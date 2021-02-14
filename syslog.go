/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import (
	"log/syslog"

	"github.com/mls-361/buffer"
	"github.com/mls-361/logfmt"
)

const (
	_maxIDLen        = 8
	_maxNameLen      = 10
	_defaultFacility = "local4"
)

type (
	// OutputSyslog AFAIRE.
	OutputSyslog struct {
		*syslog.Writer
	}
)

// NewOutputSyslog AFAIRE.
func NewOutputSyslog(facility, appName string) (*OutputSyslog, error) {
	var p syslog.Priority

	switch facility {
	case "local1":
		p = syslog.LOG_LOCAL1
	case "local2":
		p = syslog.LOG_LOCAL2
	case "local3":
		p = syslog.LOG_LOCAL3
	case "local4":
		p = syslog.LOG_LOCAL4
	case "local5":
		p = syslog.LOG_LOCAL5
	case "local6":
		p = syslog.LOG_LOCAL6
	case "local7":
		p = syslog.LOG_LOCAL7
	default:
		p = syslog.LOG_LOCAL0
	}

	writer, err := syslog.New(p, appName)
	if err != nil {
		return nil, err
	}

	o := &OutputSyslog{
		Writer: writer,
	}

	return o, nil
}

// Write AFAIRE.
func (o *OutputSyslog) Write(buf *buffer.Buffer, log *Logger, level Level, msg string, data ...interface{}) error {
	buf.AppendRJString(log.name, '.', _maxNameLen)
	buf.AppendByte(':')
	buf.AppendRJString(log.id, '.', _maxIDLen)
	buf.AppendByte(' ')
	buf.AppendString(msg)

	if len(data) > 0 {
		buf.AppendString("> ")
		logfmt.EncodeList(buf, data...)
	}

	var err error

	switch level {
	case LevelInfo:
		err = o.Info(buf.String())
	case LevelNotice:
		err = o.Notice(buf.String())
	case LevelWarning:
		err = o.Warning(buf.String())
	case LevelError:
		err = o.Err(buf.String())
	case LevelFatal:
		err = o.Crit(buf.String())
	default:
		err = o.Debug(buf.String())
	}

	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/
