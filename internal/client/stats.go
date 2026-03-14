package client

import (
	"context"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

func (c *Client) StatsCurrent(ctx context.Context) (*models.StatsResponse, error) {
	return get[models.StatsResponse](c, ctx, "/stats/current", nil)
}
