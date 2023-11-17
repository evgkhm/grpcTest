package logging

import (
	"log/slog"
	"os"
)

const (
	defaultLevel     = slog.LevelInfo
	defaultIsJSON    = false
	defaultAddSource = false
)

type LoggerOption func(*LoggerOptions)

func Logger(opts ...LoggerOption) *slog.Logger {
	config := &LoggerOptions{
		Level:     defaultLevel,
		IsJSON:    defaultIsJSON,
		AddSource: defaultAddSource,
	}

	for _, opt := range opts {
		opt(config)
	}

	if err := config.Validate(); err != nil {
		return slog.Default()
	}

	options := &slog.HandlerOptions{
		AddSource: config.AddSource,
		Level:     config.Level,
	}

	var h slog.Handler = slog.NewTextHandler(os.Stdout, options)

	if config.IsJSON {
		h = slog.NewJSONHandler(os.Stdout, options)
	}

	return slog.New(h)
}

type LoggerOptions struct {
	Level     slog.Level
	AddSource bool
	IsJSON    bool
}

func (o *LoggerOptions) Validate() error {
	return nil
}
