package litelog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (l *Logger) format(msg string, lvlPrefix string) string {
	if !l.hasPrefix && !l.hasTime {
		if lvlPrefix != "" {
			sb := new(strings.Builder)
			sb.WriteString(lvlPrefix)
			sb.WriteString(" ")
			sb.WriteString(msg)
			return sb.String()
		}
		return msg
	}

	sb := new(strings.Builder)

	if l.hasTime {
		sb.WriteString(time.Now().Format(l.timeLayout))
		sb.WriteString(": ")
	}

	if lvlPrefix != "" {
		sb.WriteString(lvlPrefix)
		sb.WriteString(" ")
	}

	if l.hasPrefix {
		sb.WriteString(l.prefix)
		sb.WriteString(" ")
	}

	sb.WriteString(msg)
	return sb.String()
}

// Println prints message in provided io.Writer
func (l *Logger) Println(msg string) {
	fmt.Fprintln(l.writer, l.format(msg, ""))
}

// Printf prints formatted message in provided io.Writer
func (l *Logger) Printf(msg string, args ...interface{}) {
	fmt.Fprintf(l.writer, l.format(msg, ""), args...)
}

func (l *Logger) levelPrintln(lvl uint8, msg, lvlPrefix string) {
	if l.level >= lvl {
		fmt.Fprintln(l.writer, l.format(msg, lvlPrefix))
	}
}

func (l *Logger) levelPrintf(lvl uint8, msg, lvlPrefix string, args ...interface{}) {
	if l.level >= lvl {
		fmt.Fprintf(l.writer, l.format(msg, lvlPrefix), args...)
	}
}

// Info prints message if current level is more or equal Info with "[INFO]" prefix
func (l *Logger) Info(msg string) {
	l.levelPrintln(Info, msg, infoPrefix)
}

// Infof prints formatted message if current level is more or equal Info with "[INFO]" prefix
func (l *Logger) Infof(msg string, args ...interface{}) {
	l.levelPrintf(Info, msg, infoPrefix, args...)
}

// Warn prints message if current level is more or equal Warn with "[WARN]" prefix
func (l *Logger) Warn(msg string) {
	l.levelPrintln(Warn, msg, warnPrefix)
}

// Warnf prints formatted message if current level is more or equal Warn with "[WARN]" prefix
func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.levelPrintf(Warn, msg, warnPrefix, args...)
}

// Error prints message if current level is more or equal Err with "[ERROR]" prefix
func (l *Logger) Error(msg string) {
	l.levelPrintln(Err, msg, errPrefix)
}

// Errorf prints formatted message if current level is more or equal Err with "[ERROR]" prefix
func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.levelPrintf(Err, msg, errPrefix, args...)
}

// Debug prints message if current level is more or equal Debug with "[DEBUG]" prefix
func (l *Logger) Debug(msg string) {
	l.levelPrintln(Debug, msg, debugPrefix)
}

// Debugf prints formatted message if current level is more or equal Debug with "[DEBUG]" prefix
func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.levelPrintf(Debug, msg, debugPrefix, args...)
}

// Trace prints message if current level is more or equal Trace with "[TRACE]" prefix
func (l *Logger) Trace(msg string) {
	l.levelPrintln(Trace, msg, tracePrefix)
}

// Tracef prints formatted message if current level is more or equal Trace with "[TRACE]" prefix
func (l *Logger) Tracef(msg string, args ...interface{}) {
	l.levelPrintf(Trace, msg, tracePrefix, args...)
}

// Fatal prints message and exit with status code 1
func (l *Logger) Fatal(msg string) {
	fmt.Fprintln(l.writer, l.format(msg, "[FATAL]"))
	os.Exit(1)
}

// Fatalf prints formatted message and exit with status code 1
func (l *Logger) Fatalf(msg string, args ...interface{}) {
	fmt.Fprintf(l.writer, l.format(msg, "[FATAL]"), args...)
	os.Exit(1)
}
