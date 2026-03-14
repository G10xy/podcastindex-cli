package client

import (
	"context"
	"net/url"
	"strconv"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

type RecentEpisodesOptions struct {
	ExcludeString string
	Max           int
	Before        int
	FullText      bool
}

type RecentFeedsOptions struct {
	Lang   string
	Cat    string
	NotCat string
	Since  int64
	Max    int
}

type RecentNewFeedsOptions struct {
	FeedID string
	Since  int64
	Max    int
	Desc   bool
}

type RecentNewValueFeedsOptions struct {
	Since int64
	Max   int
}

type RecentDataOptions struct {
	Since int64
	Max   int
}

func (c *Client) RecentEpisodes(ctx context.Context, opts RecentEpisodesOptions) (*models.RecentEpisodesResponse, error) {
	params := url.Values{}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.ExcludeString != "" {
		params.Set("excludeString", opts.ExcludeString)
	}
	if opts.Before > 0 {
		params.Set("before", strconv.Itoa(opts.Before))
	}
	addBoolFlag(params, "fulltext", opts.FullText)
	return get[models.RecentEpisodesResponse](c, ctx, "/recent/episodes", params)
}

func (c *Client) RecentFeeds(ctx context.Context, opts RecentFeedsOptions) (*models.RecentFeedsResponse, error) {
	params := url.Values{}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.Since > 0 {
		params.Set("since", strconv.FormatInt(opts.Since, 10))
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
	return get[models.RecentFeedsResponse](c, ctx, "/recent/feeds", params)
}

func (c *Client) RecentNewFeeds(ctx context.Context, opts RecentNewFeedsOptions) (*models.RecentNewFeedsResponse, error) {
	params := url.Values{}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.Since > 0 {
		params.Set("since", strconv.FormatInt(opts.Since, 10))
	}
	if opts.FeedID != "" {
		params.Set("feedid", opts.FeedID)
	}
	addBoolFlag(params, "desc", opts.Desc)
	return get[models.RecentNewFeedsResponse](c, ctx, "/recent/newfeeds", params)
}

func (c *Client) RecentNewValueFeeds(ctx context.Context, opts RecentNewValueFeedsOptions) (*models.RecentNewValueFeedsResponse, error) {
	params := url.Values{}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.Since > 0 {
		params.Set("since", strconv.FormatInt(opts.Since, 10))
	}
	return get[models.RecentNewValueFeedsResponse](c, ctx, "/recent/newvaluefeeds", params)
}

func (c *Client) RecentData(ctx context.Context, opts RecentDataOptions) (*models.RecentDataResponse, error) {
	params := url.Values{}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.Since > 0 {
		params.Set("since", strconv.FormatInt(opts.Since, 10))
	}
	return get[models.RecentDataResponse](c, ctx, "/recent/data", params)
}

func (c *Client) RecentSoundbites(ctx context.Context, max int) (*models.RecentSoundbitesResponse, error) {
	params := url.Values{}
	if max > 0 {
		params.Set("max", strconv.Itoa(max))
	}
	return get[models.RecentSoundbitesResponse](c, ctx, "/recent/soundbites", params)
}
