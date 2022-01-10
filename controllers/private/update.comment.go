package controllers

import (
	"io/ioutil"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"

	"main/database"
	"main/models"
)

// update user comment controller
func UpdateComment(c *gin.Context) {

	unparsedBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "Invalid JSON body",
		})
		return
	}

	// parse json body from request
	jsonParsed, _ := gabs.ParseJSON(unparsedBody)
	email, email_error := jsonParsed.Path("email").Data().(string)
	text, text_error := jsonParsed.Path("text").Data().(string)
	username, username_error := jsonParsed.Path("username").Data().(string)
	id, id_error := jsonParsed.Path("id").Data().(string)

	// email error
	if !email_error {
		c.JSON(400, gin.H{
			"msg":      "Missing required fields",
			"expected": "email",
		})
		return
	}

	// text error
	if !text_error {
		c.JSON(400, gin.H{
			"msg":      "Missing required fields",
			"expected": "text",
		})
		return
	}

	// username error
	if !username_error {
		c.JSON(400, gin.H{
			"msg":      "Missing required fields",
			"expected": "username",
		})
		return
	}

	// id error
	if !id_error {
		c.JSON(400, gin.H{
			"msg":      "Missing required fields",
			"expected": "id",
		})
		return
	}

	// update comment 
	var comment models.Comment
	newMap := make(map[string]interface{})
	newMap["text"] = text
	newMap["username"] = username
	newMap["email"] = email

	// update comment in database
	db := database.DB
	db.Model(&comment).Where("id = ?", id).Updates(newMap)

	c.JSON(200, gin.H{
		"msg":     "comment updated successfully",
		"success": "success",
	})

}
