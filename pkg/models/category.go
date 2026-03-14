package models

import "strconv"

type Category struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (c Category) TableHeaders() []string {
	return []string{"ID", "NAME"}
}

func (c Category) TableRow() []string {
	return []string{
		strconv.Itoa(c.ID),
		c.Name,
	}
}

type CategoriesListResponse struct {
	APIResponse
	Feeds []Category `json:"feeds"`
}
