package pterm

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"maps"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"time"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/screen"
	"github.com/pterm/pterm/internal/text"
)

// LogLevel is the severity of a log line.
type LogLevel int

// The log levels, from least to most severe. A logger prints the lines of
// its own level and every level above it. LogLevelPrint is always printed
// unless the logger is disabled.
const (
	// LogLevelDisabled never prints.
	LogLevelDisabled LogLevel = iota
	// LogLevelTrace is for very detailed output while tracking down a bug.
	LogLevelTrace
	// LogLevelDebug is for output that helps while developing.
	LogLevelDebug
	// LogLevelInfo is for normal operation.
	LogLevelInfo
	// LogLevelWarn is for things to double check.
	LogLevelWarn
	// LogLevelError is for failures.
	LogLevelError
	// LogLevelFatal is for failures that end the program.
	LogLevelFatal
	// LogLevelPrint is for lines that are always printed.
	LogLevelPrint
)

// String returns the name of l in upper case, as JSON output uses it.
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

	return "UNKNOWN"
}

// LogFormatter is the output format of a logger.
type LogFormatter int

const (
	// LogFormatterColorful prints lines for humans.
	LogFormatterColorful LogFormatter = iota
	// LogFormatterJSON prints one JSON object per line, for machines.
	LogFormatterJSON
)

// LevelLabel is how a log level looks at the start of a line.
type LevelLabel struct {
	// Glyph is drawn first. Empty uses the glyph of the level from the
	// active glyph set.
	Glyph string
	// Word names the level, such as "warn".
	Word string
	// Tone colors the glyph and the word.
	Tone Tone
}

// ErrKeyWithoutValue is the key Args uses for a value whose key is missing,
// when it gets an odd number of arguments.
const ErrKeyWithoutValue = "ERROR: key_without_value"

// LoggerArgument is a key-value pair of a log line.
type LoggerArgument struct {
	// Key names the value.
	Key string
	// Value is printed with fmt.Sprint, or as JSON by the JSON formatter.
	Value any
}

// LoggerPrinter prints structured log lines: a toned level, the message
// and muted key=value pairs.
//
//	● info  Server started port=8080 env=production
//	▲ warn  Disk almost full free="2.1 GB"
//
// Log lines go to the error output by default, so they stay apart from the
// output of a program.
type LoggerPrinter struct {
	// Formatter picks colorful lines or JSON.
	Formatter LogFormatter
	// Writer receives the output. Nil means the error output.
	Writer io.Writer
	// Level is the least severe level that gets printed.
	Level LogLevel
	// ShowCaller adds the file and line that logged, as a "caller" argument.
	ShowCaller bool
	// CallerOffset skips stack frames when finding the caller, for loggers
	// wrapped in helpers of their own.
	CallerOffset int
	// ShowTime starts each line with the time.
	ShowTime bool
	// TimeFormat is the layout of the time, see time.Layout.
	TimeFormat string
	// Labels change how levels look. Levels without an entry use the
	// default look.
	Labels map[LogLevel]LevelLabel
	// MessageTone colors the message.
	MessageTone Tone
	// KeyTone colors the keys of arguments.
	KeyTone Tone
	// ValueTone colors the values of arguments.
	ValueTone Tone
	// TimeTone colors the time.
	TimeTone Tone
	// KeyStyles style the key and value of the arguments with these keys,
	// such as "error".
	KeyStyles map[string]Style
	// MaxWidth caps the width of a line. Longer lines put each argument on
	// a line of its own. Zero or less uses the terminal width.
	MaxWidth int
	// SortArguments sorts arguments by key. Otherwise they keep the order
	// they were given in.
	SortArguments bool
}

// Logger is the default logger. It prints Info and above.
var Logger = LoggerPrinter{
	Level:       LogLevelInfo,
	TimeFormat:  "15:04:05",
	MessageTone: ToneText,
	KeyTone:     ToneMuted,
	ValueTone:   ToneMuted,
	TimeTone:    ToneMuted,
	KeyStyles: map[string]Style{
		"error": {Fg: ToneError},
		"err":   {Fg: ToneError},
	},
}

// logNow returns the time stamped on log lines. Tests replace it.
var logNow = time.Now

// WithFormatter returns a copy of l that prints in format f.
func (l LoggerPrinter) WithFormatter(f LogFormatter) *LoggerPrinter {
	l.Formatter = f
	return &l
}

// WithWriter returns a copy of l that writes to w.
func (l LoggerPrinter) WithWriter(w io.Writer) *LoggerPrinter {
	l.Writer = w
	return &l
}

// WithLevel returns a copy of l that prints level and everything more
// severe.
func (l LoggerPrinter) WithLevel(level LogLevel) *LoggerPrinter {
	l.Level = level
	return &l
}

