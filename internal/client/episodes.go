package client

import (
	"context"
	"net/url"
	"strconv"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

type EpisodesByFeedIDOptions struct {
	ID        string
	Enclosure string
	Since     int64
	Max       int
	FullText  bool
	Newest    bool
}

type EpisodesByFeedURLOptions struct {
	URL      string
	Since    int64
	Max      int
	FullText bool
}

type EpisodesByPodcastGUIDOptions struct {
	GUID     string
	Since    int64
	Max      int
	FullText bool
}

type EpisodesByItunesIDOptions struct {
	Enclosure string
	Since     int64
	ID        int
	Max       int
	FullText  bool
}

type EpisodeByGUIDOptions struct {
	GUID        string
	FeedURL     string
	FeedID      string
	PodcastGUID string
	FullText    bool
}

type EpisodesRandomOptions struct {
	Lang     string
	Cat      string
	NotCat   string
	Max      int
	FullText bool
}

func (c *Client) EpisodesByFeedID(ctx context.Context, opts EpisodesByFeedIDOptions) (*models.EpisodesByFeedResponse, error) {
	params := url.Values{}
	params.Set("id", opts.ID)
	if opts.Since > 0 {
		params.Set("since", strconv.FormatInt(opts.Since, 10))
	}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.Enclosure != "" {
		params.Set("enclosure", opts.Enclosure)
	}
	addBoolFlag(params, "fulltext", opts.FullText)
	addBoolFlag(params, "newest", opts.Newest)
	return get[models.EpisodesByFeedResponse](c, ctx, "/episodes/byfeedid", params)
}

func (c *Client) EpisodesByFeedURL(ctx context.Context, opts EpisodesByFeedURLOptions) (*models.EpisodesByFeedResponse, error) {
	params := url.Values{}
	params.Set("url", opts.URL)
	if opts.Since > 0 {
		params.Set("since", strconv.FormatInt(opts.Since, 10))
	}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	addBoolFlag(params, "fulltext", opts.FullText)
	return get[models.EpisodesByFeedResponse](c, ctx, "/episodes/byfeedurl", params)
}

func (c *Client) EpisodesByPodcastGUID(ctx context.Context, opts EpisodesByPodcastGUIDOptions) (*models.EpisodesByFeedResponse, error) {
	params := url.Values{}
	params.Set("guid", opts.GUID)
	if opts.Since > 0 {
		params.Set("since", strconv.FormatInt(opts.Since, 10))
	}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	addBoolFlag(params, "fulltext", opts.FullText)
	return get[models.EpisodesByFeedResponse](c, ctx, "/episodes/bypodcastguid", params)
}

func (c *Client) EpisodesByItunesID(ctx context.Context, opts EpisodesByItunesIDOptions) (*models.EpisodesByFeedResponse, error) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(opts.ID))
	if opts.Since > 0 {
		params.Set("since", strconv.FormatInt(opts.Since, 10))
	}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.Enclosure != "" {
		params.Set("enclosure", opts.Enclosure)
	}
	addBoolFlag(params, "fulltext", opts.FullText)
	return get[models.EpisodesByFeedResponse](c, ctx, "/episodes/byitunesid", params)
}

func (c *Client) EpisodeByID(ctx context.Context, id int, fulltext bool) (*models.EpisodeByIDResponse, error) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))
	addBoolFlag(params, "fulltext", fulltext)
	return get[models.EpisodeByIDResponse](c, ctx, "/episodes/byid", params)
}

func (c *Client) EpisodeByGUID(ctx context.Context, opts EpisodeByGUIDOptions) (*models.EpisodeByIDResponse, error) {
	params := url.Values{}
	params.Set("guid", opts.GUID)
	if opts.FeedURL != "" {
		params.Set("feedurl", opts.FeedURL)
	}
	if opts.FeedID != "" {
		params.Set("feedid", opts.FeedID)
	}
	if opts.PodcastGUID != "" {
		params.Set("podcastguid", opts.PodcastGUID)
	}
	addBoolFlag(params, "fulltext", opts.FullText)
	return get[models.EpisodeByIDResponse](c, ctx, "/episodes/byguid", params)
}

func (c *Client) EpisodesLive(ctx context.Context, max int) (*models.EpisodesLiveResponse, error) {
	params := url.Values{}
	if max > 0 {
		params.Set("max", strconv.Itoa(max))
	}
	return get[models.EpisodesLiveResponse](c, ctx, "/episodes/live", params)
}

func (c *Client) EpisodesRandom(ctx context.Context, opts EpisodesRandomOptions) (*models.EpisodesRandomResponse, error) {
	params := url.Values{}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.Lang != "" {
		params.Set("lang", opts.Lang)
	}
	if opts.Cat != "" {
		params.Set("cat", opts.Cat)
	}
	if opts.NotCat != "" {
		params.Set("notcat", opts.NotCat)
	}
	addBoolFlag(params, "fulltext", opts.FullText)
	return get[models.EpisodesRandomResponse](c, ctx, "/episodes/random", params)
}
