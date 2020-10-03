package gologger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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

type GoLogger struct {
	LogLevel  uint8
	FileFlush bool
	Scope     string
}

func NewGoLogger(logLevel uint8, fileFlush bool) Logger {
	return GoLogger{
		LogLevel:  logLevel,
		FileFlush: fileFlush,
		Scope:"",
	}
}


func (l GoLogger) SetScope(scope string) Logger {
	l.Scope = scope
	return l
}

func (l GoLogger) Info(msg interface{}) {
	if l.LogLevel >= INFO {
		val, e := json.Marshal(msg)
		if e != nil {
			// Do nothing
		}
		if l.Scope != "" {
			_, _ = fmt.Fprint(os.Stdout, string(colorGreen), fmt.Sprintf("INFO: [%s] Scope: [%s] %s\n", time.Now().Format(time.RFC3339), l.Scope, string(val)))
		}
		_, _ = fmt.Fprint(os.Stdout, string(colorGreen), fmt.Sprintf("INFO: [%s]  %s\n", time.Now().Format(time.RFC3339),  string(val)))
	}

}

func (l GoLogger) Warn(msg interface{}) {
	if l.LogLevel >= WARN {
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

func (l GoLogger) Debug(msg interface{}) {
	if l.LogLevel >= DEBUG {
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
