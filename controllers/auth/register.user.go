package controllers

import (
	"main/database"

	// "main/models"
	"main/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "github.com/thoas/go-funk"
)

type Register struct {
	Username string `json:"username"`
}

// RegisterUser is a controller function to register a user
func RegisterUser(c *gin.Context) {

	body := models.User{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"msg": "Invalid JSON",
		})
		return
	}

	if body.Username == "" || body.Password == "" || body.Email == "" {
		c.JSON(400, gin.H{
			"msg":      "Missing required fields",
			"expected": "username, password, email",
		})
		return
	}

	db := database.DB
	var result models.User
	var user models.User

	db.Find(&user, "username", body.Username).Scan(&result)

	if result.Username != "" {
		c.JSON(400, gin.H{
			"msg":     "Username already exists",
			"success": "failed",
		})
		return
	}

	// set user to its type
	user.ID = uuid.New()
	user.Username = body.Username
	user.Password = body.Password
	user.Email = body.Email

	db.Create(&user)

	c.JSON(200, gin.H{
		"msg":     "user created successfully",
		"success": "success",
	})
	return

}
