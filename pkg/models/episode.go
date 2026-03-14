package models

import (
	"strconv"
	"time"
)

type Transcript struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type Episode struct {
	Title           string       `json:"title"`
	Link            string       `json:"link"`
	Description     string       `json:"description"`
	GUID            string       `json:"guid"`
	EnclosureURL    string       `json:"enclosureUrl"`
	EnclosureType   string       `json:"enclosureType"`
	Image           string       `json:"image"`
	FeedImage       string       `json:"feedImage"`
	FeedURL         string       `json:"feedUrl"`
	FeedAuthor      string       `json:"feedAuthor"`
	FeedTitle       string       `json:"feedTitle"`
	FeedLanguage    string       `json:"feedLanguage"`
	Transcripts     []Transcript `json:"transcripts,omitempty"`
	DatePublished   int64        `json:"datePublished"`
	DateCrawled     int64        `json:"dateCrawled"`
	EpisodeType     *string      `json:"episodeType"`
	ChaptersURL     *string      `json:"chaptersUrl"`
	TranscriptURL   *string      `json:"transcriptUrl"`
	Duration        *int         `json:"duration"`
	Episode         *int         `json:"episode"`
	Season          *int         `json:"season"`
	FeedItunesID    *int         `json:"feedItunesId"`
	ID              int          `json:"id"`
	EnclosureLength int          `json:"enclosureLength"`
	FeedID          int          `json:"feedId"`
	Explicit        int          `json:"explicit"`
}

func (e Episode) TableHeaders() []string {
	return []string{"ID", "TITLE", "FEED", "PUBLISHED", "DURATION", "TYPE"}
}

func (e Episode) TableRow() []string {
	published := ""
	if e.DatePublished > 0 {
		published = time.Unix(e.DatePublished, 0).Format("2006-01-02")
	}
	epType := ""
	if e.EpisodeType != nil {
		epType = *e.EpisodeType
	}
	return []string{
		strconv.Itoa(e.ID),
		e.Title,
		e.FeedTitle,
		published,
		formatDuration(e.Duration),
		epType,
	}
}

type EpisodesByFeedResponse struct {
	APIResponse
	Items []Episode         `json:"items"`
	Query map[string]string `json:"query,omitempty"`
}

type EpisodeByIDResponse struct {
	APIResponse
	Episode Episode           `json:"episode"`
	Query   map[string]string `json:"query,omitempty"`
}

type EpisodesLiveResponse struct {
	APIResponse
	Items []Episode `json:"items"`
}

type EpisodesRandomResponse struct {
	APIResponse
	Episodes []Episode `json:"episodes"`
}
