package controllers

// 5d4ee9e5-b48c-4b6d-8c06-ec5d4313441b

import (
	// "fmt"
	_ "fmt"
	"main/database"
	"main/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thoas/go-funk"
)

type Comment struct {
	ID string `json:"id"`
}

// checks if uuid valid or not
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func ViewComment(c *gin.Context) {

	// set body of type Comment
	body := Comment{}

	// bind json to body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"msg":     "Undefined comment UUID",
			"success": "failed",
		})
		return
	}

	// check if comment id is valid
	if !IsValidUUID(body.ID) {
		c.JSON(400, gin.H{
			"msg":     "Invalid comment UUID",
			"success": "failed",
		})
		return
	}

	var result models.Comment
	var comment models.Comment

	// get db
	db := database.DB

	// find comment by id and store in result
	db.Limit(1).Where(&Comment{
		ID: body.ID,
	}).Find(&comment).Scan(&result)

	// check if comment is found if not return error
	if funk.IsEmpty(result.ID) {
		c.JSON(400, gin.H{
			"msg":     "comment not found",
			"success": "failed",
		})
		return
	}

	// send comment
	c.JSON(200, gin.H{
		"msg":     "comment read successfully",
		"success": "success",
		"result":  result,
	})
}
