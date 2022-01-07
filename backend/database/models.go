package database

import (
	"time"
)

type Post struct {
	Date    time.Time `json:"Date" gorm:"primaryKey"`
	Content string    `json:"content"`
}
