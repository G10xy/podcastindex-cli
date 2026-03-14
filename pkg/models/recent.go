package models

import (
	"strconv"
	"time"
)

type Soundbite struct {
	EnclosureURL string `json:"enclosureUrl"`
	Title        string `json:"title"`
	EpisodeTitle string `json:"episodeTitle"`
	FeedTitle    string `json:"feedTitle"`
	FeedURL      string `json:"feedUrl"`
	EpisodeID    int    `json:"episodeId"`
	FeedID       int    `json:"feedId"`
	StartTime    int    `json:"startTime"`
	Duration     int    `json:"duration"`
}

func (s Soundbite) TableHeaders() []string {
	return []string{"EPISODE ID", "TITLE", "FEED", "START", "DURATION"}
}

func (s Soundbite) TableRow() []string {
	return []string{
		strconv.Itoa(s.EpisodeID),
		s.Title,
		s.FeedTitle,
		strconv.Itoa(s.StartTime),
		strconv.Itoa(s.Duration),
	}
}

type RecentDataItem struct {
	FeedURL      string `json:"feedUrl,omitempty"`
	FeedTitle    string `json:"feedTitle,omitempty"`
	EpisodeTitle string `json:"episodeTitle,omitempty"`
	DatePublished int64 `json:"datePublished,omitempty"`
	FeedID       int    `json:"feedId,omitempty"`
	EpisodeID    int    `json:"episodeId,omitempty"`
}

type NewFeed struct {
	URL         string `json:"url"`
	Status      string `json:"status"`
	ContentHash string `json:"contentHash"`
	Language    string `json:"language"`
	TimeAdded   int64  `json:"timeAdded"`
	ID          int    `json:"id"`
}

func (n NewFeed) TableHeaders() []string {
	return []string{"ID", "URL", "STATUS", "LANGUAGE", "ADDED"}
}

func (n NewFeed) TableRow() []string {
	added := ""
	if n.TimeAdded > 0 {
		added = time.Unix(n.TimeAdded, 0).Format("2006-01-02 15:04")
	}
	return []string{
		strconv.Itoa(n.ID),
		n.URL,
		n.Status,
		n.Language,
		added,
	}
}

type RecentEpisodesResponse struct {
	APIResponse
	Items []Episode `json:"items"`
	Max   int       `json:"max,omitempty"`
}

type RecentFeedsResponse struct {
	APIResponse
	Feeds []Feed `json:"feeds"`
	Since int64  `json:"since,omitempty"`
}

type RecentNewFeedsResponse struct {
	APIResponse
	Feeds []NewFeed `json:"feeds"`
	Since int64     `json:"since,omitempty"`
}

type RecentNewValueFeedsResponse struct {
	APIResponse
	Feeds []Feed `json:"feeds"`
	Since int64  `json:"since,omitempty"`
}

type RecentDataResponse struct {
	APIResponse
	Data struct {
		Feeds []RecentDataItem `json:"feeds"`
		Items []RecentDataItem `json:"items"`
	} `json:"data"`
}

type RecentSoundbitesResponse struct {
	APIResponse
	Items []Soundbite `json:"items"`
}
