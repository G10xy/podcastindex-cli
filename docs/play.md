# play

Stream an episode to a local media player. Supports mpv, vlc, and ffplay with auto-detection.

```
podcastindex play [flags]
```

## Flags

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--id` | int | no* | Episode ID |
| `--guid` | string | no* | Episode GUID |
| `--url` | string | no* | Direct enclosure URL (skips API lookup) |
| `--player` | string | no | Media player to use (default: auto-detect) |

*One of `--id`, `--guid`, or `--url` is required.

## Player detection

The command auto-detects an available media player in this order:

1. mpv
2. vlc
3. ffplay

You can override this by passing `--player` or setting the `player` key in the config file:

```bash
podcastindex config set player vlc
```

## Examples

```bash
# Play an episode by ID
podcastindex play --id 12345

# Play an episode by GUID
podcastindex play --guid "abc-def-123"

# Play a direct audio URL with a specific player
podcastindex play --url https://example.com/episode.mp3 --player vlc
```
