package media

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"sync/atomic"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

// BatchDownloadOptions configures a batch download of multiple episodes.
type BatchDownloadOptions struct {
	Episodes []models.Episode
	Dir      string
	Workers  int       // concurrent downloads, default 3
	Output   io.Writer // progress output (stderr)
}

// BatchResult summarizes the outcome of a batch download.
type BatchResult struct {
	TotalFiles int
	TotalBytes int64
	Errors     []error
}

const maxWorkers = 20

// BatchDownload downloads multiple episodes concurrently using bounded parallelism.
func BatchDownload(ctx context.Context, httpClient *http.Client, opts BatchDownloadOptions) (*BatchResult, error) {
	if opts.Workers <= 0 {
		opts.Workers = 3
	}
	if opts.Workers > maxWorkers {
		opts.Workers = maxWorkers
	}

	total := len(opts.Episodes)
	if total == 0 {
		return &BatchResult{}, nil
	}

	sem := make(chan struct{}, opts.Workers)
	var (
		mu         sync.Mutex
		errors     []error
		wg         sync.WaitGroup
		completed  atomic.Int32
		downloaded atomic.Int32
		totalBytes atomic.Int64
	)

	for _, ep := range opts.Episodes {
		if ctx.Err() != nil {
			break
		}

		wg.Add(1)
		select {
		case sem <- struct{}{}: // acquire slot
		case <-ctx.Done():
			wg.Done()
			continue
		}

		go func(episode models.Episode) {
			defer wg.Done()
			defer func() { <-sem }() // release slot

			if episode.EnclosureURL == "" {
				n := completed.Add(1)
				if opts.Output != nil {
					fmt.Fprintf(opts.Output, "[%d/%d] Skipped %q (no enclosure URL)\n", n, total, episode.Title)
				}
				return
			}

			dest, err := Download(ctx, httpClient, DownloadOptions{
				URL: episode.EnclosureURL,
				Dir: opts.Dir,
			})
			n := completed.Add(1)

			if err != nil {
				mu.Lock()
				errors = append(errors, fmt.Errorf("episode %q: %w", episode.Title, err))
				mu.Unlock()
				if opts.Output != nil {
					fmt.Fprintf(opts.Output, "[%d/%d] Failed %q: %v\n", n, total, episode.Title, err)
				}
				return
			}

			downloaded.Add(1)
			var size int64
			if fi, err := os.Stat(dest); err == nil {
				size = fi.Size()
			}
			totalBytes.Add(size)

			if opts.Output != nil {
				fmt.Fprintf(opts.Output, "[%d/%d] Downloaded %q (%s)\n", n, total, episode.Title, FormatBytes(size))
			}
		}(ep)
	}

	wg.Wait()

	return &BatchResult{
		TotalFiles: int(downloaded.Load()),
		TotalBytes: totalBytes.Load(),
		Errors:     errors,
	}, nil
}
