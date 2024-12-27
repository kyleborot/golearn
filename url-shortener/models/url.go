package models

import "time"

type UrlStruct struct {
	ID        string    `json:"id"`
	LongURL   string    `json:"long_url"`
	Timestamp time.Time `json:"timestamp"`
}
