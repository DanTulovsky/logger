package logger

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// LogLevel is a log level
type LogLevel int

const (
	levelDebug LogLevel = iota
	levelInfo
	levelWarn
	levelError
	levelFatal
)

var (
	level = flag.Int("log_level", int(levelInfo), "log level")
)

// Logger is the custom logger
type Logger struct {
	color  *color.Color
	prefix string
	level  LogLevel
	logger *log.Logger
}

// New returns a new logger
func New(prefix string, color *color.Color) *Logger {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	return &Logger{
		prefix: prefix,
		color:  color,
		level:  LogLevel(*level),
		logger: logger,
	}
}

// SetLevel sets the logger level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// SetColor sets the logger color
func (l *Logger) SetColor(clr *color.Color) {
	l.color = clr
}

// Log logs a message
func (l *Logger) Log(format string, v ...interface{}) {
	prefix := l.color.Sprintf("%v ", l.prefix)
	// The 2 gives us the line that called this function
	l.logger.Output(3, fmt.Sprintf(prefix+format, v...))
}

// Debug logs levelDebug and higher messages
func (l *Logger) Debug(v ...interface{}) {
	if l.level > levelDebug {
		return
	}

	l.Log("%v", v...)
}

// Info logs levelInfo and higher messages
func (l *Logger) Info(v ...interface{}) {
	if l.level > levelInfo {
		return
	}

	l.Log("%v", v...)
}

// Warn logs levelWarn and higher messages
func (l *Logger) Warn(v ...interface{}) {
	if l.level > levelWarn {
		return
	}

	l.Log("%v", v...)
}

// Error logs levelError and higher messages
func (l *Logger) Error(v ...interface{}) {
	if l.level > levelError {
		return
	}

	l.Log("%v", v...)
}

// Fatal logs levelFatal and higher messages
func (l *Logger) Fatal(v ...interface{}) {
	if l.level > levelFatal {
		return
	}

	l.Log("%v", v...)
	os.Exit(1)
}

// Debugf logs levelDebug and higher messages
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level > levelDebug {
		return
	}

	l.Log(format, v...)
}

// Infof logs levelInfo and higher messages
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level > levelInfo {
		return
	}

	l.Log(format, v...)
}

// Warnf logs levelWarn and higher messages
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level > levelWarn {
		return
	}

	l.Log(format, v...)
}

// Errorf logs levelError and higher messages
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level > levelError {
		return
	}

	l.Log(format, v...)
}

// Fatalf logs levelFatal and higher messages
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.level > levelFatal {
		return
	}

	l.Log(format, v...)
	os.Exit(1)
}
