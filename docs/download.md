# download

Download episode audio files. Supports single episode downloads and batch downloading all episodes from a feed.

```
podcastindex download [flags]
```

## Single episode download

Download one episode by ID, GUID, or direct URL.

### Flags

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--id` | int | no* | Episode ID |
| `--guid` | string | no* | Episode GUID |
| `--url` | string | no* | Direct enclosure URL (skips API lookup) |
| `--dir` | string | no | Output directory (default: current directory) |
| `--filename` | string | no | Override output filename |

*One of `--id`, `--guid`, or `--url` is required.

### Examples

```bash
# Download by episode ID
podcastindex download --id 12345

# Download to a specific directory with a custom filename
podcastindex download --id 12345 --dir ~/podcasts --filename episode.mp3

# Download from a direct URL
podcastindex download --url https://example.com/episode.mp3
```

## Batch download

Download all episodes from a feed using `--all` with a feed identifier.

### Flags

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--all` | bool | yes | Enable batch mode |
| `--feedid` | int | no* | Feed ID |
| `--feedurl` | string | no* | Feed URL |
| `--podcastguid` | string | no* | Podcast GUID |
| `--dir` | string | no | Output directory (default: current directory) |
| `--workers` | int | no | Concurrent downloads (default: 3, max: 20) |

*One of `--feedid`, `--feedurl`, or `--podcastguid` is required when using `--all`.

### Examples

```bash
# Download all episodes from a feed by ID
podcastindex download --all --feedid 12345 --dir ~/podcasts

# Download all episodes with more concurrency
podcastindex download --all --feedurl https://example.com/feed.xml --workers 10

# Download all episodes by podcast GUID
podcastindex download --all --podcastguid "abc-def-123" --dir ~/podcasts
```
