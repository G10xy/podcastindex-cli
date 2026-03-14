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
	EpisodeType     *string      `json:"episodeType"`
	FeedItunesID    *int         `json:"feedItunesId"`
	Season          *int         `json:"season"`
	Episode         *int         `json:"episode"`
	Duration        *int         `json:"duration"`
	TranscriptURL   *string      `json:"transcriptUrl"`
	ChaptersURL     *string      `json:"chaptersUrl"`
	Image           string       `json:"image"`
	Description     string       `json:"description"`
	FeedAuthor      string       `json:"feedAuthor"`
	FeedTitle       string       `json:"feedTitle"`
	FeedLanguage    string       `json:"feedLanguage"`
	Link            string       `json:"link"`
	FeedURL         string       `json:"feedUrl"`
	GUID            string       `json:"guid"`
	FeedImage       string       `json:"feedImage"`
	Title           string       `json:"title"`
	EnclosureType   string       `json:"enclosureType"`
	EnclosureURL    string       `json:"enclosureUrl"`
	Transcripts     []Transcript `json:"transcripts,omitempty"`
	DateCrawled     int64        `json:"dateCrawled"`
	DatePublished   int64        `json:"datePublished"`
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
	Query map[string]string `json:"query,omitempty"`
	APIResponse
	Items []Episode `json:"items"`
}

type EpisodeByIDResponse struct {
	Query map[string]string `json:"query,omitempty"`
	APIResponse
	Episode Episode `json:"episode"`
}

type EpisodesLiveResponse struct {
	APIResponse
	Items []Episode `json:"items"`
}

type EpisodesRandomResponse struct {
	APIResponse
	Episodes []Episode `json:"episodes"`
}
