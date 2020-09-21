package litelog

import (
	"io"
	"os"
	"time"
)

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

type Logger struct {
	hasTime    bool
	hasPrefix  bool
	level      uint8
	prefix     string
	timeLayout string
	writer     io.Writer
}

type setter func(l *Logger)

func WithWriter(wr io.Writer) setter {
	return func(l *Logger) {
		l.writer = wr
	}
}

func WithPrefix(pr string) setter {
	return func(l *Logger) {
		l.hasPrefix = true
		l.prefix = pr
	}
}

func WithLevel(lev int) setter {
	return func(l *Logger) {
		l.level = uint8(lev)
	}
}

func WithTime(layout ...string) setter {
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

func New(options ...setter) *Logger {
	logger := &Logger{
		writer: os.Stdout,
	}

	for _, opt := range options {
		opt(logger)
	}

	return logger
}
