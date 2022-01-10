package controllers

import (
	"main/database"

	"github.com/gin-gonic/gin"
	"main/models"
)

// find dataBlock by ID
func GetUsers() gin.HandlerFunc {

	db := database.DB
	var users []models.User
	db.Find(&users)

	return gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":    "Wasup Kai",
			"result": users,
		})
	})

}
