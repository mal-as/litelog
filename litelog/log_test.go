package litelog

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestFatal(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(WithWriter(buf))
	l.Fatal("some text")
	if buf.String() != "[FATAL] some text\n" {
		t.Errorf("expected string '[FATAL] some text', but got - %s\n", buf.String())
	}
}

func TestFatalf(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(WithWriter(buf))
	l.Fatalf("some text %d - %v", 1, true)
	if buf.String() != "[FATAL] some text 1 - true" {
		t.Errorf("expected string '[FATAL] some text 1 - true', but got - %s\n", buf.String())
	}
}

func TestEmptyLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(WithWriter(buf))
	l.Println("some text")
	if buf.String() != "some text\n" {
		t.Errorf("expectedt 'some text', but got - %s", buf.String())
	}
}

func TestPrintln(t *testing.T) {
	in := "some text"
	buf := &bytes.Buffer{}

	l := New(
		WithLevel(Info),
		WithPrefix("<test>"),
		WithTime(time.RFC3339),
		WithWriter(buf),
	)

	l.Println(in)

	if !strings.HasPrefix(buf.String(), "2020") {
		t.Errorf("result don't have year prefix: %s\n", buf.String())
	}

	if !strings.Contains(buf.String(), l.prefix) {
		t.Errorf("result don't contains prefix %s: %s\n", l.prefix, buf.String())
	}

	if !strings.HasSuffix(buf.String(), in+"\n") {
		t.Errorf("result lost message: %s\n", buf.String())
	}
}

func TestPrintf(t *testing.T) {
	inTmpl := "some text %d - %v"
	in := "some text 1 - true"
	buf := &bytes.Buffer{}

	l := New(
		WithLevel(Info),
		WithTime(time.RFC3339),
		WithWriter(buf),
	)

	l.Printf(inTmpl, 1, true)

	if !strings.HasPrefix(buf.String(), "2020") {
		t.Errorf("result don't have year prefix: %s\n", buf.String())
	}

	if !strings.HasSuffix(buf.String(), in) {
		t.Errorf("result lost message: %s\n", buf.String())
	}
}

func TestLVL(t *testing.T) {
	for _, lvl := range []uint8{Info, Warn, Err, Debug, Trace} {
		lvlTestHelper(t, lvl)
	}
}

func lvlTestHelper(t *testing.T, lvl uint8) {
	var (
		NL     = true
		in     = "some text"
		inTmpl = "some text %d - %v"
		inFmt  = "some text 1 - true"
		inLN   = in + "\n"
		buf    = &bytes.Buffer{}
	)

	l := New(
		WithLevel(int(lvl)),
		WithPrefix("<test>"),
		WithTime(time.RFC3339),
		WithWriter(buf),
	)

	prefix := helperLVL(l, lvl, NL, in)
	res := buf.String()

	if !strings.HasPrefix(res, "2020") {
		t.Errorf("result don't have year prefix: %s\n", res)
	}

	if !strings.Contains(res, prefix) {
		t.Errorf("result don't contains prefix %s: %s\n", prefix, res)
	}

	if !strings.HasSuffix(res, inLN) {
		t.Errorf("result lost message: %s\n", res)
	}

	buf.Reset()
	NL = false
	prefix = helperLVL(l, lvl, NL, inTmpl, 1, true)
	res = buf.String()

	if !strings.HasPrefix(res, "2020") {
		t.Errorf("result don't have year prefix: %s\n", res)
	}

	if !strings.Contains(res, prefix) {
		t.Errorf("result don't contains prefix %s: %s\n", prefix, res)
	}

	if !strings.HasSuffix(res, inFmt) {
		t.Errorf("result have incorrect message: %s\n", res)
	}

	// fmt.Println(buf.String())
}

func helperLVL(l *Logger, lvl uint8, newLine bool, msg string, args ...interface{}) string {
	switch lvl {
	case Info:
		if newLine {
			l.Info(msg)
		} else {
			l.Infof(msg, args...)
		}
		return infoPrefix
	case Warn:
		if newLine {
			l.Warn(msg)
		} else {
			l.Warnf(msg, args...)
		}
		return warnPrefix
	case Err:
		if newLine {
			l.Error(msg)
		} else {
			l.Errorf(msg, args...)
		}
		return errPrefix
	case Debug:
		if newLine {
			l.Debug(msg)
		} else {
			l.Debugf(msg, args...)
		}
		return debugPrefix
	case Trace:
		if newLine {
			l.Trace(msg)
		} else {
			l.Tracef(msg, args...)
		}
		return tracePrefix
	}
	return ""
}

func TestDefaultLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(WithWriter(buf))
	in := "some text"
	l.Info(in)
	if buf.String() != "[INFO] "+in+"\n" {
		t.Errorf("expected - %s but got - %s\n", in, buf.String())
	}

	inTmpl := "some text %d - %v"
	inFmt := "some text 1 - true"

	buf.Reset()

	l.Infof(inTmpl, 1, true)
	if buf.String() != "[INFO] "+inFmt {
		t.Errorf("expected - %s but got - %s\n", inFmt, buf.String())
	}
}

func TestDefaultTime(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(WithWriter(buf), WithTime())

	l.Println("some text")

	if !strings.HasPrefix(buf.String(), "2020") {
		t.Errorf("result doesn't have time prefix: %s", buf.String())
	}
}
