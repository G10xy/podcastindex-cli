package formatter

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

// TableRow is implemented by types that can render themselves as table rows.
type TableRow interface {
	TableHeaders() []string
	TableRow() []string
}

func formatTable(w io.Writer, data interface{}) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	switch v := data.(type) {
	case TableRow:
		fmt.Fprintln(tw, strings.Join(v.TableHeaders(), "\t"))
		fmt.Fprintln(tw, strings.Join(v.TableRow(), "\t"))
	case []TableRow:
		if len(v) == 0 {
			fmt.Fprintln(w, "No results found.")
			return nil
		}
		fmt.Fprintln(tw, strings.Join(v[0].TableHeaders(), "\t"))
		for _, row := range v {
			fmt.Fprintln(tw, strings.Join(row.TableRow(), "\t"))
		}
	default:
		return fmt.Errorf("table format not supported for type %T", data)
	}

	return nil
}
