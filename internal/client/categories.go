package client

import (
	"context"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

func (c *Client) CategoriesList(ctx context.Context) (*models.CategoriesListResponse, error) {
	return get[models.CategoriesListResponse](c, ctx, "/categories/list", nil)
}
