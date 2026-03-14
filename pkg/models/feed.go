package models

import (
	"fmt"
	"strconv"
	"time"
)

type Feed struct {
	PodcastGUID            string            `json:"podcastGuid"`
	Title                  string            `json:"title"`
	URL                    string            `json:"url"`
	OriginalURL            string            `json:"originalUrl"`
	Link                   string            `json:"link"`
	Description            string            `json:"description"`
	Author                 string            `json:"author"`
	OwnerName              string            `json:"ownerName"`
	Image                  string            `json:"image"`
	Artwork                string            `json:"artwork"`
	ContentType            string            `json:"contentType"`
	Generator              string            `json:"generator"`
	Language               string            `json:"language"`
	Medium                 string            `json:"medium"`
	Categories             map[string]string `json:"categories"`
	LastUpdateTime         int64             `json:"lastUpdateTime"`
	LastCrawlTime          int64             `json:"lastCrawlTime"`
	LastParseTime          int64             `json:"lastParseTime"`
	LastGoodHttpStatusTime int64             `json:"lastGoodHttpStatusTime"`
	ImageURLHash           int64             `json:"imageUrlHash"`
	NewestItemPubdate      int64             `json:"newestItemPubdate"`
	ItunesID               *int              `json:"itunesId"`
	ID                     int               `json:"id"`
	LastHttpStatus         int               `json:"lastHttpStatus"`
	Type                   int               `json:"type"`
	Dead                   int               `json:"dead"`
	EpisodeCount           int               `json:"episodeCount"`
	CrawlErrors            int               `json:"crawlErrors"`
	ParseErrors            int               `json:"parseErrors"`
	Locked                 int               `json:"locked"`
	Explicit               bool              `json:"explicit"`
}

func (f Feed) TableHeaders() []string {
	return []string{"ID", "TITLE", "AUTHOR", "LANGUAGE", "EPISODES", "LAST UPDATED"}
}

func (f Feed) TableRow() []string {
	updated := ""
	if f.LastUpdateTime > 0 {
		updated = time.Unix(f.LastUpdateTime, 0).Format("2006-01-02")
	}
	return []string{
		strconv.Itoa(f.ID),
		f.Title,
		f.Author,
		f.Language,
		strconv.Itoa(f.EpisodeCount),
		updated,
	}
}

type PodcastByFeedResponse struct {
	APIResponse
	Feed  Feed              `json:"feed"`
	Query map[string]string `json:"query"`
}

type PodcastsByTagResponse struct {
	APIResponse
	Feeds       []Feed `json:"feeds"`
	NextStartAt int    `json:"nextStartAt,omitempty"`
}

type PodcastsByMediumResponse struct {
	APIResponse
	Feeds []Feed `json:"feeds"`
}

type PodcastsTrendingResponse struct {
	APIResponse
	Feeds []Feed `json:"feeds"`
	Since int64  `json:"since,omitempty"`
}

type PodcastsDeadResponse struct {
	APIResponse
	Feeds []Feed `json:"feeds"`
}

type PodcastsBatchByGUIDResponse struct {
	APIResponse
	Feeds []Feed `json:"feeds"`
}

type SearchByTermResponse struct {
	APIResponse
	Feeds []Feed `json:"feeds"`
	Query string `json:"query"`
}

func formatDuration(seconds *int) string {
	if seconds == nil {
		return ""
	}
	m := *seconds / 60
	s := *seconds % 60
	return fmt.Sprintf("%d:%02d", m, s)
}
