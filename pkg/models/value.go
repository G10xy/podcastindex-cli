package models

import (
	"strconv"
)

type ValueDestination struct {
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Type        string  `json:"type"`
	CustomKey   *string `json:"customKey,omitempty"`
	CustomValue *string `json:"customValue,omitempty"`
	Split       int     `json:"split"`
	Fee         *bool   `json:"fee,omitempty"`
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
	Type      string  `json:"type"`
	Method    string  `json:"method"`
	Suggested *string `json:"suggested,omitempty"`
}

type Value struct {
	Model        ValueModel         `json:"model"`
	Destinations []ValueDestination `json:"destinations"`
}

type ValueByFeedResponse struct {
	APIResponse
	Value Value             `json:"value"`
	Query map[string]string `json:"query,omitempty"`
}

type ValueByEpisodeGUIDResponse struct {
	APIResponse
	Value Value `json:"value"`
}

type ValueBatchByEpisodeGUIDResponse struct {
	APIResponse
	Items map[string]interface{} `json:"items"`
}
