package database

import (
	"time"

	"github.com/go-pg/pg/v10/types"
)

type Post struct {
	Date      time.Time      `json:"Date" pg:"default:now(), pk, notnull"`
	Content   string         `json:"Content" pg:", notnull"`
	UpdatedAt types.NullTime `json:"UpdatedAt,omitempty"`
}
