package pterm

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm/internal"
)

type LogLevel int

// Style returns the style of the log level.
func (l LogLevel) Style() Style {
	baseStyle := NewStyle(Bold)
	switch l {
	case LogLevelTrace:
		return baseStyle.Add(*FgCyan.ToStyle())
	case LogLevelDebug:
		return baseStyle.Add(*FgBlue.ToStyle())
	case LogLevelInfo:
		return baseStyle.Add(*FgGreen.ToStyle())
	case LogLevelWarn:
		return baseStyle.Add(*FgYellow.ToStyle())
	case LogLevelError:
		return baseStyle.Add(*FgRed.ToStyle())
	case LogLevelFatal:
		return baseStyle.Add(*FgRed.ToStyle())
	case LogLevelPrint:
		return baseStyle.Add(*FgWhite.ToStyle())
	}

	return baseStyle.Add(*FgWhite.ToStyle())
}

func (l LogLevel) String() string {
	switch l {
	case LogLevelDisabled:
		return ""
	case LogLevelTrace:
		return "TRACE"
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	case LogLevelPrint:
		return "PRINT"
	}
	return "Unknown"
}

const (
	// LogLevelDisabled does never print.
	LogLevelDisabled LogLevel = iota
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
	// LogLevelPrint is the log level for printing.
	LogLevelPrint
)

// LogFormatter is the log formatter.
// Can be either LogFormatterColorful or LogFormatterJSON.
type LogFormatter int

const (
	// LogFormatterColorful is a colorful log formatter.
	LogFormatterColorful LogFormatter = iota
	// LogFormatterJSON is a JSON log formatter.
	LogFormatterJSON
)

// DefaultLogger is the default logger.
var DefaultLogger = Logger{
	Formatter:  LogFormatterColorful,
	Writer:     os.Stdout,
	Level:      LogLevelInfo,
	ShowTime:   true,
	TimeFormat: "2006-01-02 15:04:05",
	MaxWidth:   80,
	KeyStyles: map[string]Style{
		"error":  *NewStyle(FgRed, Bold),
		"err":    *NewStyle(FgRed, Bold),
		"caller": *NewStyle(FgGray, Bold),
	},
}

// loggerMutex syncs all loggers, so that they don't print at the exact same time.
var loggerMutex sync.Mutex

type Logger struct {
	// Formatter is the log formatter of the logger.
	Formatter LogFormatter
	// Writer is the writer of the logger.
	Writer io.Writer
	// Level is the log level of the logger.
	Level LogLevel
	// ShowCaller defines if the logger should print the caller.
	ShowCaller bool
	// CallerOffset defines the offset of the caller.
	CallerOffset int
	// ShowTime defines if the logger should print a timestamp.
	ShowTime bool
	// TimestampLayout defines the layout of the timestamp.
	TimeFormat string
	// KeyStyles defines the styles for specific keys.
	KeyStyles map[string]Style
	// MaxWidth defines the maximum width of the logger.
	// If the text (including the arguments) is longer than the max width, it will be split into multiple lines.
	MaxWidth int
}

// WithFormatter sets the log formatter of the logger.
func (l Logger) WithFormatter(formatter LogFormatter) *Logger {
	l.Formatter = formatter
	return &l
}

// WithWriter sets the writer of the logger.
func (l Logger) WithWriter(writer io.Writer) *Logger {
	l.Writer = writer
	return &l
}

// WithLevel sets the log level of the logger.
func (l Logger) WithLevel(level LogLevel) *Logger {
	l.Level = level
	return &l
}

// WithCaller enables or disables the caller.
func (l Logger) WithCaller(b ...bool) *Logger {
	l.ShowCaller = internal.WithBoolean(b)
	return &l
}

// WithCallerOffset sets the caller offset.
func (l Logger) WithCallerOffset(offset int) *Logger {
	l.CallerOffset = offset
	return &l
}

// WithTime enables or disables the timestamp.
func (l Logger) WithTime(b ...bool) *Logger {
	l.ShowTime = internal.WithBoolean(b)
	return &l
}

// WithTimeFormat sets the timestamp layout.
func (l Logger) WithTimeFormat(format string) *Logger {
	l.TimeFormat = format
	return &l
}

// WithKeyStyles sets the style for a specific key.
func (l Logger) WithKeyStyles(styles map[string]Style) *Logger {
	l.KeyStyles = styles
	return &l
}

// WithMaxWidth sets the maximum width of the logger.
func (l Logger) WithMaxWidth(width int) *Logger {
	l.MaxWidth = width
	return &l
}

// AppendKeyStyles appends a style for a specific key.
func (l Logger) AppendKeyStyles(styles map[string]Style) *Logger {
	for k, v := range styles {
		l.KeyStyles[k] = v
	}
	return &l
}

// AppendKeyStyle appends a style for a specific key.
func (l Logger) AppendKeyStyle(key string, style Style) *Logger {
	l.KeyStyles[key] = style
	return &l
}

// CanPrint checks if the logger can print a specific log level.
func (l Logger) CanPrint(level LogLevel) bool {
	if l.Level == LogLevelDisabled {
		return false
	}
	return l.Level <= level
}

// Args converts any arguments to a slice of LoggerArgument.
func (l Logger) Args(args ...any) []LoggerArgument {
	var loggerArgs []LoggerArgument

	// args are in the format of: key, value, key, value, key, value, ...
	for i := 0; i < len(args); i += 2 {
		key := Sprint(args[i])
		value := args[i+1]

		loggerArgs = append(loggerArgs, LoggerArgument{
			Key:   key,
			Value: value,
		})
	}

	return loggerArgs
}

