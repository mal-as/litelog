package litelog

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func BenchmarkConcat(b *testing.B) {
	s := "some test string"
	p := "some prefix"
	for i := 0; i < b.N; i++ {
		concat(s, p)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	s := "some test string"
	p := "some prefix"
	for i := 0; i < b.N; i++ {
		builder(s, p)
	}
}

func BenchmarkFormater(b *testing.B) {
	s := "some test string"
	p := "some prefix"
	for i := 0; i < b.N; i++ {
		formater(s, p)
	}
}

func concat(msg, prefix string) string {
	msg = prefix + " " + msg
	msg = time.Now().Format(time.RFC1123) + ": " + msg
	return msg
}

func builder(msg, prefix string) string {
	sb := new(strings.Builder)
	sb.WriteString(time.Now().Format(time.RFC1123))
	sb.WriteString(": ")
	sb.WriteString(prefix)
	sb.WriteString(" ")
	sb.WriteString(msg)
	return sb.String()
}

func formater(msg, prefix string) string {
	return fmt.Sprintf("%s: %s %s", time.Now().Format(time.RFC1123), prefix, msg)
}
