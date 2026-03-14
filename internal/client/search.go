package client

import (
	"context"
	"net/url"
	"strconv"

	"github.com/G10xy/podcastindex-cli/pkg/models"
)

type SearchByTermOptions struct {
	Query    string
	Val      string
	Max      int
	APOnly   bool
	Clean    bool
	Similar  bool
	FullText bool
}

type SearchByTitleOptions struct {
	Query    string
	Val      string
	Max      int
	Clean    bool
	Similar  bool
	FullText bool
}

type SearchByPersonOptions struct {
	Query    string
	Max      int
	FullText bool
}

type SearchMusicByTermOptions struct {
	Query    string
	Val      string
	Max      int
	APOnly   bool
	Clean    bool
	FullText bool
}

func (c *Client) SearchByTerm(ctx context.Context, opts SearchByTermOptions) (*models.SearchByTermResponse, error) {
	params := url.Values{}
	params.Set("q", opts.Query)
	if opts.Val != "" {
		params.Set("val", opts.Val)
	}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	addBoolFlag(params, "aponly", opts.APOnly)
	addBoolFlag(params, "clean", opts.Clean)
	addBoolFlag(params, "similar", opts.Similar)
	addBoolFlag(params, "fulltext", opts.FullText)

	var result models.SearchByTermResponse
	resp, err := c.doRequest(ctx, "GET", "/search/byterm", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SearchByTitle(ctx context.Context, opts SearchByTitleOptions) (*models.SearchByTermResponse, error) {
	params := url.Values{}
	params.Set("q", opts.Query)
	if opts.Val != "" {
		params.Set("val", opts.Val)
	}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	addBoolFlag(params, "clean", opts.Clean)
	addBoolFlag(params, "similar", opts.Similar)
	addBoolFlag(params, "fulltext", opts.FullText)

	var result models.SearchByTermResponse
	resp, err := c.doRequest(ctx, "GET", "/search/bytitle", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SearchByPerson(ctx context.Context, opts SearchByPersonOptions) (*models.SearchByPersonResponse, error) {
	params := url.Values{}
	params.Set("q", opts.Query)
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	addBoolFlag(params, "fulltext", opts.FullText)

	var result models.SearchByPersonResponse
	resp, err := c.doRequest(ctx, "GET", "/search/byperson", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SearchMusicByTerm(ctx context.Context, opts SearchMusicByTermOptions) (*models.SearchByTermResponse, error) {
	params := url.Values{}
	params.Set("q", opts.Query)
	if opts.Val != "" {
		params.Set("val", opts.Val)
	}
	if opts.Max > 0 {
		params.Set("max", strconv.Itoa(opts.Max))
	}
	addBoolFlag(params, "aponly", opts.APOnly)
	addBoolFlag(params, "clean", opts.Clean)
	addBoolFlag(params, "fulltext", opts.FullText)

	var result models.SearchByTermResponse
	resp, err := c.doRequest(ctx, "GET", "/search/music/byterm", params, nil, true)
	if err != nil {
		return nil, err
	}
	if err := c.decodeResponse(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
