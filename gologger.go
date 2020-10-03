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
		val, e := json.Marshal(msg)
		if e != nil {
			// Do nothing
		}
		if l.Scope != "" {
			_, _ = fmt.Fprint(os.Stdout, string(colorGreen), fmt.Sprintf("INFO: [%s] Scope: [%s] %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
		}
		_, _ = fmt.Fprint(os.Stdout, string(colorGreen), fmt.Sprintf("INFO: [%s]  %s\n", time.Now().Format(time.RFC3339), string(val)))
	}

}

// Warn level log print to console
func (l GoLogger) Warn(msg interface{}) {
	if l.config.LogLevel >= WARN {
		val, e := json.Marshal(msg)
		if e != nil {
			// Do nothing
		}
		if l.Scope != "" {
			_, _ = fmt.Fprint(os.Stdout, string(colorYellow), fmt.Sprintf("Warn: [%s] Scope: [%s]  %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
		}
		_, _ = fmt.Fprint(os.Stdout, string(colorYellow), fmt.Sprintf("Warn: [%s]  %s\n", time.Now().Format(time.RFC3339), string(val)))
	}
}

// Debug level log print to console
func (l GoLogger) Debug(msg interface{}) {
	if l.config.LogLevel >= DEBUG {
		val, e := json.Marshal(msg)
		if e != nil {
			// Do nothing
		}
		if l.Scope != "" {
			_, _ = fmt.Fprint(os.Stdout, string(colorCyan), fmt.Sprintf("DEBUG: [%s] Scope: [%s] %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
		}
		_, _ = fmt.Fprint(os.Stdout, string(colorCyan), fmt.Sprintf("DEBUG: [%s] %s\n", time.Now().Format(time.RFC3339), string(val)))
	}
}

// Error level log print to console
func (l GoLogger) Error(msg interface{}) {

	val, e := json.Marshal(msg)
	if e != nil {
		// Do nothing
	}
	if l.Scope != "" {
		_, _ = fmt.Fprint(os.Stdout, string(colorRed), fmt.Sprintf("ERROR: [%s] Scope: [%s] %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
	}
	_, _ = fmt.Fprint(os.Stdout, string(colorRed), fmt.Sprintf("ERROR: [%s] %s\n", time.Now().Format(time.RFC3339), string(val)))

}
