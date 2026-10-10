package pterm

import (
	"context"
	"log/slog"
	"slices"
)

// SlogHandler is a slog.Handler that prints records with a LoggerPrinter,
// so code that logs with log/slog gets pterm's look:
//
//	slog.SetDefault(slog.New(pterm.NewSlogHandler(pterm.Logger.WithLevel(pterm.LogLevelDebug))))
type SlogHandler struct {
	logger *LoggerPrinter
	attrs  []LoggerArgument
	group  string
}

// NewSlogHandler returns a handler that prints with logger, or with Logger
// when logger is nil. The handler reads logger for every record, so later
// changes to it take effect.
func NewSlogHandler(logger *LoggerPrinter) *SlogHandler {
	if logger == nil {
		logger = &Logger
	}

	return &SlogHandler{logger: logger}
}

// logLevelOf maps a slog level to the pterm level that covers it.
func logLevelOf(level slog.Level) LogLevel {
	switch {
	case level < slog.LevelDebug:
		return LogLevelTrace
	case level < slog.LevelInfo:
		return LogLevelDebug
	case level < slog.LevelWarn:
		return LogLevelInfo
	case level < slog.LevelError:
		return LogLevelWarn
	default:
		return LogLevelError
	}
}

// Enabled reports whether the logger prints records of level.
func (s *SlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return s.logger.CanPrint(logLevelOf(level))
}

// Handle prints record. Attributes keep their order: those added with
// WithAttrs first, then the record's own. Groups become key prefixes, as in
// "request.id".
func (s *SlogHandler) Handle(_ context.Context, record slog.Record) error {
	args := slices.Clone(s.attrs)

	record.Attrs(func(a slog.Attr) bool {
		args = appendAttr(args, s.group, a)
		return true
	})

	s.logger.log(logLevelOf(record.Level), record.Message, args, record.PC)

	return nil
}

// WithAttrs returns a handler that adds attrs to every record.
func (s *SlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	next := *s
	next.attrs = slices.Clone(s.attrs)

	for _, a := range attrs {
		next.attrs = appendAttr(next.attrs, s.group, a)
	}

	return &next
}

// WithGroup returns a handler that prefixes the keys of later attributes
// with name.
func (s *SlogHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return s
	}

	next := *s
	next.group = s.group + name + "."

	return &next
}

// appendAttr appends a as arguments, flattening groups into dotted keys.
func appendAttr(args []LoggerArgument, prefix string, a slog.Attr) []LoggerArgument {
	a.Value = a.Value.Resolve()

	if a.Equal(slog.Attr{}) {
		return args
	}

	if a.Value.Kind() != slog.KindGroup {
		return append(args, LoggerArgument{Key: prefix + a.Key, Value: a.Value})
	}

	// A group without a key adds its attributes at the current level, as
	// slog's own handlers do.
	if a.Key != "" {
		prefix += a.Key + "."
	}

	for _, ga := range a.Value.Group() {
		args = appendAttr(args, prefix, ga)
	}

	return args
}
