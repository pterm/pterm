package pterm_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// newBufferedLogger returns a logger writing into the returned buffer, with
// timestamps disabled so output is deterministic.
func newBufferedLogger() (*pterm.Logger, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	logger := pterm.DefaultLogger.WithWriter(buf).WithTime(false)

	return logger, buf
}

// Level filtering: messages below the configured level must produce no
// output at all, messages at or above it must be printed.

func TestLoggerLevelFiltering(t *testing.T) {
	methods := []struct {
		name  string
		level pterm.LogLevel
		call  func(l *pterm.Logger, msg string)
	}{
		{"Trace", pterm.LogLevelTrace, func(l *pterm.Logger, msg string) { l.Trace(msg) }},
		{"Debug", pterm.LogLevelDebug, func(l *pterm.Logger, msg string) { l.Debug(msg) }},
		{"Info", pterm.LogLevelInfo, func(l *pterm.Logger, msg string) { l.Info(msg) }},
		{"Warn", pterm.LogLevelWarn, func(l *pterm.Logger, msg string) { l.Warn(msg) }},
		{"Error", pterm.LogLevelError, func(l *pterm.Logger, msg string) { l.Error(msg) }},
		{"Print", pterm.LogLevelPrint, func(l *pterm.Logger, msg string) { l.Print(msg) }},
	}

	configuredLevels := []pterm.LogLevel{
		pterm.LogLevelDisabled,
		pterm.LogLevelTrace,
		pterm.LogLevelDebug,
		pterm.LogLevelInfo,
		pterm.LogLevelWarn,
		pterm.LogLevelError,
		pterm.LogLevelFatal,
		pterm.LogLevelPrint,
	}

	for _, configured := range configuredLevels {
		for _, method := range methods {
			t.Run(configured.String()+"/"+method.name, func(t *testing.T) {
				logger, buf := newBufferedLogger()
				method.call(logger.WithLevel(configured), "the message")

				shouldPrint := configured != pterm.LogLevelDisabled && configured <= method.level
				if shouldPrint {
					out := stripANSI(buf.String())
					assert.Contains(t, out, "the message")
					assert.Contains(t, out, method.level.String(), "the printed line must carry the level name")
				} else {
					assert.Empty(t, buf.String(), "%s must print nothing at level %s", method.name, configured)
				}
			})
		}
	}
}

func TestLoggerCanPrint(t *testing.T) {
	assert.True(t, pterm.Logger{Level: pterm.LogLevelInfo}.CanPrint(pterm.LogLevelInfo))
	assert.True(t, pterm.Logger{Level: pterm.LogLevelInfo}.CanPrint(pterm.LogLevelError))
	assert.False(t, pterm.Logger{Level: pterm.LogLevelInfo}.CanPrint(pterm.LogLevelDebug))
	assert.False(t, pterm.Logger{Level: pterm.LogLevelDisabled}.CanPrint(pterm.LogLevelFatal), "a disabled logger never prints")
}

func TestLogLevelString(t *testing.T) {
	expected := map[pterm.LogLevel]string{
		pterm.LogLevelDisabled: "",
		pterm.LogLevelTrace:    "TRACE",
		pterm.LogLevelDebug:    "DEBUG",
		pterm.LogLevelInfo:     "INFO",
		pterm.LogLevelWarn:     "WARN",
		pterm.LogLevelError:    "ERROR",
		pterm.LogLevelFatal:    "FATAL",
		pterm.LogLevelPrint:    "PRINT",
	}

	for level, name := range expected {
		assert.Equal(t, name, level.String())
	}

	assert.Equal(t, "Unknown", pterm.LogLevel(255).String())
}

// Args conversion.

func TestLoggerArgs(t *testing.T) {
	logger := pterm.DefaultLogger

	t.Run("pairs", func(t *testing.T) {
		args := logger.Args("k1", "v1", "k2", 2)
		assert.Equal(t, []pterm.LoggerArgument{
			{Key: "k1", Value: "v1"},
			{Key: "k2", Value: 2},
		}, args)
	})

	t.Run("no args", func(t *testing.T) {
		assert.Empty(t, logger.Args())
	})

	t.Run("odd number of args marks the dangling value", func(t *testing.T) {
		assert.Equal(t, []pterm.LoggerArgument{
			{Key: pterm.ErrKeyWithoutValue, Value: "lonely"},
		}, logger.Args("lonely"))

		assert.Equal(t, []pterm.LoggerArgument{
			{Key: "k", Value: "v"},
			{Key: pterm.ErrKeyWithoutValue, Value: "extra"},
		}, logger.Args("k", "v", "extra"))
	})
}

func TestLoggerArgsFromMap(t *testing.T) {
	args := pterm.DefaultLogger.ArgsFromMap(map[string]any{"a": 1, "b": "x"})

	asMap := make(map[string]any, len(args))
	for _, arg := range args {
		asMap[arg.Key] = arg.Value
	}

	assert.Equal(t, map[string]any{"a": 1, "b": "x"}, asMap)
}

// Colorful (text) formatter.

