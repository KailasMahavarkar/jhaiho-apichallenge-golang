package controllers

import (
	// local imports
	"fmt"
	"main/database"
	"main/models"

	// import go-funk
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"

	// import uuid
	"github.com/google/uuid"

	// import time
	"time"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUser is a controller function to login a user
func LoginUser(c *gin.Context) {

	body := Login{}

	// bind json
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"msg": "Invalid JSON",
		})
		return
	}

	// check if username and password are empty
	if body.Username == "" || body.Password == "" {
		c.JSON(400, gin.H{
			"msg":     "username or password is missing",
			"success": "failed",
		})
		return
	}

	// get db
	db := database.DB
	var user models.User

	// find user
	db.Where(&Login{
		Username: body.Username,
		Password: body.Password,
	}).Find(&user)

	// check if user is found
	if funk.IsEmpty(user) {
		c.JSON(400, gin.H{
			"msg":     "username or password does not match",
			"success": "invalid",
		})
		return
	}

	// this stores result of the query
	var result models.Session

	// this stores type session
	var session models.Session

	// query db
	db.Limit(1).Find(&session, "username = ?", body.Username).Scan(&result)

	// check if session is found
	if funk.IsEmpty(result.Username) {
		// creates a new session and send reponse
		var session models.Session
		session.ID = uuid.New()
		session.Username = body.Username

		db := database.DB
		db.Create(&session)


		c.JSON(200, gin.H{
			"msg":     "user logged in successfully",
			"success": "success",
			"ssid":    session.ID,
		})
		return
	}

	sessionTimeElapsed := time.Now().Unix() - result.CreatedAt.Unix()

	if sessionTimeElapsed > 900 {
		// 	// creates a new session and send reponse
		fmt.Println("session is expired --> creating a new one")
		db.First(&result)
		result.CreatedAt = time.Now()
		db.Save(&result)
		c.JSON(200, gin.H{
			"msg":     "updated session & logged in successfully",
			"success": "success",
			"ssid":    session.ID,
		})
		return
	} else {
		c.JSON(400, gin.H{
			"msg":     "username already logged in",
			"success": "invalid",
		})
		return
	}

}
