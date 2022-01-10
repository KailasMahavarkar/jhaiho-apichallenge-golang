package controllers

import (
	"main/database"
	"main/models"

	"github.com/gin-gonic/gin"
)

func GetSessions() gin.HandlerFunc {
	db := database.DB
	var sessions []models.Session
	db.Find(&sessions)

	return gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "All Session",
			"result": sessions,
			"success": "success",
		})
	})

}
