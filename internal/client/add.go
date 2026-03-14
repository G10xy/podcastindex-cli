package client

import (
	"context"
	"net/url"
	"strconv"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

type AddByFeedURLOptions struct {
	URL      string
	CHash    string
	ItunesID int
}

func (c *Client) AddByFeedURL(ctx context.Context, opts AddByFeedURLOptions) (*models.AddByFeedURLResponse, error) {
	params := url.Values{}
	params.Set("url", opts.URL)
	if opts.CHash != "" {
		params.Set("chash", opts.CHash)
	}
	if opts.ItunesID > 0 {
		params.Set("itunesid", strconv.Itoa(opts.ItunesID))
	}
	return get[models.AddByFeedURLResponse](c, ctx, "/add/byfeedurl", params)
}

func (c *Client) AddByItunesID(ctx context.Context, id int) (*models.AddByItunesIDResponse, error) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))
	return get[models.AddByItunesIDResponse](c, ctx, "/add/byitunesid", params)
}
