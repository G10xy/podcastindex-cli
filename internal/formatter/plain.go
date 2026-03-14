package formatter

import (
	"fmt"
	"io"
	"strings"
)

func formatPlain(w io.Writer, data interface{}) error {
	switch v := data.(type) {
	case TableRow:
		fmt.Fprintln(w, strings.Join(v.TableRow(), "\t"))
	case []TableRow:
		for _, row := range v {
			fmt.Fprintln(w, strings.Join(row.TableRow(), "\t"))
		}
	default:
		return fmt.Errorf("plain format not supported for type %T", data)
	}
	return nil
}
