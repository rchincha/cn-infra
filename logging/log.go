// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

// LogLevel represents severity of log record
type LogLevel int

const (
	// DebugLevel - the most verbose logging
	DebugLevel LogLevel = iota
	// InfoLevel level - general operational entries about what's going on inside the application.
	InfoLevel
	// WarnLevel - non-critical entries that deserve eyes.
	WarnLevel
	// ErrorLevel level - used for errors that should definitely be noted.
	ErrorLevel
	// FatalLevel - logs and then calls `os.Exit(1)`.
	FatalLevel
	// PanicLevel - highest level of severity. Logs and then calls panic with the message passed in.
	PanicLevel
)

// Logger provides logging capabilities
type Logger interface {
	LogWithLevel
	// SetLevel modifies the LogLevel
	SetLevel(level LogLevel)
	// GetLevel returns currently set logLevel
	GetLevel() LogLevel
	// WithField creates one structured field
	WithField(key string, value interface{}) LogWithLevel
	// WithFields creates multiple structured fields
	WithFields(fields map[string]interface{}) LogWithLevel
}

// LogWithLevel allows to log with different log levels
type LogWithLevel interface {
	// Debug logs using Debug level
	Debug(args ...interface{})
	// Info logs using Info level
	Info(args ...interface{})
	// Warning logs using Warning level
	Warn(args ...interface{})
	// Error logs using Error level
	Error(args ...interface{})
	// Panic logs using Panic level and panics
	Panic(args ...interface{})
	// Fatal logs using Fatal level and calls os.Exit(1)
	Fatal(args ...interface{})
}

// Registry groups multiple Logger instances and allows to mange their log levels.
type Registry interface {
	// List Loggers returns a map (loggerName => log level)
	ListLoggers() map[string]string
	// SetLevel modifies log level of selected logger in the registry
	SetLevel(logger, level string) error
	// GetLevel returns the currently set log level of the logger from registry
	GetLevel(logger string) (string, error)
	// GetLoggerByName returns a logger instance identified by name from registry
	GetLoggerByName(name string) (*Logger, bool)
	// ClearRegistry removes all loggers except the default one from registry
	ClearRegistry()
}

// String converts the Level to a string. E.g. PanicLevel becomes "panic".
func (level LogLevel) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}

	return "unknown"
}