// WithCaller returns a copy of l that shows or hides the caller.
func (l LoggerPrinter) WithCaller(show ...bool) *LoggerPrinter {
	l.ShowCaller = flag(show)
	return &l
}

// WithCallerOffset returns a copy of l that skips offset more stack frames
// to find the caller.
func (l LoggerPrinter) WithCallerOffset(offset int) *LoggerPrinter {
	l.CallerOffset = offset
	return &l
}

// WithTime returns a copy of l that shows or hides the time.
func (l LoggerPrinter) WithTime(show ...bool) *LoggerPrinter {
	l.ShowTime = flag(show)
	return &l
}

// WithTimeFormat returns a copy of l that formats the time with layout.
func (l LoggerPrinter) WithTimeFormat(layout string) *LoggerPrinter {
	l.TimeFormat = layout
	return &l
}

// WithLabel returns a copy of l in which level looks like label.
func (l LoggerPrinter) WithLabel(level LogLevel, label LevelLabel) *LoggerPrinter {
	labels := maps.Clone(l.Labels)
	if labels == nil {
		labels = map[LogLevel]LevelLabel{}
	}

	labels[level] = label
	l.Labels = labels

	return &l
}

// WithMessageTone returns a copy of l whose messages use tone t.
func (l LoggerPrinter) WithMessageTone(t Tone) *LoggerPrinter {
	l.MessageTone = t
	return &l
}

// WithKeyTone returns a copy of l whose argument keys use tone t.
func (l LoggerPrinter) WithKeyTone(t Tone) *LoggerPrinter {
	l.KeyTone = t
	return &l
}

// WithValueTone returns a copy of l whose argument values use tone t.
func (l LoggerPrinter) WithValueTone(t Tone) *LoggerPrinter {
	l.ValueTone = t
	return &l
}

// WithTimeTone returns a copy of l whose time uses tone t.
func (l LoggerPrinter) WithTimeTone(t Tone) *LoggerPrinter {
	l.TimeTone = t
	return &l
}

// WithKeyStyles returns a copy of l that uses styles for arguments with
// these keys, replacing the styles it had.
func (l LoggerPrinter) WithKeyStyles(styles map[string]Style) *LoggerPrinter {
	l.KeyStyles = styles
	return &l
}

// AppendKeyStyles returns a copy of l that also uses styles for arguments
// with these keys.
func (l LoggerPrinter) AppendKeyStyles(styles map[string]Style) *LoggerPrinter {
	merged := maps.Clone(l.KeyStyles)
	if merged == nil {
		merged = map[string]Style{}
	}

	maps.Copy(merged, styles)
	l.KeyStyles = merged

	return &l
}

// AppendKeyStyle returns a copy of l that also uses style for arguments
// with key.
func (l LoggerPrinter) AppendKeyStyle(key string, style Style) *LoggerPrinter {
	return l.AppendKeyStyles(map[string]Style{key: style})
}

// WithMaxWidth returns a copy of l whose lines are at most w columns wide.
func (l LoggerPrinter) WithMaxWidth(w int) *LoggerPrinter {
	l.MaxWidth = w
	return &l
}

// WithSortArguments returns a copy of l that sorts arguments by key.
func (l LoggerPrinter) WithSortArguments(sort ...bool) *LoggerPrinter {
	l.SortArguments = flag(sort)
	return &l
}

// CanPrint reports whether l prints lines of level.
func (l LoggerPrinter) CanPrint(level LogLevel) bool {
	return l.Level != LogLevelDisabled && level != LogLevelDisabled && l.Level <= level
}

// Args pairs up keys and values: Args("user", "ada", "id", 42). With an odd
// number of arguments the last value gets the key ErrKeyWithoutValue.
func (l LoggerPrinter) Args(args ...any) []LoggerArgument {
	out := make([]LoggerArgument, 0, (len(args)+1)/2)

	for i := 0; i < len(args); i += 2 {
		if i+1 == len(args) {
			out = append(out, LoggerArgument{Key: ErrKeyWithoutValue, Value: args[i]})
			break
		}

		out = append(out, LoggerArgument{Key: fmt.Sprint(args[i]), Value: args[i+1]})
	}

	return out
}

// ArgsFromMap turns a map into arguments, sorted by key so lines are
// stable from run to run.
func (l LoggerPrinter) ArgsFromMap(m map[string]any) []LoggerArgument {
	out := make([]LoggerArgument, 0, len(m))
	for _, k := range slices.Sorted(maps.Keys(m)) {
		out = append(out, LoggerArgument{Key: k, Value: m[k]})
	}

	return out
}

