package pterm

import (
	"context"

	"log/slog"
)

// SlogHandler is a slog.Handler that prints log records via a pterm Logger.
type SlogHandler struct {
	logger *Logger
	attrs  []slog.Attr
}

// slogLevel maps each band to its lower severity. Values outside slog's
// standard range use Trace below Debug and Error above Error.
func slogLevel(level slog.Level) LogLevel {
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

// Enabled returns true if the given level is enabled.
func (s *SlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return s.logger.CanPrint(slogLevel(level))
}

// Handle handles the given record.
func (s *SlogHandler) Handle(_ context.Context, record slog.Record) error {
	level := slogLevel(record.Level)
	message := record.Message

	// Collect the attrs in the order they were supplied, so the logger can print
	// them predictably instead of in Go's randomized map order. Attrs bound via
	// WithAttrs come first, followed by the record's own attrs, matching the
	// behavior of slog's own text handler. Alphabetical ordering is available
	// through the logger's SortArguments option.
	args := make([]LoggerArgument, 0, len(s.attrs)+record.NumAttrs())

	for _, attr := range s.attrs {
		args = append(args, LoggerArgument{Key: attr.Key, Value: attr.Value})
	}

	record.Attrs(func(attr slog.Attr) bool {
		args = append(args, LoggerArgument{Key: attr.Key, Value: attr.Value})
		return true
	})

	// Wrapping args inside another slice to match [][]LoggerArgument
	argsWrapped := [][]LoggerArgument{args}

	logger := s.logger

	// Must be done here, see https://github.com/pterm/pterm/issues/608#issuecomment-1876001650
	if logger.CallerOffset == 0 {
		logger = logger.WithCallerOffset(3)
	}

	switch level {
	case LogLevelTrace:
		logger.Trace(message, argsWrapped...)
	case LogLevelDebug:
		logger.Debug(message, argsWrapped...)
	case LogLevelInfo:
		logger.Info(message, argsWrapped...)
	case LogLevelWarn:
		logger.Warn(message, argsWrapped...)
	case LogLevelError:
		logger.Error(message, argsWrapped...)
	}

	return nil
}

// WithAttrs returns a new handler with the given attributes.
func (s *SlogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newS := *s
	newS.attrs = attrs

	return &newS
}

// WithGroup is not yet supported.
func (s *SlogHandler) WithGroup(_ string) slog.Handler {
	// Grouping is not yet supported by pterm.
	return s
}

// NewSlogHandler returns a handler for log/slog. Custom levels between Debug,
// Info, Warn, and Error use the lower band. Levels below Debug use Trace;
// levels above Error use Error. For example:
//
//	logger := slog.New(pterm.NewSlogHandler(pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)))
//	logger.Log(context.Background(), slog.Level(-5), "tracing a request")
func NewSlogHandler(logger *Logger) *SlogHandler {
	return &SlogHandler{logger: logger}
}
