package controllers

import (
	// "encoding/json"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm/clause"

	"main/database"
	"main/models"
)

// update user comment controller
func ListComment(c *gin.Context) {

	// get all comments
	db := database.DB
	var comments []models.Comment

	// find all comments order by id desc
	db.Raw("SELECT * FROM comments ORDER BY created_at DESC").Scan(&comments)

	// send comments
	c.JSON(200, gin.H{
		"msg":     "comment updated successfully",
		"success": "success",
		"result":  comments,
	})

}
