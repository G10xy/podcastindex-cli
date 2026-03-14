package models

import (
	"strconv"
)

type ValueDestination struct {
	CustomKey   *string `json:"customKey,omitempty"`
	CustomValue *string `json:"customValue,omitempty"`
	Fee         *bool   `json:"fee,omitempty"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Type        string  `json:"type"`
	Split       int     `json:"split"`
}

func (v ValueDestination) TableHeaders() []string {
	return []string{"NAME", "ADDRESS", "TYPE", "SPLIT", "FEE"}
}

func (v ValueDestination) TableRow() []string {
	fee := ""
	if v.Fee != nil && *v.Fee {
		fee = "yes"
	}
	return []string{
		v.Name,
		v.Address,
		v.Type,
		strconv.Itoa(v.Split),
		fee,
	}
}

type ValueModel struct {
	Suggested *string `json:"suggested,omitempty"`
	Type      string  `json:"type"`
	Method    string  `json:"method"`
}

type Value struct {
	Model        ValueModel         `json:"model"`
	Destinations []ValueDestination `json:"destinations"`
}

type ValueByFeedResponse struct {
	Query map[string]string `json:"query,omitempty"`
	Value Value             `json:"value"`
	APIResponse
}

type ValueByEpisodeGUIDResponse struct {
	APIResponse
	Value Value `json:"value"`
}

type ValueBatchByEpisodeGUIDResponse struct {
	Items map[string]interface{} `json:"items"`
	APIResponse
}
