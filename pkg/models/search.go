package models

import (
	"strconv"
	"time"
)

type PersonResult struct {
	Title           string `json:"title"`
	Link            string `json:"link"`
	Description     string `json:"description"`
	GUID            string `json:"guid"`
	EnclosureURL    string `json:"enclosureUrl"`
	EnclosureType   string `json:"enclosureType"`
	EpisodeType     string `json:"episodeType"`
	Image           string `json:"image"`
	FeedImage       string `json:"feedImage"`
	FeedURL         string `json:"feedUrl"`
	FeedAuthor      string `json:"feedAuthor"`
	FeedTitle       string `json:"feedTitle"`
	FeedLanguage    string `json:"feedLanguage"`
	ChaptersURL     string `json:"chaptersUrl"`
	TranscriptURL   string `json:"transcriptUrl"`
	DatePublished   int64  `json:"datePublished"`
	DateCrawled     int64  `json:"dateCrawled"`
	Duration        *int   `json:"duration"`
	Episode         *int   `json:"episode"`
	Season          *int   `json:"season"`
	FeedItunesID    *int   `json:"feedItunesId"`
	ID              int    `json:"id"`
	EnclosureLength int    `json:"enclosureLength"`
	FeedID          int    `json:"feedId"`
	Explicit        int    `json:"explicit"`
}

func (p PersonResult) TableHeaders() []string {
	return []string{"ID", "TITLE", "FEED", "PUBLISHED", "DURATION"}
}

func (p PersonResult) TableRow() []string {
	published := ""
	if p.DatePublished > 0 {
		published = time.Unix(p.DatePublished, 0).Format("2006-01-02")
	}
	return []string{
		strconv.Itoa(p.ID),
		p.Title,
		p.FeedTitle,
		published,
		formatDuration(p.Duration),
	}
}

type SearchByPersonResponse struct {
	APIResponse
	Items []PersonResult `json:"items"`
	Query string         `json:"query"`
}
