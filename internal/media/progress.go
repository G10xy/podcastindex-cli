package media

import (
	"fmt"
	"io"
	"time"
)

type progressWriter struct {
	writer    io.Writer
	output    io.Writer
	total     int64
	written   int64
	startTime time.Time
}

func newProgressWriter(w, output io.Writer, total int64) *progressWriter {
	return &progressWriter{
		writer:    w,
		output:    output,
		total:     total,
		startTime: time.Now(),
	}
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	n, err := pw.writer.Write(p)
	pw.written += int64(n)

	if pw.output != nil {
		elapsed := time.Since(pw.startTime).Seconds()
		speed := float64(pw.written) / elapsed / (1024 * 1024) // MB/s

		if pw.total > 0 {
			pct := float64(pw.written) / float64(pw.total) * 100
			fmt.Fprintf(pw.output, "\r  %s / %s (%.0f%%) %.1f MB/s",
				formatBytes(pw.written), formatBytes(pw.total), pct, speed)
		} else {
			fmt.Fprintf(pw.output, "\r  %s downloaded  %.1f MB/s",
				formatBytes(pw.written), speed)
		}
	}

	return n, err
}

func formatBytes(b int64) string {
	const (
		kb = 1024
		mb = kb * 1024
		gb = mb * 1024
	)
	switch {
	case b >= gb:
		return fmt.Sprintf("%.1f GB", float64(b)/float64(gb))
	case b >= mb:
		return fmt.Sprintf("%.1f MB", float64(b)/float64(mb))
	case b >= kb:
		return fmt.Sprintf("%.1f KB", float64(b)/float64(kb))
	default:
		return fmt.Sprintf("%d B", b)
	}
}
