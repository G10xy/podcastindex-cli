package client

import (
	"context"
	"net/url"
	"strconv"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

type HubPubNotifyOptions struct {
	URL string
	ID  int
}

func (c *Client) HubPubNotify(ctx context.Context, opts HubPubNotifyOptions) (*models.HubPubNotifyResponse, error) {
	params := url.Values{}
	if opts.ID > 0 {
		params.Set("id", strconv.Itoa(opts.ID))
	}
	if opts.URL != "" {
		params.Set("url", opts.URL)
	}
	return getNoAuth[models.HubPubNotifyResponse](c, ctx, "/hub/pubnotify", params)
}
