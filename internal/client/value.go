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

func (c *Client) ValueByFeedID(ctx context.Context, id int) (*models.ValueByFeedResponse, error) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(id))
	return getNoAuth[models.ValueByFeedResponse](c, ctx, "/value/byfeedid", params)
}

func (c *Client) ValueByFeedURL(ctx context.Context, feedURL string) (*models.ValueByFeedResponse, error) {
	params := url.Values{}
	params.Set("url", feedURL)
	return getNoAuth[models.ValueByFeedResponse](c, ctx, "/value/byfeedurl", params)
}

func (c *Client) ValueByPodcastGUID(ctx context.Context, guid string) (*models.ValueByFeedResponse, error) {
	params := url.Values{}
	params.Set("guid", guid)
	return getNoAuth[models.ValueByFeedResponse](c, ctx, "/value/bypodcastguid", params)
}

func (c *Client) ValueByEpisodeGUID(ctx context.Context, podcastGUID, episodeGUID string) (*models.ValueByEpisodeGUIDResponse, error) {
	params := url.Values{}
	params.Set("podcastguid", podcastGUID)
	params.Set("episodeguid", episodeGUID)
	return getNoAuth[models.ValueByEpisodeGUIDResponse](c, ctx, "/value/byepisodeguid", params)
}

func (c *Client) ValueBatchByEpisodeGUID(ctx context.Context, batch map[string][]string) (*models.ValueBatchByEpisodeGUIDResponse, error) {
	body, err := json.Marshal(batch)
	if err != nil {
		return nil, fmt.Errorf("encoding batch: %w", err)
	}
	return do[models.ValueBatchByEpisodeGUIDResponse](c, ctx, "POST", "/value/batch/byepisodeguid", nil, bytes.NewReader(body), false)
}
