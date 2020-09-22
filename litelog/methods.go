package litelog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (l *Logger) format(msg string, lvlPrefix string) string {
	if !l.hasPrefix && !l.hasTime {
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

func (l *Logger) Println(msg string) {
	fmt.Fprintln(l.writer, l.format(msg, ""))
}

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

func (l *Logger) Info(msg string) {
	l.levelPrintln(Info, msg, infoPrefix)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.levelPrintf(Info, msg, infoPrefix, args...)
}

func (l *Logger) Warn(msg string) {
	l.levelPrintln(Warn, msg, warnPrefix)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.levelPrintf(Warn, msg, warnPrefix, args...)
}

func (l *Logger) Error(msg string) {
	l.levelPrintln(Err, msg, errPrefix)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.levelPrintf(Err, msg, errPrefix, args...)
}

func (l *Logger) Debug(msg string) {
	l.levelPrintln(Debug, msg, debugPrefix)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.levelPrintf(Debug, msg, debugPrefix, args...)
}

func (l *Logger) Trace(msg string) {
	l.levelPrintln(Trace, msg, tracePrefix)
}

func (l *Logger) Tracef(msg string, args ...interface{}) {
	l.levelPrintf(Trace, msg, tracePrefix, args...)
}

func (l *Logger) Fatal(msg string) {
	fmt.Fprintln(l.writer, l.format(msg, "[FATAL]"))
	os.Exit(1)
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	fmt.Fprintf(l.writer, l.format(msg, "[FATAL]"), args...)
	os.Exit(1)
}
