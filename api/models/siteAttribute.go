package models

import "time"

type SiteAttribute struct {
	ID        int
	Name      string
	Value     string
	SiteID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SiteAttributeViewResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
