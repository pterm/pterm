package pterm

import (
	"context"

	"log/slog"
)

type SlogHandler struct {
	logger *Logger
	attrs  []slog.Attr
}

// Enabled returns true if the given level is enabled.
func (s *SlogHandler) Enabled(ctx context.Context, level slog.Level) bool {
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
func (s *SlogHandler) Handle(ctx context.Context, record slog.Record) error {
	level := record.Level
	message := record.Message

	// Convert slog Attrs to a map.
	keyValsMap := make(map[string]interface{})

	record.Attrs(func(attr slog.Attr) bool {
		keyValsMap[attr.Key] = attr.Value
		return true
	})

	for _, attr := range s.attrs {
		keyValsMap[attr.Key] = attr.Value
	}

	args := s.logger.ArgsFromMap(keyValsMap)

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
func (s *SlogHandler) WithGroup(name string) slog.Handler {
	// Grouping is not yet supported by pterm.
	return s
}

// NewSlogHandler returns a new logging handler that can be intrgrated with log/slog.
func NewSlogHandler(logger *Logger) *SlogHandler {
	return &SlogHandler{logger: logger}
}
