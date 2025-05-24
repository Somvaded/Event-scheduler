package routes

import (
	"github.com/Somvaded/event-scheduler/controller"
	"github.com/Somvaded/event-scheduler/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EventRoutes(r *gin.Engine, db *gorm.DB) {

	eventController := controllers.NewEventController(db)
	userController := controllers.NewUserController(db)
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	

	logged := r.Group("/user").Use(middlewares.AuthMiddleware())

	logged.POST("/events", eventController.CreateEvent)
	logged.GET("/events", eventController.GetEvents)
	logged.GET("/events/:id", eventController.GetEvent)
	logged.PUT("/events/:id", eventController.UpdateEvent)
	logged.DELETE("/events/:id", eventController.DeleteEvent)
}
