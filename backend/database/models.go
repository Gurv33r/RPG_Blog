package database

import (
	"time"
)

type Post struct {
	Date      string    `json:"date" gorm:"primaryKey"`
	CreatedAt time.Time `json:"timeCreatedAt"`
	Content   string    `json:"content"`
}