func TestLoggerColorfulInlineFormat(t *testing.T) {
	logger, buf := newBufferedLogger()

	logger.Info("service started", logger.Args("port", 8080, "env", "prod"))

	// The level is padded to a fixed width so messages align across levels.
	assert.Equal(t, "INFO  service started port=8080 env=prod\n", stripANSI(buf.String()))

	buf.Reset()
	logger.Error("boom")
	assert.Equal(t, "ERROR boom\n", stripANSI(buf.String()))
}

func TestLoggerQuotesValuesWithSpaces(t *testing.T) {
	logger, buf := newBufferedLogger()

	logger.Info("msg", logger.Args("key", "two words", "plain", "word"))

	assert.Equal(t, `INFO  msg key="two words" plain=word`+"\n", stripANSI(buf.String()))
}

func TestLoggerWithTimeAndFormat(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := pterm.DefaultLogger.WithWriter(buf).WithTime(true).WithTimeFormat("2006")

	logger.Info("timed")

	assert.Equal(t, time.Now().Format("2006")+" INFO  timed\n", stripANSI(buf.String()))

	buf.Reset()
	logger.WithTime(false).Info("untimed")
	assert.Equal(t, "INFO  untimed\n", stripANSI(buf.String()))
}

func TestLoggerWithCallerAddsCallerArg(t *testing.T) {
	logger, buf := newBufferedLogger()

	logger.WithCaller().Info("who called")

	out := stripANSI(buf.String())
	assert.Contains(t, out, "caller=")
	assert.Contains(t, out, "logger_test.go:", "the caller must point at this test file")
}

// A multiline argument value must force block rendering with the tree rail,
// and continuation lines must align under the start of the value.
func TestLoggerMultilineValueRendersAsBlock(t *testing.T) {
	logger, buf := newBufferedLogger()

	logger.Info("msg", logger.Args("stack", "a\nb"))

	lines := strings.Split(strings.TrimSuffix(stripANSI(buf.String()), "\n"), "\n")
	require.Len(t, lines, 3)
	assert.Equal(t, "INFO  msg", lines[0])
	assert.Equal(t, "      └ stack=a", lines[1])
	assert.Equal(t, strings.Repeat(" ", 14)+"b", lines[2], "the continuation line must align under the value start")
}

// WithMaxWidth: overflowing content wraps into a block where every line stays
// within the width and arguments hang off a tree rail. (The exact wrapped
// layout is locked by snapshot tests; this verifies the structure.)
func TestLoggerMaxWidthWrapsOverflowingLines(t *testing.T) {
	const width = 30

	logger, buf := newBufferedLogger()
	logger = logger.WithMaxWidth(width)

	logger.Info("this message is far too long to fit into thirty columns",
		logger.Args("first", "value", "second", "value"))

	out := strings.TrimSuffix(stripANSI(buf.String()), "\n")
	lines := strings.Split(out, "\n")
	require.Greater(t, len(lines), 2, "overflowing content must wrap into multiple lines")

	for _, line := range lines {
		assert.LessOrEqual(t, utf8.RuneCountInString(line), width, "line %q exceeds the max width", line)
	}

	assert.Contains(t, out, "├ first=", "all but the last argument hang off a branch connector")
	assert.Contains(t, out, "└ second=", "the last argument hangs off the end connector")
}

func TestLoggerShortLineStaysInline(t *testing.T) {
	logger, buf := newBufferedLogger()
	logger = logger.WithMaxWidth(80)

	logger.Info("short", logger.Args("k", "v"))

	assert.Equal(t, "INFO  short k=v\n", stripANSI(buf.String()))
}

// Raw output mode keeps every log on a single line so it stays grep-friendly.
func TestLoggerRawOutputStaysOnSingleLine(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	logger, buf := newBufferedLogger()
	logger = logger.WithMaxWidth(20)

	logger.Info("this would normally wrap into a block", logger.Args("key", "value with spaces"))

	out := buf.String()
	assert.NotContains(t, out, "\x1b", "raw output must not contain escape codes")
	assert.Equal(t, 1, strings.Count(out, "\n"), "raw output must stay on one line")
	assert.Contains(t, out, `key="value with spaces"`)
}

func TestLoggerAppendKeyStyle(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := pterm.Logger{
		Formatter: pterm.LogFormatterColorful,
		Writer:    buf,
		Level:     pterm.LogLevelInfo,
		KeyStyles: map[string]pterm.Style{},
	}

	logger.AppendKeyStyle("special", *pterm.NewStyle(pterm.FgMagenta)).Info("msg", logger.Args("special", "v"))

	assert.Contains(t, buf.String(), "\x1b[35mspecial\x1b[0m", "the key must be rendered with its configured style")
}

// JSON formatter.

func TestLoggerJSONFormat(t *testing.T) {
	logger, buf := newBufferedLogger()
	logger = logger.WithFormatter(pterm.LogFormatterJSON)

	logger.Warn("something odd", logger.Args("code", 42, "detail", "disk full"))

	var m map[string]any

	require.NoError(t, json.Unmarshal(buf.Bytes(), &m), "JSON output must parse")
	assert.Equal(t, "WARN", m["level"])
	assert.Equal(t, "something odd", m["msg"])
	assert.Equal(t, float64(42), m["code"])
	assert.Equal(t, "disk full", m["detail"])

	timestamp, ok := m["timestamp"].(string)
	require.True(t, ok, "timestamp must be present")

	_, err := time.Parse(logger.TimeFormat, timestamp)
	assert.NoError(t, err, "timestamp must follow the configured TimeFormat")
}

