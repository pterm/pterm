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

// Enabled returns true if the given level is enabled.
func (s *SlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	switch level {
	case slog.LevelDebug:
		return s.logger.CanPrint(LogLevelDebug)
	case slog.LevelInfo:
		return s.logger.CanPrint(LogLevelInfo)
	case slog.LevelWarn:
		return s.logger.CanPrint(LogLevelWarn)
	case slog.LevelError:
		return s.logger.CanPrint(LogLevelError)
	}

	return false
}

// Handle handles the given record.
func (s *SlogHandler) Handle(_ context.Context, record slog.Record) error {
	level := record.Level
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
	case slog.LevelDebug:
		logger.Debug(message, argsWrapped...)
	case slog.LevelInfo:
		logger.Info(message, argsWrapped...)
	case slog.LevelWarn:
		logger.Warn(message, argsWrapped...)
	case slog.LevelError:
		logger.Error(message, argsWrapped...)
	default:
		logger.Print(message, argsWrapped...)
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

// NewSlogHandler returns a new logging handler that can be integrated with log/slog.
func NewSlogHandler(logger *Logger) *SlogHandler {
	return &SlogHandler{logger: logger}
}
