package gologger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Logger Interface
type Logger interface {
	Info(msg interface{})
	Debug(msg interface{})
	Warn(msg interface{})
	Error(msg interface{})
	SetScope(scope string) Logger
}

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	ERROR       = 0
	INFO        = 1
	WARN        = 2
	DEBUG       = 3
)

// GoLogger struct implements all Logger Method
type GoLogger struct {

	// config for our Logger
	config Config

	// Scope of logger
	Scope string
}

// NewGoLogger constructor for GoLogger
func NewGoLogger(config Config) Logger {
	return GoLogger{
		config: config,
		Scope:  "",
	}
}

// SetScope for setting scope of logger
func (l GoLogger) SetScope(scope string) Logger {
	l.Scope = scope
	return l
}

// Info level log print to console
func (l GoLogger) Info(msg interface{}) {
	if l.config.LogLevel >= INFO {

		val, e := marshall(msg, l.config.PrettyPrint)
		if e != nil {
			// Do nothing
		}
		if l.Scope != "" {
			log(INFO, fmt.Sprintf("INFO: [%s] Scope: [%s] %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
		} else {
			log(INFO, fmt.Sprintf("INFO: [%s]  %s\n", time.Now().Format(time.RFC3339), string(val)))
		}
	}

}

// Warn level log print to console
func (l GoLogger) Warn(msg interface{}) {
	if l.config.LogLevel >= WARN {
		val, e := marshall(msg, l.config.PrettyPrint)
		if e != nil {
			// Do nothing
		}
		if l.Scope != "" {
			log(WARN, fmt.Sprintf("Warn: [%s] Scope: [%s]  %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
		} else {
			log(WARN, fmt.Sprintf("Warn: [%s]  %s\n", time.Now().Format(time.RFC3339), string(val)))
		}
	}
}

// Debug level log print to console
func (l GoLogger) Debug(msg interface{}) {
	if l.config.LogLevel >= DEBUG {
		val, e := marshall(msg, l.config.PrettyPrint)
		if e != nil {
			// Do nothing
		}
		if l.Scope != "" {
			log(DEBUG, fmt.Sprintf("DEBUG: [%s] Scope: [%s] %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
		} else {
			log(DEBUG, fmt.Sprintf("DEBUG: [%s] %s\n", time.Now().Format(time.RFC3339), string(val)))
		}
	}
}

// Error level log print to console
func (l GoLogger) Error(msg interface{}) {

	val, e := marshall(msg, l.config.PrettyPrint)
	if e != nil {
		// Do nothing
	}
	if l.Scope != "" {
		log(ERROR, fmt.Sprintf("ERROR: [%s] Scope: [%s] %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
	} else {
		log(ERROR, fmt.Sprintf("ERROR: [%s] %s\n", time.Now().Format(time.RFC3339), string(val)))
	}

}


// log for final logging
func log(lvl uint8, msg string) {
	var color string
	if lvl == ERROR {
		color = colorRed
	}
	if lvl == INFO {
		color = colorGreen
	}
	if lvl == WARN {
		color = colorYellow
	}
	if lvl == DEBUG {
		color = colorCyan
	}
	_, _ = fmt.Fprint(os.Stdout, color, msg)
}

func marshall (msg interface{}, pretty bool) ([]byte, error)  {
	var val []byte
	var e error
	if pretty == true {
		val, e = json.MarshalIndent(msg, "", " ")
	} else {
		val, e = json.Marshal(msg)
	}
	return val, e
}
