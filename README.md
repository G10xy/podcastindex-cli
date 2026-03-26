# podcastindex-cli

CLI for interfacing with the [PodcastIndex.org](https://podcastindex.org/) API.

## Quick Start

```bash
# Install
go install github.com/G10xy/podcastindex-cli/cmd/podcastindex@latest

# Configure your API credentials (get them at https://api.podcastindex.org/)
podcastindex config set api_key <your-key>
podcastindex config set api_secret <your-secret>

# Search for a podcast
podcastindex search byterm -q "linux"

# Get episodes from a feed
podcastindex episodes byfeedid --id 75075

# Play an episode
podcastindex play --id 12345

# Download an episode
podcastindex download --id 12345 --dir ~/podcasts
```

## Installation

```bash
go install github.com/G10xy/podcastindex-cli/cmd/podcastindex@latest
```

Or download a pre-built binary from the [releases page](https://github.com/G10xy/podcastindex-cli/releases).

## Configuration

Set your API credentials via flags, environment variables, or a config file.

**Flags (available on all commands):**

```
-k, --api-key       PodcastIndex API key
-s, --api-secret    PodcastIndex API secret
-o, --output        Output format: table, json, plain (default: table)
```

**Environment variables:**

```
PODCASTINDEX_API_KEY
PODCASTINDEX_API_SECRET
```

**Config file** (`~/.config/podcastindex-cli/config.yaml`):

```bash
podcastindex config set api_key <your-key>
podcastindex config set api_secret <your-secret>
```

## Commands

```
podcastindex
├── search
│   ├── byterm          Search podcasts by term
│   ├── bytitle         Search podcasts by title
│   ├── byperson        Search by person
│   └── music           Search music feeds
├── podcasts
│   ├── byfeedid        Lookup podcast by feed ID
│   ├── byfeedurl       Lookup podcast by feed URL
│   ├── byitunesid      Lookup podcast by iTunes ID
│   ├── byguid          Lookup podcast by GUID
│   ├── bytag           Lookup podcasts by tag
│   ├── bymedium        Lookup podcasts by medium
│   ├── trending        Get trending podcasts
│   ├── dead            Get dead podcasts
│   └── batch           Batch lookup podcasts
├── episodes
│   ├── byfeedid        Get episodes by feed ID
│   ├── byfeedurl       Get episodes by feed URL
│   ├── bypodcastguid   Get episodes by podcast GUID
│   ├── byitunesid      Get episodes by iTunes ID
│   ├── byid            Get episode by ID
│   ├── byguid          Get episode by GUID
│   ├── live            Get live episodes
│   └── random          Get random episodes
├── recent
│   ├── episodes        Get recent episodes
│   ├── feeds           Get recent feeds
│   ├── newfeeds        Get new feeds
│   ├── newvaluefeeds   Get new value-enabled feeds
│   ├── data            Get recent data
│   └── soundbites      Get recent soundbites
├── value
│   ├── byfeedid        Get value info by feed ID
│   ├── byfeedurl       Get value info by feed URL
│   ├── bypodcastguid   Get value info by podcast GUID
│   ├── byepisodeguid   Get value info by episode GUID
│   └── batch           Batch lookup value info
├── add
│   ├── byfeedurl       Add podcast by feed URL
│   └── byitunesid      Add podcast by iTunes ID
├── play                Stream an episode to a local media player
├── download            Download episode audio files
├── categories
│   └── list            List categories
├── stats
│   └── current         Get current stats
├── hub
│   └── pubnotify       Send a PubSub notification
├── config
│   ├── show            Show current configuration
│   └── set             Set a configuration value
└── version             Print the CLI version
```

## Documentation

- [add](docs/add.md) — Add podcasts to the index
- [play](docs/play.md) — Stream episodes to a media player
- [download](docs/download.md) — Download episode audio files

## Contributing

```bash
# Clone the repository
git clone https://github.com/G10xy/podcastindex-cli.git
cd podcastindex-cli

# Run tests
go test ./...

# Build
go build -o podcastindex ./cmd/podcastindex
```

## License

See [LICENSE](LICENSE).
