package database

import (
	"fmt"
	"github.com/Somvaded/event-scheduler/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Connect DB
	db, err := gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		fmt.Errorf("failed to connect database: %v", err)
		return nil,err	
	}

	// Migrate schema
	err= db.AutoMigrate(&models.User{},&models.Event{})
	if err != nil {
		fmt.Errorf("failed to migrate database: %v", err)
		return nil,err
	}
	return db,nil
}