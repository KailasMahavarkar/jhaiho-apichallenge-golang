package controllers

import (
	"main/database"
	"main/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateComment() gin.HandlerFunc {

	return gin.HandlerFunc(func(c *gin.Context) {
		body := models.Comment{}
		db := database.DB

		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{
				"msg": "Invalid JSON",
			})
			return
		}

		for _, x := range []string{body.Username, body.Text, body.Email} {
			if x == "" {
				c.JSON(400, gin.H{
					"msg":      "Missing required fields",
					"expected": "expected " + x,
				})
				return
			}
		}

		var comment models.Comment
		comment.ID = uuid.New()
		comment.Username = body.Username
		comment.Email = body.Email
		comment.Text = body.Text
		db.Create(&comment)

		c.JSON(200, gin.H{
			"msg":        "comment created successfully",
			"success":    "success",
			"comment-id": comment.ID,
		})
	})

}
