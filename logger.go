package pterm

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"

	"github.com/pterm/pterm/internal"
)

type LogLevel int

func (l LogLevel) String() string {
	switch l {
	case LogLevelDisabled:
		return ""
	case LogLevelPrint:
		return NewStyle(Bold, FgLightMagenta).Sprint("PRINT")
	case LogLevelTrace:
		return NewStyle(Bold, FgLightCyan).Sprint("TRACE")
	case LogLevelDebug:
		return NewStyle(Bold, FgLightBlue).Sprint("DEBUG")
	case LogLevelInfo:
		return NewStyle(Bold, FgLightGreen).Sprint("INFO ")
	case LogLevelWarn:
		return NewStyle(Bold, FgLightYellow).Sprint("WARN ")
	case LogLevelError:
		return NewStyle(Bold, FgLightRed).Sprint("ERROR")
	case LogLevelFatal:
		return NewStyle(Bold, FgLightRed).Sprint("FATAL")
	}
	return "Unknown"
}

const (
	// LogLevelDisabled does never print.
	LogLevelDisabled LogLevel = iota
	// LogLevelPrint is the log level for printing.
	LogLevelPrint
	// LogLevelTrace is the log level for traces.
	LogLevelTrace
	// LogLevelDebug is the log level for debug.
	LogLevelDebug
	// LogLevelInfo is the log level for info.
	LogLevelInfo
	// LogLevelWarn is the log level for warnings.
	LogLevelWarn
	// LogLevelError is the log level for errors.
	LogLevelError
	// LogLevelFatal is the log level for fatal errors.
	LogLevelFatal
)

type LogFormatter int

const (
	// LogFormatterColorful is a colorful log formatter.
	LogFormatterColorful LogFormatter = iota
	// LogFormatterJSON is a JSON log formatter.
	LogFormatterJSON
)

// DefaultLogger is the default logger.
var DefaultLogger = Logger{
	Formatter:       LogFormatterColorful,
	Writer:          os.Stdout,
	Level:           LogLevelInfo,
	ShowTimestamp:   true,
	TimestampLayout: time.RFC3339,
}

// loggerMutex syncs all loggers, so that they don't print at the exact same time.
var loggerMutex sync.Mutex

type Logger struct {
	Formatter LogFormatter
	Writer    io.Writer
	// Level is the log level of the logger.
	Level LogLevel
	// ShowTimestamp defines if the logger should print a timestamp.
	ShowTimestamp bool
	// TimestampLayout defines the layout of the timestamp.
	TimestampLayout string
	// KeyStyles defines the styles for specific keys.
	KeyStyles map[string]Style
}

// WithLevel sets the log level of the logger.
func (l Logger) WithLevel(level LogLevel) *Logger {
	l.Level = level
	return &l
}

// WithTimestamp enables or disables the timestamp.
func (l Logger) WithTimestamp(b ...bool) *Logger {
	l.ShowTimestamp = internal.WithBoolean(b)
	return &l
}

// CanPrint checks if the logger can print a specific log level.
func (l Logger) CanPrint(level LogLevel) bool {
	return l.Level <= level
}

func (l Logger) print(level LogLevel, args ...any) {
	if l.Level > level {
		return
	}

	var line string

	switch l.Formatter {
	case LogFormatterColorful:
		line = l.renderColorful(level, args...)
	case LogFormatterJSON:
		line = l.renderJSON(level, args...)
	}

	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	_, _ = l.Writer.Write([]byte(line + "\n"))
}

func (l Logger) renderColorful(level LogLevel, args ...any) string {
	return ""
}

func (l Logger) renderJSON(level LogLevel, args ...any) string {
	m := l.argsToMap(args...)

	m["level"] = level.String()
	m["timestamp"] = time.Now().Format(l.TimestampLayout)

	b, _ := json.Marshal(m)
	return string(b)
}

func (l Logger) argsToMap(args ...any) map[string]string {
	// args are always in this order: "key", "value", "key", "value", ...
	m := make(map[string]string)
	for i := 0; i < len(args); i += 2 {
		m[Sprint(args[i])] = Sprint(args[i+1])
	}

	return m
}

func (l Logger) Trace(msg string, args ...any) {
	l.print(LogLevelTrace, args...)
}

func (l Logger) Debug(msg string, args ...any) {
	l.print(LogLevelDebug, args...)
}

func (l Logger) Info(msg string, args ...any) {
	DefaultLogger.print(LogLevelInfo, args...)
}

func (l Logger) Warn(msg string, args ...any) {
	l.print(LogLevelWarn, args...)
}

func (l Logger) Error(msg string, args ...any) {
	l.print(LogLevelError, args...)
}

func (l Logger) Fatal(msg string, args ...any) {
	l.print(LogLevelFatal, args...)
	os.Exit(1)
}

func (l Logger) Print(msg string, args ...any) {
	l.print(LogLevelPrint, args...)
}
