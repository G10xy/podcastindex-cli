package media

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// DownloadOptions configures a file download.
type DownloadOptions struct {
	URL      string
	Dir      string    // output directory, defaults to "."
	Filename string    // override filename, auto-derived if empty
	Output   io.Writer // for progress reporting (typically os.Stderr)
}

// Download fetches the URL to a local file, streaming to disk.
// Returns the path of the downloaded file.
func Download(ctx context.Context, httpClient *http.Client, opts DownloadOptions) (string, error) {
	if opts.Dir == "" {
		opts.Dir = "."
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opts.URL, nil)
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("downloading: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download failed: HTTP %d", resp.StatusCode)
	}

	filename := opts.Filename
	if filename == "" {
		filename = deriveFilename(opts.URL)
	}

	dest := filepath.Join(opts.Dir, filename)

	if err := os.MkdirAll(opts.Dir, 0755); err != nil {
		return "", fmt.Errorf("creating output directory: %w", err)
	}

	f, err := os.Create(dest)
	if err != nil {
		return "", fmt.Errorf("creating file: %w", err)
	}
	defer f.Close()

	var writer io.Writer = f
	if opts.Output != nil {
		pw := newProgressWriter(f, opts.Output, resp.ContentLength)
		writer = pw
	}

	if _, err := io.Copy(writer, resp.Body); err != nil {
		return "", fmt.Errorf("writing file: %w", err)
	}

	// Clear the progress line
	if opts.Output != nil {
		fmt.Fprintln(opts.Output)
	}

	return dest, nil
}

func deriveFilename(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err == nil {
		base := path.Base(parsed.Path)
		base = sanitizeFilename(base)
		if base != "" && base != "." && base != "/" {
			return base
		}
	}
	return fmt.Sprintf("episode-%d.mp3", time.Now().Unix())
}

func sanitizeFilename(name string) string {
	// Remove characters that are problematic in filenames
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
	)
	name = replacer.Replace(name)
	return strings.TrimSpace(name)
}