// ArgsFromMap converts a map to a slice of LoggerArgument.
func (l Logger) ArgsFromMap(m map[string]any) []LoggerArgument {
	var loggerArgs []LoggerArgument

	for k, v := range m {
		loggerArgs = append(loggerArgs, LoggerArgument{
			Key:   k,
			Value: v,
		})
	}

	return loggerArgs
}

func (l Logger) getCallerInfo() (path string, line int) {
	if !l.ShowCaller {
		return
	}

	_, path, line, _ = runtime.Caller(l.CallerOffset + 4)
	_, callerBase, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(callerBase)
	basepath = strings.ReplaceAll(basepath, "\\", "/")

	path = strings.TrimPrefix(path, basepath)

	return
}

func (l Logger) combineArgs(args ...[]LoggerArgument) []LoggerArgument {
	var result []LoggerArgument

	for _, arg := range args {
		result = append(result, arg...)
	}

	return result
}

func (l Logger) print(level LogLevel, msg string, args []LoggerArgument) {
	if !l.CanPrint(level) {
		return
	}

	var line string

	switch l.Formatter {
	case LogFormatterColorful:
		line = l.renderColorful(level, msg, args)
	case LogFormatterJSON:
		line = l.renderJSON(level, msg, args)
	}

	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	_, _ = l.Writer.Write([]byte(line + "\n"))
}

func (l Logger) renderColorful(level LogLevel, msg string, args []LoggerArgument) (result string) {
	if l.ShowTime {
		result += Gray(time.Now().Format(l.TimeFormat)) + " "
	}

	if GetTerminalWidth() > 0 && GetTerminalWidth() < l.MaxWidth {
		l.MaxWidth = GetTerminalWidth()
	}

	var argumentsInNewLine bool

	result += level.Style().Sprintf("%-5s", level.String()) + " "

	// if msg is too long, wrap it to multiple lines with the same length
	remainingWidth := l.MaxWidth - internal.GetStringMaxWidth(result)
	if internal.GetStringMaxWidth(msg) > remainingWidth {
		argumentsInNewLine = true
		msg = DefaultParagraph.WithMaxWidth(remainingWidth).Sprint(msg)
		padding := len(time.Now().Format(l.TimeFormat) + " ")
		msg = strings.ReplaceAll(msg, "\n", "\n"+strings.Repeat(" ", padding)+"  │   ")
	}

	result += msg

	if l.ShowCaller {
		path, line := l.getCallerInfo()
		args = append(args, LoggerArgument{
			Key:   "caller",
			Value: FgGray.Sprintf("%s:%d", path, line),
		})
	}

	arguments := make([]string, len(args))

	// add arguments
	if len(args) > 0 {
		for i, arg := range args {
			if style, ok := l.KeyStyles[arg.Key]; ok {
				arguments[i] = style.Sprintf("%s: ", arg.Key)
			} else {
				arguments[i] = level.Style().Sprintf("%s: ", arg.Key)
			}

			arguments[i] += Sprintf("%s", Sprint(arg.Value))
		}
	}

	fullLine := result + " " + strings.Join(arguments, " ")

	// if the full line is too long, wrap the arguments to multiple lines
	if internal.GetStringMaxWidth(fullLine) > l.MaxWidth {
		argumentsInNewLine = true
	}

	if !argumentsInNewLine {
		result = fullLine
	} else {
		padding := 4
		if l.ShowTime {
			padding = len(time.Time{}.Format(l.TimeFormat)) + 3
		}

		for i, argument := range arguments {
			var pipe string
			if i < len(arguments)-1 {
				pipe = "├"
			} else {
				pipe = "└"
			}
			result += "\n" + strings.Repeat(" ", padding) + pipe + " " + argument
		}
	}

	return
}

func (l Logger) renderJSON(level LogLevel, msg string, args []LoggerArgument) string {
	m := l.argsToMap(args)

	m["level"] = level.String()
	m["timestamp"] = time.Now().Format(l.TimeFormat)
	m["msg"] = msg

	if file, line := l.getCallerInfo(); file != "" {
		m["caller"] = Sprintf("%s:%d", file, line)
	}

	b, _ := json.Marshal(m)
	return string(b)
}

func (l Logger) argsToMap(args []LoggerArgument) map[string]any {
	m := make(map[string]any)

	for _, arg := range args {
		m[arg.Key] = arg.Value
	}

	return m
}

// Trace prints a trace log.
func (l Logger) Trace(msg string, args ...[]LoggerArgument) {
	l.print(LogLevelTrace, msg, l.combineArgs(args...))
}

// Debug prints a debug log.
func (l Logger) Debug(msg string, args ...[]LoggerArgument) {
	l.print(LogLevelDebug, msg, l.combineArgs(args...))
}

// Info prints an info log.
func (l Logger) Info(msg string, args ...[]LoggerArgument) {
	l.print(LogLevelInfo, msg, l.combineArgs(args...))
}

// Warn prints a warning log.
func (l Logger) Warn(msg string, args ...[]LoggerArgument) {
	l.print(LogLevelWarn, msg, l.combineArgs(args...))
}

// Error prints an error log.
func (l Logger) Error(msg string, args ...[]LoggerArgument) {
	l.print(LogLevelError, msg, l.combineArgs(args...))
}

// Fatal prints a fatal log and exits the program.
func (l Logger) Fatal(msg string, args ...[]LoggerArgument) {
	l.print(LogLevelFatal, msg, l.combineArgs(args...))
	if l.CanPrint(LogLevelFatal) {
		os.Exit(1)
	}
}

// Print prints a log.
func (l Logger) Print(msg string, args ...[]LoggerArgument) {
	l.print(LogLevelPrint, msg, l.combineArgs(args...))
}

// LoggerArgument is a key-value pair for a logger.
type LoggerArgument struct {
	// Key is the key of the argument.
	Key string
	// Value is the value of the argument.
	Value any
}
