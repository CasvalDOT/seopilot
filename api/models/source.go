package models

import "time"

type Source struct {
	ID        int
	Name      string
	Code      string
	Token     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type SourceViewResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type SourceViewSimpleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

const SourceCodeManual = "manual"
