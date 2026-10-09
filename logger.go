package pterm

import (
	"encoding/json"
	"io"
	"log/slog"
	"maps"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm/internal"
)

// LogLevel is the severity level used by the Logger.
type LogLevel int

// Style returns the style of the log level, as configured in ThemeDefault.
func (l LogLevel) Style() Style {
	switch l {
	case LogLevelTrace:
		return ThemeDefault.LoggerTraceStyle
	case LogLevelDebug:
		return ThemeDefault.LoggerDebugStyle
	case LogLevelInfo:
		return ThemeDefault.LoggerInfoStyle
	case LogLevelWarn:
		return ThemeDefault.LoggerWarnStyle
	case LogLevelError:
		return ThemeDefault.LoggerErrorStyle
	case LogLevelFatal:
		return ThemeDefault.LoggerFatalStyle
	case LogLevelPrint:
		return ThemeDefault.LoggerPrintStyle
	}

	return ThemeDefault.LoggerPrintStyle
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

// Logger is a fully configurable, structured logger.
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
	// The width is always capped to the current terminal width.
	// A value of zero or less uses the full terminal width.
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
	maps.Copy(l.KeyStyles, styles)

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
	args = l.sanitizeArgs(args)

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
	loggerArgs := make([]LoggerArgument, 0, len(m))

	for k, v := range m {
		loggerArgs = append(loggerArgs, LoggerArgument{
			Key:   k,
			Value: v,
		})
	}

	return loggerArgs
}

// sanitizeArgs inserts an error message into an args slice if an odd number of arguments is provided.
func (l Logger) sanitizeArgs(args []any) []any {
	numArgs := len(args)
	if numArgs > 0 && numArgs%2 != 0 {
		if numArgs > 1 {
			lastArg := args[numArgs-1]
			args = append(args[:numArgs-1], []any{ErrKeyWithoutValue, lastArg}...)
		} else {
			args = []any{ErrKeyWithoutValue, args[0]}
		}
	}

	return args
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

	Fprintln(l.Writer, line)
}

// loggerLevelWidth is the width the level name is padded to, so that the
// messages of all levels start at the same column.
const loggerLevelWidth = 5

// loggerMinContentWidth is the narrowest the wrapped content is allowed to
// get. In terminals too narrow to hold the prefix plus this width, lines
// overflow instead of degrading into one word per line.
const loggerMinContentWidth = 16

// lineWidth returns the width a rendered log line may occupy: MaxWidth capped
// to the current terminal width. A MaxWidth of zero or less means the full
// terminal width may be used.
func (l Logger) lineWidth() int {
	width := l.MaxWidth

	if terminalWidth := GetTerminalWidth(); terminalWidth > 0 && (width <= 0 || terminalWidth < width) {
		width = terminalWidth
	}

	return width
}

// quoteValue wraps a value that contains spaces in dimmed quotes, so that
// inline key=value pairs stay unambiguous.
func (l Logger) quoteValue(value string) string {
	stripped := internal.RemoveEscapeCodes(value)

	if stripped == "" || (strings.Contains(stripped, " ") && !strings.Contains(stripped, "\n")) {
		return Gray(`"`) + value + Gray(`"`)
	}

	return value
}

func (l Logger) renderColorful(level LogLevel, msg string, args []LoggerArgument) string {
	var prefix string

	if l.ShowTime {
		prefix += ThemeDefault.LoggerTimestampStyle.Sprint(time.Now().Format(l.TimeFormat)) + " "
	}

	prefix += level.Style().Sprintf("%-*s", loggerLevelWidth, level.String()) + " "

	if l.ShowCaller {
		path, line := l.getCallerInfo()
		args = append(args, LoggerArgument{
			Key:   "caller",
			Value: ThemeDefault.LoggerCallerStyle.Sprintf("%s:%d", path, line),
		})
	}

	keys := make([]string, len(args))
	values := make([]string, len(args))

	// Keys take the level color; the FATAL background is too heavy to repeat
	// on every key, so its keys fall back to plain red.
	keyStyle := level.Style()
	if level == LogLevelFatal {
		keyStyle = ThemeDefault.LoggerFatalKeyStyle
	}

	var multilineValues bool

	for i, arg := range args {
		style, ok := l.KeyStyles[arg.Key]
		if !ok {
			style = keyStyle
		}

		keys[i] = style.Sprint(arg.Key) + Gray("=")
		values[i] = Sprint(arg.Value)

		if strings.Contains(values[i], "\n") {
			multilineValues = true
		}
	}

	width := l.lineWidth()

	var inlineBuilder strings.Builder

	inlineBuilder.WriteString(prefix)
	inlineBuilder.WriteString(msg)

	for i := range keys {
		inlineBuilder.WriteString(" ")
		inlineBuilder.WriteString(keys[i])
		inlineBuilder.WriteString(l.quoteValue(values[i]))
	}

	inline := inlineBuilder.String()

	// Raw output always stays on a single line, so it remains grep-friendly
	// when piped into files or other tools.
	if rawOutput() {
		return inline
	}

	if !multilineValues && !strings.Contains(msg, "\n") && (width <= 0 || internal.GetStringMaxWidth(inline) <= width) {
		return inline
	}

	return l.renderBlock(prefix, msg, keys, values, width)
}

// renderBlock renders a log whose inline form would overflow the line width:
// the message wraps under its own first line and every argument moves onto its
// own line, connected by a dimmed tree rail. Wrapped and multiline argument
// values align under the start of the value.
func (l Logger) renderBlock(prefix, msg string, keys, values []string, width int) string {
	prefixWidth := internal.GetStringMaxWidth(prefix)
	indent := strings.Repeat(" ", prefixWidth)

	contentWidth := 0
	if width > 0 {
		contentWidth = max(width-prefixWidth, loggerMinContentWidth)
	}

	var sb strings.Builder

	sb.WriteString(prefix)

	for i, line := range internal.WrapText(msg, contentWidth) {
		if i > 0 {
			sb.WriteString("\n" + indent)
		}

		sb.WriteString(line)
	}

	for i := range keys {
		connector, rail := Gray("├ "), Gray("│ ")
		if i == len(keys)-1 {
			connector, rail = Gray("└ "), "  "
		}

		valueIndent := internal.GetStringMaxWidth(keys[i])
		if contentWidth > 0 && valueIndent > contentWidth/2 {
			valueIndent = 2
		}

		valueWidth := 0
		if contentWidth > 0 {
			valueWidth = contentWidth - 2 - valueIndent
		}

		valueLines := internal.WrapText(values[i], valueWidth)

		sb.WriteString("\n" + indent + connector + keys[i] + valueLines[0])

		for _, line := range valueLines[1:] {
			sb.WriteString("\n" + indent + rail + strings.Repeat(" ", valueIndent) + line)
		}
	}

	return sb.String()
}

func (l Logger) renderJSON(level LogLevel, msg string, args []LoggerArgument) string {
	m := l.argsToMap(args)

	m["level"] = level.String()
	m["timestamp"] = time.Now().Format(l.TimeFormat)
	m["msg"] = msg

	if file, line := l.getCallerInfo(); file != "" {
		m["caller"] = Sprintf("%s:%d", file, line)
	}

	b, err := json.Marshal(m)
	if err != nil {
		return Sprintf("%v", m)
	}

	return string(b)
}

func (l Logger) argsToMap(args []LoggerArgument) map[string]any {
	m := make(map[string]any)

	for _, arg := range args {
		v := arg.Value
		if sv, ok := arg.Value.(slog.Value); ok {
			v = sv.Any()
		}

		m[arg.Key] = v
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
