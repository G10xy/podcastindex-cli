package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

type PodcastsByTagOptions struct {
	Max                   int
	StartAt               int
	PodcastValue          bool
	PodcastValueTimeSplit bool
}

type PodcastsByMediumOptions struct {
	Medium string
	Max    int
}

type PodcastsTrendingOptions struct {
	Lang   string
	Cat    string
	NotCat string
	Since  int64
	Max    int
}

func (c *Client) PodcastByFeedID(ctx context.Context, id int) (*models.PodcastByFeedResponse, error) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))

	var result models.PodcastByFeedResponse
	resp, err := c.doRequest(ctx, "GET", "/podcasts/byfeedid", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PodcastByFeedURL(ctx context.Context, feedURL string) (*models.PodcastByFeedResponse, error) {
	params := url.Values{}
	params.Set("url", feedURL)

	var result models.PodcastByFeedResponse
	resp, err := c.doRequest(ctx, "GET", "/podcasts/byfeedurl", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PodcastByItunesID(ctx context.Context, itunesID int) (*models.PodcastByFeedResponse, error) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(itunesID))

	var result models.PodcastByFeedResponse
	resp, err := c.doRequest(ctx, "GET", "/podcasts/byitunesid", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PodcastByGUID(ctx context.Context, guid string) (*models.PodcastByFeedResponse, error) {
	params := url.Values{}
	params.Set("guid", guid)

	var result models.PodcastByFeedResponse
	resp, err := c.doRequest(ctx, "GET", "/podcasts/byguid", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PodcastsByTag(ctx context.Context, opts PodcastsByTagOptions) (*models.PodcastsByTagResponse, error) {
	params := url.Values{}
	addBoolFlag(params, "podcast-value", opts.PodcastValue)
	addBoolFlag(params, "podcast-valueTimeSplit", opts.PodcastValueTimeSplit)
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	if opts.StartAt > 0 {
		params.Set("start_at", strconv.Itoa(opts.StartAt))
	}

	var result models.PodcastsByTagResponse
	resp, err := c.doRequest(ctx, "GET", "/podcasts/bytag", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PodcastsByMedium(ctx context.Context, opts PodcastsByMediumOptions) (*models.PodcastsByMediumResponse, error) {
	params := url.Values{}
	params.Set("medium", opts.Medium)
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}

	var result models.PodcastsByMediumResponse
	resp, err := c.doRequest(ctx, "GET", "/podcasts/bymedium", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PodcastsTrending(ctx context.Context, opts PodcastsTrendingOptions) (*models.PodcastsTrendingResponse, error) {
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

	var result models.PodcastsTrendingResponse
	resp, err := c.doRequest(ctx, "GET", "/podcasts/trending", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PodcastsDead(ctx context.Context) (*models.PodcastsDeadResponse, error) {
	var result models.PodcastsDeadResponse
	resp, err := c.doRequest(ctx, "GET", "/podcasts/dead", nil, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) PodcastsBatchByGUID(ctx context.Context, guids []string) (*models.PodcastsBatchByGUIDResponse, error) {
	body, err := json.Marshal(guids)
	if err != nil {
		return nil, fmt.Errorf("encoding GUIDs: %w", err)
	}

	var result models.PodcastsBatchByGUIDResponse
	resp, err := c.doRequest(ctx, "POST", "/podcasts/batch/byguid", nil, bytes.NewReader(body), true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
