# add

Add podcasts to the PodcastIndex. Requires a write-enabled API key.

## Subcommands

### byfeedurl

Add a podcast by its RSS feed URL.

```
podcastindex add byfeedurl --url <feed-url> [flags]
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--url` | string | yes | Podcast feed URL |
| `--chash` | string | no | MD5 hash of the feed content for duplicate checking |
| `--itunesid` | int | no | iTunes ID to associate if none already exists |

**Examples:**

```bash
# Add a podcast by feed URL
podcastindex add byfeedurl --url https://example.com/feed.xml

# Add with iTunes ID association
podcastindex add byfeedurl --url https://example.com/feed.xml --itunesid 123456
```

### byitunesid

Add a podcast by its iTunes ID.

```
podcastindex add byitunesid --id <itunes-id>
```

**Flags:**

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--id` | int | yes | iTunes ID |

**Examples:**

```bash
podcastindex add byitunesid --id 123456
```
