package formatter

import (
	"fmt"
	"io"
)

// Formatter defines how API results are rendered to the user.
type Formatter interface {
	Format(w io.Writer, data interface{}) error
}

// Func adapts a plain function into a Formatter.
type Func func(w io.Writer, data interface{}) error

func (f Func) Format(w io.Writer, data interface{}) error {
	return f(w, data)
}

// New returns a Formatter for the given format name.
// Supported: "table", "json", "plain".
func New(format string) (Formatter, error) {
	switch format {
	case "table":
		return Func(formatTable), nil
	case "json":
		return Func(formatJSON), nil
	case "plain":
		return Func(formatPlain), nil
	default:
		return nil, fmt.Errorf("unsupported output format: %q (supported: table, json, plain)", format)
	}
}
