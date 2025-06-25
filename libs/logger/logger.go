package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger provides structured logging with trace ID support
type Logger struct {
	logger  zerolog.Logger
	traceID string
}

// GetTraceID returns the current trace ID
func (l *Logger) GetTraceID() string {
	return l.traceID
}

// Config holds basic logger configuration
type Config struct {
	Level  string `json:"level"`  // debug, info, warn, error
	Format string `json:"format"` // json, console
}

// DefaultConfig returns default configuration
func DefaultConfig() Config {
	return Config{
		Level:  "info",
		Format: "json",
	}
}

var globalLogger *Logger

// Init initializes the global logger
func Init(config Config) *Logger {
	level, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	var logger zerolog.Logger
	if config.Format == "console" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	} else {
		logger = log.Logger
	}

	globalLogger = &Logger{
		logger: logger,
	}
	return globalLogger
}

// GetLogger returns the global logger
func GetLogger() *Logger {
	if globalLogger == nil {
		globalLogger = Init(DefaultConfig())
	}
	return globalLogger
}

// WithTraceID creates a new logger with trace ID
func (l *Logger) WithTraceID(traceID string) *Logger {
	return &Logger{
		logger:  l.logger.With().Str("trace_id", traceID).Logger(),
		traceID: traceID,
	}
}

// NewTraceID generates a new trace ID and returns logger with it
func (l *Logger) NewTraceID() *Logger {
	traceID := generateTraceID()
	return l.WithTraceID(traceID)
}

// Debug logs debug message with fixed JSON structure
func (l *Logger) Debug(component, action, message string, data interface{}) {
	l.logger.Debug().
		Str("level", "debug").
		Str("component", component).
		Str("action", action).
		Str("message", message).
		Interface("data", data).
		Time("timestamp", time.Now()).
		Msg("")
}

// Info logs info message with fixed JSON structure
func (l *Logger) Info(component, action, message string, data interface{}) {
	l.logger.Info().
		Str("level", "info").
		Str("component", component).
		Str("action", action).
		Str("message", message).
		Interface("data", data).
		Time("timestamp", time.Now()).
		Msg("")
}

// Warn logs warning message with fixed JSON structure
func (l *Logger) Warn(component, action, message string, data interface{}) {
	l.logger.Warn().
		Str("level", "warn").
		Str("component", component).
		Str("action", action).
		Str("message", message).
		Interface("data", data).
		Time("timestamp", time.Now()).
		Msg("")
}

// Error logs error message with fixed JSON structure
func (l *Logger) Error(component, action, message string, err error, data interface{}) {
	l.logger.Error().
		Str("level", "error").
		Str("component", component).
		Str("action", action).
		Str("message", message).
		Err(err).
		Interface("data", data).
		Time("timestamp", time.Now()).
		Msg("")
}

// Package level functions using global logger

// WithTraceID creates a new logger with trace ID using global logger
func WithTraceID(traceID string) *Logger {
	return GetLogger().WithTraceID(traceID)
}

// NewTraceID generates a new trace ID using global logger
func NewTraceID() *Logger {
	return GetLogger().NewTraceID()
}

// Debug logs debug message using global logger
func Debug(component, action, message string, data interface{}) {
	GetLogger().Debug(component, action, message, data)
}

// Info logs info message using global logger
func Info(component, action, message string, data interface{}) {
	GetLogger().Info(component, action, message, data)
}

// Warn logs warning message using global logger
func Warn(component, action, message string, data interface{}) {
	GetLogger().Warn(component, action, message, data)
}

// Error logs error message using global logger
func Error(component, action, message string, err error, data interface{}) {
	GetLogger().Error(component, action, message, err, data)
}

// generateTraceID generates a simple trace ID
func generateTraceID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString generates a random string of given length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
