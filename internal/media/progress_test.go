package media

import (
	"bytes"
	"strings"
	"testing"
)

func TestProgressWriterTracksBytes(t *testing.T) {
	var buf bytes.Buffer
	var progressOutput bytes.Buffer

	pw := newProgressWriter(&buf, &progressOutput, 100)

	data := []byte("hello world") // 11 bytes
	n, err := pw.Write(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 11 {
		t.Fatalf("wrote %d bytes, want 11", n)
	}
	if pw.written != 11 {
		t.Fatalf("written = %d, want 11", pw.written)
	}

	// Verify the data was written to the underlying writer
	if buf.String() != "hello world" {
		t.Fatalf("underlying writer got %q, want %q", buf.String(), "hello world")
	}

	// Verify progress output contains percentage
	if !strings.Contains(progressOutput.String(), "%") {
		t.Errorf("progress output missing percentage: %q", progressOutput.String())
	}
}

func TestProgressWriterUnknownTotal(t *testing.T) {
	var buf bytes.Buffer
	var progressOutput bytes.Buffer

	pw := newProgressWriter(&buf, &progressOutput, 0) // unknown total

	pw.Write([]byte("data"))

	if !strings.Contains(progressOutput.String(), "downloaded") {
		t.Errorf("expected 'downloaded' in output for unknown total, got %q", progressOutput.String())
	}
}

func TestProgressWriterNilOutput(t *testing.T) {
	var buf bytes.Buffer
	pw := newProgressWriter(&buf, nil, 100)

	n, err := pw.Write([]byte("test"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 4 {
		t.Fatalf("wrote %d, want 4", n)
	}
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		input int64
		want  string
	}{
		{0, "0 B"},
		{512, "512 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{1048576, "1.0 MB"},
		{1073741824, "1.0 GB"},
	}
	for _, tt := range tests {
		got := formatBytes(tt.input)
		if got != tt.want {
			t.Errorf("formatBytes(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
