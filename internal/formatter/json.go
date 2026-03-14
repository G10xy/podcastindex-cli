package formatter

import (
	"encoding/json"
	"fmt"
	"io"
)

func formatJSON(w io.Writer, data interface{}) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling JSON: %w", err)
	}
	_, err = fmt.Fprintln(w, string(b))
	return err
}
