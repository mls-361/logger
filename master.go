/*
------------------------------------------------------------------------------------------------------------------------
####### logger ####### (c) 2020-2021 mls-361 ####################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package logger

import "sync"

type (
	loggers struct {
		all   map[string]*logger
		mutex sync.Mutex
	}
)

func (ls *loggers) add(log *logger) {
	ls.mutex.Lock()
	ls.all[log.id] = log
	ls.mutex.Unlock()
}

func (ls *loggers) remove(id string) {
	ls.mutex.Lock()
	delete(ls.all, id)
	ls.mutex.Unlock()
}

type (
	// Master AFAIRE.
	Master struct {
		*logger
		output  Output
		loggers *loggers
	}
)

// New AFAIRE.
func New() *Master {
	return &Master{
		loggers: &loggers{
			all: make(map[string]*logger),
		},
	}
}

// Build AFAIRE.
func (m *Master) Build(id, name, level string, output Output) error {
	if output == nil {
		var err error

		output, err = NewOutputSyslog(_defaultFacility, name)
		if err != nil {
			return err
		}
	}

	m.output = output

	m.logger = &logger{
		id:      id,
		name:    name,
		level:   StringToLevel(level),
		output:  output,
		loggers: m.loggers,
	}

	m.loggers.add(m.logger)

	return nil
}

// Close AFAIRE.
func (m *Master) Close() {
	if m.output != nil {
		_ = m.output.Close() // AFINIR
		m.output = nil
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
