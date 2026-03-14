package models

import (
	"strconv"
	"time"
)

type PersonResult struct {
	Duration        *int   `json:"duration"`
	FeedItunesID    *int   `json:"feedItunesId"`
	Season          *int   `json:"season"`
	Episode         *int   `json:"episode"`
	EpisodeType     string `json:"episodeType"`
	Link            string `json:"link"`
	Title           string `json:"title"`
	Image           string `json:"image"`
	FeedImage       string `json:"feedImage"`
	FeedURL         string `json:"feedUrl"`
	FeedAuthor      string `json:"feedAuthor"`
	FeedTitle       string `json:"feedTitle"`
	FeedLanguage    string `json:"feedLanguage"`
	ChaptersURL     string `json:"chaptersUrl"`
	TranscriptURL   string `json:"transcriptUrl"`
	EnclosureType   string `json:"enclosureType"`
	Description     string `json:"description"`
	EnclosureURL    string `json:"enclosureUrl"`
	GUID            string `json:"guid"`
	DateCrawled     int64  `json:"dateCrawled"`
	DatePublished   int64  `json:"datePublished"`
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
	Query string `json:"query"`
	APIResponse
	Items []PersonResult `json:"items"`
}
