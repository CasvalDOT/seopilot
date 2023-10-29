package models

import (
	"time"

	"gorm.io/gorm"
)

type Site struct {
	ID         int
	URL        string
	Status     string
	Source     Source
	SourceID   int
	Alert      Alert
	AlertID    *int
	Attributes []SiteAttribute
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type SiteCreateRequest struct {
	URL     string `json:"url" binding:"required,url,max=255"`
	AlertID *int   `json:"alert_id" binding:"omitempty,number"`
}

type SiteUpdateRequest struct {
	URL     string `json:"url" binding:"required,url,max=255"`
	AlertID *int   `json:"alert_id" binding:"omitempty,number"`
}

type SiteViewResponse struct {
	ID         int                         `json:"id"`
	URL        string                      `json:"url"`
	Status     string                      `json:"status"`
	Alert      AlertViewResponse           `json:"alert"`
	Attributes []SiteAttributeViewResponse `json:"attributes"`
	Source     SourceViewResponse          `json:"source"`
	CreatedAt  time.Time                   `json:"created_at"`
	UpdatedAt  time.Time                   `json:"updated_at"`
}

type SiteSimpleViewResponse struct {
	ID        int                      `json:"id"`
	URL       string                   `json:"url"`
	Status    string                   `json:"status"`
	Alert     AlertViewSimpleResponse  `json:"alert"`
	Source    SourceViewSimpleResponse `json:"source"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

type SiteViewAnyResponse struct {
	Data []SiteSimpleViewResponse `json:"data"`
}
