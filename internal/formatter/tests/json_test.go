package formatter

import (
	"bytes"
	"strings"
	"testing"
)

func TestJSONFormatter(t *testing.T) {
	f, _ := New("json")
	data := struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}{Name: "test", Count: 42}

	var buf bytes.Buffer
	if err := f.Format(&buf, data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, `"name": "test"`) {
		t.Errorf("output missing name field: %s", output)
	}
	if !strings.Contains(output, `"count": 42`) {
		t.Errorf("output missing count field: %s", output)
	}
}