// Trace logs at the trace level.
func (l LoggerPrinter) Trace(msg string, args ...[]LoggerArgument) {
	l.log(LogLevelTrace, msg, slices.Concat(args...), 0)
}

// Debug logs at the debug level.
func (l LoggerPrinter) Debug(msg string, args ...[]LoggerArgument) {
	l.log(LogLevelDebug, msg, slices.Concat(args...), 0)
}

// Info logs at the info level.
func (l LoggerPrinter) Info(msg string, args ...[]LoggerArgument) {
	l.log(LogLevelInfo, msg, slices.Concat(args...), 0)
}

// Warn logs at the warn level.
func (l LoggerPrinter) Warn(msg string, args ...[]LoggerArgument) {
	l.log(LogLevelWarn, msg, slices.Concat(args...), 0)
}

// Error logs at the error level.
func (l LoggerPrinter) Error(msg string, args ...[]LoggerArgument) {
	l.log(LogLevelError, msg, slices.Concat(args...), 0)
}

// Fatal logs at the fatal level and exits the program with code 1, also
// when the level of l hides the line.
func (l LoggerPrinter) Fatal(msg string, args ...[]LoggerArgument) {
	l.log(LogLevelFatal, msg, slices.Concat(args...), 0)
	screen.Exit(1)
}

// Print logs a line that is printed at every level but LogLevelDisabled.
func (l LoggerPrinter) Print(msg string, args ...[]LoggerArgument) {
	l.log(LogLevelPrint, msg, slices.Concat(args...), 0)
}

func (l LoggerPrinter) out() io.Writer { return orErrorOutput(l.Writer) }

// log prints one line. pc is the program counter of the caller, or 0 to
// look it up on the stack.
func (l LoggerPrinter) log(level LogLevel, msg string, args []LoggerArgument, pc uintptr) {
	if !l.CanPrint(level) || !OutputEnabled() {
		return
	}

	if l.ShowCaller {
		if pc == 0 {
			// Skip runtime.Callers, log and the level method.
			pcs := make([]uintptr, 1)
			if runtime.Callers(3+l.CallerOffset, pcs) > 0 {
				pc = pcs[0]
			}
		}

		if caller := callerOf(pc); caller != "" {
			args = append(slices.Clip(args), LoggerArgument{Key: "caller", Value: caller})
		}
	}

	if l.SortArguments {
		args = slices.Clone(args)
		slices.SortStableFunc(args, func(a, b LoggerArgument) int { return strings.Compare(a.Key, b.Key) })
	}

	w := l.out()

	var line string

	switch l.Formatter {
	case LogFormatterJSON:
		line = l.renderJSON(level, msg, args)
	default:
		line = l.render(canvasFor(w), level, msg, args)
	}

	emit(w, line+"\n")
}

// callerOf returns "dir/file.go:line" for pc, or "" when it is unknown.
func callerOf(pc uintptr) string {
	if pc == 0 {
		return ""
	}

	frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()
	if frame.File == "" {
		return ""
	}

	file := filepath.ToSlash(frame.File)
	if i := strings.LastIndexByte(file, '/'); i > 0 {
		if j := strings.LastIndexByte(file[:i], '/'); j >= 0 {
			file = file[j+1:]
		}
	}

	return file + ":" + strconv.Itoa(frame.Line)
}

// label returns how level looks, with the glyph filled in.
func (l LoggerPrinter) label(c canvas, level LogLevel) LevelLabel {
	if lb, ok := l.Labels[level]; ok {
		if lb.Glyph == "" {
			lb.Glyph = defaultLabel(c.g, level).Glyph
		}

		return lb
	}

	return defaultLabel(c.g, level)
}

func defaultLabel(g Glyphs, level LogLevel) LevelLabel {
	switch level {
	case LogLevelTrace:
		return LevelLabel{Glyph: g.Debug, Word: "trace", Tone: ToneMuted}
	case LogLevelDebug:
		return LevelLabel{Glyph: g.Debug, Word: "debug", Tone: ToneMuted}
	case LogLevelWarn:
		return LevelLabel{Glyph: g.Warning, Word: "warn", Tone: ToneWarning}
	case LogLevelError:
		return LevelLabel{Glyph: g.Error, Word: "error", Tone: ToneError}
	case LogLevelFatal:
		return LevelLabel{Glyph: g.Error, Word: "fatal", Tone: ToneError}
	case LogLevelPrint:
		return LevelLabel{Glyph: g.Bullet, Word: "", Tone: ToneMuted}
	default:
		return LevelLabel{Glyph: g.Info, Word: "info", Tone: ToneInfo}
	}
}

// logWordWidth is the width level words are padded to, so messages line
// up.
const logWordWidth = 5

