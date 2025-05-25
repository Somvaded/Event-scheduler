package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Somvaded/event-scheduler/database"
	"github.com/Somvaded/event-scheduler/routes"
	"github.com/Somvaded/event-scheduler/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

)
var webPort = "8080"
func main(){

	r := gin.Default()

	// Setup database

	db , err := database.InitDB()
	if err != nil {
		fmt.Errorf("database initialization failed: %v", err)
		return
	}
	// Setup routes
	routes.EventRoutes(r, db)

	//Load Port
	err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	if os.Getenv("PORT") !=""{
		webPort = os.Getenv("PORT")
	}
	//go routine to check for upcoming events
	go func() {
		for {
			utils.CheckUpcomingEvents(db)
			time.Sleep(1 * time.Minute)
		}
	}()

	//run the ser ver
	r.Run(fmt.Sprintf(":%s", webPort))
}

