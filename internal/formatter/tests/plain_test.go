package formatter

import (
	"bytes"
	"strings"
	"testing"
)

func TestPlainSingleRow(t *testing.T) {
	f, _ := New("plain")
	var buf bytes.Buffer
	if err := f.Format(&buf, mockRow{"1", "Alice"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := strings.TrimSpace(buf.String())
	if out != "1\tAlice" {
		t.Errorf("got %q, want %q", out, "1\tAlice")
	}
}

func TestPlainMultipleRows(t *testing.T) {
	f, _ := New("plain")
	rows := []TableRow{mockRow{"1", "Alice"}, mockRow{"2", "Bob"}}
	var buf bytes.Buffer
	if err := f.Format(&buf, rows); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("got %d lines, want 2", len(lines))
	}
	if lines[0] != "1\tAlice" {
		t.Errorf("line 0 = %q", lines[0])
	}
	if lines[1] != "2\tBob" {
		t.Errorf("line 1 = %q", lines[1])
	}
}

func TestPlainNoHeaders(t *testing.T) {
	f, _ := New("plain")
	var buf bytes.Buffer
	f.Format(&buf, mockRow{"1", "Alice"})
	if strings.Contains(buf.String(), "ID") || strings.Contains(buf.String(), "NAME") {
		t.Error("plain output should not contain headers")
	}
}

func TestPlainUnsupportedType(t *testing.T) {
	f, _ := New("plain")
	var buf bytes.Buffer
	if err := f.Format(&buf, 42); err == nil {
		t.Error("expected error for unsupported type")
	}
}
