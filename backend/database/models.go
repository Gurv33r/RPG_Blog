package database

import (
	"time"

	"github.com/go-pg/pg/v10"
)

type Post struct {
	Date      time.Time   `json:"Date" pg:"default:now(), pk, notnull"`
	Content   string      `json:"Content" pg:", notnull"`
	UpdatedAt pg.NullTime `json:"UpdatedAt,omitempty"`
}
