package litelog

import (
	"io"
	"os"
	"time"
)

// LOG levels
const (
	Info = iota
	Warn
	Err
	Debug
	Trace
)

const (
	infoPrefix  = "[INFO]"
	warnPrefix  = "[WARN]"
	errPrefix   = "[ERROR]"
	debugPrefix = "[DEBUG]"
	tracePrefix = "[TRACE]"
)

// Logger represents logger object
type Logger struct {
	hasTime    bool
	hasPrefix  bool
	level      uint8
	prefix     string
	timeLayout string
	writer     io.Writer
}

// Setter sets Logger parameters
type Setter func(l *Logger)

// WithWriter sets io.Writer fo Logger type (default os.Stdout)
func WithWriter(wr io.Writer) Setter {
	return func(l *Logger) {
		l.writer = wr
	}
}

// WithPrefix sets prefix in log messages
func WithPrefix(pr string) Setter {
	return func(l *Logger) {
		l.hasPrefix = true
		l.prefix = pr
	}
}

// WithLevel sets log level. There is consts Info, Warn, Err, Debug, Trace which define log levels.
// By default level is Info.
// If level is less then called method expect, for example setted level if Info but called Logger.Warn("some text") method,
// then method wouldn't do anything.
func WithLevel(lev int) Setter {
	switch lev {
	case Info, Warn, Debug, Err, Trace:
		break
	default:
		lev = Info
	}

	return func(l *Logger) {
		l.level = uint8(lev)
	}
}

// WithTime sets time format for log messages. Without parameters it sets time.RFC3339 layout
func WithTime(layout ...string) Setter {
	return func(l *Logger) {
		l.hasTime = true
		if len(layout) > 0 {
			if layout[0] != "" {
				l.timeLayout = layout[0]
			} else {
				l.timeLayout = time.RFC3339
			}
		} else {
			l.timeLayout = time.RFC3339
		}
	}
}

// New return *Logger object
func New(options ...Setter) *Logger {
	logger := &Logger{
		writer: os.Stdout,
	}

	for _, opt := range options {
		opt(logger)
	}

	return logger
}
