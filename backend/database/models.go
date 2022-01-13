package database

import (
	"time"
)

type Post struct {
	Date      time.Time `json:"Date" pg:"default:now(), pk, notnull"`
	Content   string    `json:"Content" pg:", notnull"`
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
}