// render formats a line for humans.
func (l LoggerPrinter) render(c canvas, level LogLevel, msg string, args []LoggerArgument) string {
	lb := l.label(c, level)
	st := Style{Fg: lb.Tone}

	if level == LogLevelFatal {
		st.Attrs = color.Bold
	}

	prefix := c.paint(lb.Tone, lb.Glyph) + " " + c.style(st, text.PadRight(lb.Word, logWordWidth)) + " "
	if l.ShowTime {
		prefix = c.paint(l.TimeTone, logNow().Format(l.TimeFormat)) + " " + prefix
	}

	msg = strings.TrimRight(msg, "\n")
	width := l.lineWidth(c)

	pairs := make([]string, len(args))
	multiline := strings.Contains(msg, "\n")

	for i, a := range args {
		pairs[i] = l.pair(c, a, true)
		if strings.Contains(fmt.Sprint(a.Value), "\n") {
			multiline = true
		}
	}

	inline := prefix + c.paint(l.MessageTone, msg)
	if len(pairs) > 0 {
		inline += " " + strings.Join(pairs, " ")
	}

	// Without styling or with an unknown width a line stays one line, so
	// files and pipes can be grepped.
	if !StylingEnabled() || width <= 0 || (!multiline && text.Width(inline) <= width) {
		return node(inline)
	}

	return node(l.renderBlock(c, prefix, msg, args, width))
}

// renderBlock formats a line that does not fit: the message wraps under
// itself and every argument gets a line of its own.
func (l LoggerPrinter) renderBlock(c canvas, prefix, msg string, args []LoggerArgument, width int) string {
	indentWidth := text.Width(prefix)
	indent := strings.Repeat(" ", indentWidth)
	room := max(width-indentWidth, 20)

	var b strings.Builder

	b.WriteString(prefix)

	for i, line := range text.Wrap(msg, room) {
		if i > 0 {
			b.WriteString("\n" + indent)
		}

		b.WriteString(c.paint(l.MessageTone, line))
	}

	for _, a := range args {
		value := fmt.Sprint(a.Value)
		if !strings.Contains(value, "\n") {
			b.WriteString("\n" + indent + l.pair(c, a, false))
			continue
		}

		key, valueStyle := l.pairStyles(a.Key)
		b.WriteString("\n" + indent + c.style(key, a.Key) + c.paint(l.KeyTone, "="))

		for _, line := range text.Wrap(strings.TrimRight(value, "\n"), max(room-2, 10)) {
			b.WriteString("\n" + indent + "  " + c.style(valueStyle, line))
		}
	}

	return b.String()
}

// pair formats key=value. Inline, values that contain spaces, quotes or
// line breaks are quoted so the line stays unambiguous.
func (l LoggerPrinter) pair(c canvas, a LoggerArgument, inline bool) string {
	key, value := l.pairStyles(a.Key)

	v := fmt.Sprint(a.Value)
	if needsQuotes(v, inline) {
		v = strconv.Quote(v)
	}

	return c.style(key, a.Key) + c.paint(l.KeyTone, "=") + c.style(value, v)
}

// pairStyles returns the styles of the key and the value of an argument.
func (l LoggerPrinter) pairStyles(key string) (keyStyle, valueStyle Style) {
	if st, ok := l.KeyStyles[key]; ok {
		return st, st
	}

	return Style{Fg: l.KeyTone}, Style{Fg: l.ValueTone}
}

func needsQuotes(v string, inline bool) bool {
	if v == "" {
		return true
	}

	if inline && strings.ContainsAny(v, "\n\r\t") {
		return true
	}

	return strings.ContainsAny(v, " \"=")
}

// lineWidth returns the width a line may take, or 0 when it is unknown.
func (l LoggerPrinter) lineWidth(c canvas) int {
	switch {
	case c.width <= 0:
		return 0
	case l.MaxWidth > 0 && l.MaxWidth < c.width:
		return l.MaxWidth
	default:
		return c.width
	}
}

// renderJSON formats a line as a JSON object.
func (l LoggerPrinter) renderJSON(level LogLevel, msg string, args []LoggerArgument) string {
	m := make(map[string]any, len(args)+3)

	for _, a := range args {
		m[a.Key] = jsonValue(a.Value)
	}

	m["level"] = level.String()
	m["timestamp"] = logNow().Format(l.TimeFormat)
	m["msg"] = msg

	b, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprint(m)
	}

	return string(b)
}

// jsonValue returns v in a form json.Marshal prints usefully.
func jsonValue(v any) any {
	switch x := v.(type) {
	case slog.Value:
		return jsonValue(x.Resolve().Any())
	case error:
		return x.Error()
	case json.Marshaler:
		return x
	case fmt.Stringer:
		return x.String()
	default:
		return v
	}
}
