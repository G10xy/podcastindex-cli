package media

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
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

	// Validate and create output directory before making the HTTP request
	if err := os.MkdirAll(opts.Dir, 0755); err != nil {
		return "", fmt.Errorf("creating output directory: %w", err)
	}

	filename := opts.Filename
	if filename == "" {
		filename = deriveFilename(opts.URL)
	}

	dest := filepath.Join(opts.Dir, filename)

	// Path traversal check: ensure dest stays within opts.Dir
	absDir, err := filepath.Abs(opts.Dir)
	if err != nil {
		return "", fmt.Errorf("resolving output directory: %w", err)
	}
	absDest, err := filepath.Abs(dest)
	if err != nil {
		return "", fmt.Errorf("resolving destination path: %w", err)
	}
	if !strings.HasPrefix(absDest, absDir+string(os.PathSeparator)) && absDest != absDir {
		return "", fmt.Errorf("filename %q resolves outside output directory", filename)
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

	// Use O_CREATE|O_EXCL to avoid silent overwrites; append suffix on collision
	f, dest, err := createFile(dest)
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

// createFile attempts to create the file exclusively. If it already exists,
// it appends a timestamp suffix to avoid overwriting.
func createFile(dest string) (*os.File, string, error) {
	f, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err == nil {
		return f, dest, nil
	}
	if !isExist(err) {
		return nil, "", err
	}
	// File exists — add timestamp suffix
	ext := filepath.Ext(dest)
	base := strings.TrimSuffix(dest, ext)
	dest = fmt.Sprintf("%s_%d%s", base, time.Now().UnixNano(), ext)
	f, err = os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return nil, "", err
	}
	return f, dest, nil
}

func isExist(err error) bool {
	return err != nil && (os.IsExist(err) || errors.Is(err, fs.ErrExist))
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
