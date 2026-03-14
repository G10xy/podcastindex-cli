package models

import "strconv"

type Stats struct {
	FeedCountTotal       int `json:"feedCountTotal"`
	EpisodeCountTotal    int `json:"episodeCountTotal"`
	FeedsWithNewEpisodes int `json:"feedsWithNewEpisodes"`
	FeedsWithValueBlocks int `json:"feedsWithValueBlocks"`
}

func (s Stats) TableHeaders() []string {
	return []string{"FEEDS", "EPISODES", "NEW EPISODES", "VALUE FEEDS"}
}

func (s Stats) TableRow() []string {
	return []string{
		strconv.Itoa(s.FeedCountTotal),
		strconv.Itoa(s.EpisodeCountTotal),
		strconv.Itoa(s.FeedsWithNewEpisodes),
		strconv.Itoa(s.FeedsWithValueBlocks),
	}
}

type StatsResponse struct {
	APIResponse
	Stats Stats `json:"stats"`
}
