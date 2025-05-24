package controllers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/Somvaded/event-scheduler/models"
)

func NewEventController(db *gorm.DB) *EventController {
	return &EventController{db: db}
}

type EventController struct {
	db *gorm.DB
}


func (e *EventController)CreateEvent(c *gin.Context) {
	var event models.Event
	userID := c.MustGet("userID").(uint)
	
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if event.Time.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Event time must be in the future"})
		return
	}
	event.UserID = userID
	e.db.Create(&event)
	c.JSON(http.StatusCreated, event)
}

func (e *EventController) GetEvents(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var events []models.Event
	e.db.Where("user_id = ?", userID).Find(&events)
	c.JSON(http.StatusOK, events)
}

func (e *EventController) GetEvent(c *gin.Context) {

	var event models.Event
	if err := e.db.First(&event, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func (e *EventController)UpdateEvent(c *gin.Context) {
	var event models.Event
	if err := e.db.First(&event, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var input models.Event
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.Title = input.Title
	event.Description = input.Description
	event.Time = input.Time
	event.Reminded = input.Reminded
	e.db.Save(&event)

	c.JSON(http.StatusOK, event)
}

func (e *EventController) DeleteEvent(c *gin.Context) {
	var event models.Event
	if err := e.db.First(&event, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	e.db.Delete(&event)
	c.Status(http.StatusNoContent)
}
