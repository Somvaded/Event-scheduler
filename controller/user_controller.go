package controllers

import (
	"fmt"
	"net/http"

	"github.com/Somvaded/event-scheduler/models"
	"github.com/Somvaded/event-scheduler/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}
type UserController struct {
	db *gorm.DB
}


func (u *UserController) Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "Email and password are required"})
		return
	}
	hashed_password , err := utils.HashPassword(user.Password) 
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = hashed_password
	if err := u.db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user" + err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": fmt.Sprintf("User %d created successfully", user.ID)})
}

func (u *UserController) Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := u.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if check := utils.CheckPasswordHash(input.Password, user.Password); !check {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
