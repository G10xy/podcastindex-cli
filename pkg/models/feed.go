package models

import (
	"fmt"
	"strconv"
	"time"
)

type Feed struct {
	ItunesID               *int              `json:"itunesId"`
	Categories             map[string]string `json:"categories"`
	OwnerName              string            `json:"ownerName"`
	Artwork                string            `json:"artwork"`
	Link                   string            `json:"link"`
	Description            string            `json:"description"`
	Author                 string            `json:"author"`
	PodcastGUID            string            `json:"podcastGuid"`
	Image                  string            `json:"image"`
	OriginalURL            string            `json:"originalUrl"`
	ContentType            string            `json:"contentType"`
	Generator              string            `json:"generator"`
	Language               string            `json:"language"`
	Medium                 string            `json:"medium"`
	Title                  string            `json:"title"`
	URL                    string            `json:"url"`
	LastCrawlTime          int64             `json:"lastCrawlTime"`
	LastParseTime          int64             `json:"lastParseTime"`
	LastUpdateTime         int64             `json:"lastUpdateTime"`
	ImageURLHash           int64             `json:"imageUrlHash"`
	NewestItemPubdate      int64             `json:"newestItemPubdate"`
	LastGoodHttpStatusTime int64             `json:"lastGoodHttpStatusTime"`
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
	Query map[string]string `json:"query"`
	APIResponse
	Feed Feed `json:"feed"`
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
	Query string `json:"query"`
	APIResponse
	Feeds []Feed `json:"feeds"`
}

func formatDuration(seconds *int) string {
	if seconds == nil {
		return ""
	}
	m := *seconds / 60
	s := *seconds % 60
	return fmt.Sprintf("%d:%02d", m, s)
}
