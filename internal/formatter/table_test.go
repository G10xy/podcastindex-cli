package formatter

import (
	"bytes"
	"strings"
	"testing"
)

type mockRow struct {
	id   string
	name string
}

func (m mockRow) TableHeaders() []string { return []string{"ID", "NAME"} }
func (m mockRow) TableRow() []string     { return []string{m.id, m.name} }

func TestTableSingleRow(t *testing.T) {
	f, _ := New("table")
	var buf bytes.Buffer
	if err := f.Format(&buf, mockRow{"1", "Alice"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "ID") || !strings.Contains(out, "NAME") {
		t.Errorf("missing headers: %s", out)
	}
	if !strings.Contains(out, "Alice") {
		t.Errorf("missing row data: %s", out)
	}
}

func TestTableMultipleRows(t *testing.T) {
	f, _ := New("table")
	rows := []TableRow{mockRow{"1", "Alice"}, mockRow{"2", "Bob"}}
	var buf bytes.Buffer
	if err := f.Format(&buf, rows); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out := buf.String()
	if !strings.Contains(out, "Alice") || !strings.Contains(out, "Bob") {
		t.Errorf("missing row data: %s", out)
	}
}

func TestTableEmptySlice(t *testing.T) {
	f, _ := New("table")
	var buf bytes.Buffer
	if err := f.Format(&buf, []TableRow{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "No results found.") {
		t.Errorf("expected 'No results found.' message, got: %s", buf.String())
	}
}

func TestTableUnsupportedType(t *testing.T) {
	f, _ := New("table")
	var buf bytes.Buffer
	if err := f.Format(&buf, "not a table row"); err == nil {
		t.Error("expected error for unsupported type")
	}
}
