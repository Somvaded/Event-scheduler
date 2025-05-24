package models

import (
	"time"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
	Reminded    bool      `json:"reminded"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"-"` // omit from JSON but preload when needed
}
