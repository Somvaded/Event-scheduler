package utils

import (
	"fmt"
	"time"
	"github.com/Somvaded/event-scheduler/models"
	"gorm.io/gorm"
)

func CheckUpcomingEvents(db *gorm.DB) {

	var events []models.Event
	now := time.Now()
	soon := now.Add(24 * time.Hour)
	// Only get events not already reminded
	db.Preload("User").Where("time BETWEEN ? AND ? AND reminded = false", now, soon).Find(&events)
	
	for _, e := range events {
		subject := "Reminder: " + e.Title
		body := fmt.Sprintf("Hi! Just a reminder: your event '%s' starts at %s.\n\nDetails: %s",
			e.Title,
			e.Time.Format("15:04 on 02 Jan 2006"),
			e.Description,
		)
		SendEmail(e.User.Email, subject, body)
		// Mark as reminded
		e.Reminded = true
		db.Save(&e)
	}
}