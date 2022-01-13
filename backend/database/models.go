package database

import (
	"database/sql"
	"time"
)

type Post struct {
	Date      time.Time    `json:"Date" pg:"default:now(), pk, notnull"`
	Content   string       `json:"Content" pg:", notnull"`
	UpdatedAt sql.NullTime `json:"UpdatedAt,omitempty"`
}