func TestLoggerJSONCaller(t *testing.T) {
	logger, buf := newBufferedLogger()

	logger.WithFormatter(pterm.LogFormatterJSON).WithCaller().Error("with caller")

	var m map[string]any

	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))

	caller, ok := m["caller"].(string)
	require.True(t, ok, "caller key must be present")
	assert.Contains(t, caller, "logger_test.go:")
}

// Fatal handling. Fatal calls os.Exit(1) after printing, so the exiting path
// runs in a subprocess; the non-exiting path (disabled logger) runs inline.

func TestLoggerFatalOnDisabledLoggerDoesNotExitOrPrint(t *testing.T) {
	logger, buf := newBufferedLogger()

	// If this exited, the whole test binary would die here.
	logger.WithLevel(pterm.LogLevelDisabled).Fatal("must not be printed")

	assert.Empty(t, buf.String())
}

func TestLoggerFatalExitsWithCode1(t *testing.T) {
	if os.Getenv("PTERM_TEST_LOGGER_FATAL") == "1" {
		pterm.DefaultLogger.WithTime(false).WithWriter(os.Stderr).Fatal("fatal crash")
		return // unreachable: Fatal exits the process
	}

	//nolint:gosec // re-runs this test binary to observe the os.Exit.
	cmd := exec.Command(os.Args[0], "-test.run", "^TestLoggerFatalExitsWithCode1$")

	cmd.Env = append(os.Environ(), "PTERM_TEST_LOGGER_FATAL=1")

	out, err := cmd.CombinedOutput()

	var exitErr *exec.ExitError

	require.ErrorAs(t, err, &exitErr, "the subprocess must exit non-zero")
	assert.Equal(t, 1, exitErr.ExitCode())
	assert.Contains(t, string(out), "FATAL")
	assert.Contains(t, string(out), "fatal crash")
}

// slog integration: pterm as a slog backend.

func TestSlogHandlerEnabledHonorsLoggerLevel(t *testing.T) {
	logger, _ := newBufferedLogger()
	handler := pterm.NewSlogHandler(logger.WithLevel(pterm.LogLevelWarn))

	assert.False(t, handler.Enabled(t.Context(), slog.LevelDebug))
	assert.False(t, handler.Enabled(t.Context(), slog.LevelInfo))
	assert.True(t, handler.Enabled(t.Context(), slog.LevelWarn))
	assert.True(t, handler.Enabled(t.Context(), slog.LevelError))
}

func TestSlogHandlerRoutesRecordsThroughTheLogger(t *testing.T) {
	logger, buf := newBufferedLogger()
	slogger := slog.New(pterm.NewSlogHandler(logger.WithFormatter(pterm.LogFormatterJSON).WithLevel(pterm.LogLevelDebug)))

	levels := map[string]func(msg string, args ...any){
		"DEBUG": slogger.Debug,
		"INFO":  slogger.Info,
		"WARN":  slogger.Warn,
		"ERROR": slogger.Error,
	}

	for levelName, log := range levels {
		t.Run(levelName, func(t *testing.T) {
			buf.Reset()
			log("slog message", "key", "value", "answer", 42)

			var m map[string]any

			require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
			assert.Equal(t, levelName, m["level"])
			assert.Equal(t, "slog message", m["msg"])
			assert.Equal(t, "value", m["key"], "slog attrs must flow through as logger args")
			assert.Equal(t, float64(42), m["answer"])
		})
	}
}

func TestSlogHandlerFiltersBelowLoggerLevel(t *testing.T) {
	logger, buf := newBufferedLogger()
	slogger := slog.New(pterm.NewSlogHandler(logger.WithLevel(pterm.LogLevelInfo)))

	slogger.Debug("hidden")
	assert.Empty(t, buf.String(), "records below the logger level must be dropped")

	slogger.Info("visible")
	assert.Contains(t, stripANSI(buf.String()), "visible")
}

func TestSlogHandlerWithAttrs(t *testing.T) {
	logger, buf := newBufferedLogger()
	slogger := slog.New(pterm.NewSlogHandler(logger.WithFormatter(pterm.LogFormatterJSON)))

	slogger.With("request_id", "abc123").Info("handled")

	var m map[string]any

	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, "abc123", m["request_id"], "attrs bound via With must be attached to every record")
}

func TestSlogHandlerWithGroupFlattensAttrs(t *testing.T) {
	logger, buf := newBufferedLogger()
	slogger := slog.New(pterm.NewSlogHandler(logger.WithFormatter(pterm.LogFormatterJSON)))

	// Groups are not supported by pterm (yet); attrs inside a group must
	// still be printed, just without the group prefix.
	slogger.WithGroup("http").Info("request", "method", "GET")

	var m map[string]any

	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, "GET", m["method"])
}
