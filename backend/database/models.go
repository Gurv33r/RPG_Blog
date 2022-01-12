package database

import (
	"time"
)

type Post struct {
	Date      time.Time `json:"date" pg:"default:now(),pk,notnull"`
	Content   string    `json:"content" pg:",notnull"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
