package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thoas/go-funk"

	"main/database"
	"main/helpers"
	"main/models"
)

func AuthSession(c *gin.Context) {

	ssid := c.GetHeader("ssid")

	// check if ssid is empty
	if funk.IsEmpty(ssid) {
		c.Abort()
		c.JSON(400, gin.H{
			"msg":     "Missing ssid header",
			"success": "failed",
		})
		return

	} else if !helpers.IsValidUUID(ssid) {
		// check if ssid is valid
		c.Abort()
		c.JSON(400, gin.H{
			"msg":     "Invalid ssid header",
			"success": "failed",
		})
		return
	}

	// check if session exists
	var session models.Session
	var session_result models.Session
	parsed_uuid, _ := uuid.Parse(ssid)

	db := database.DB

	// check if session exists and store in session_result
	db.Where(&models.Session{
		ID: parsed_uuid,
	}).Find(&session).Scan(&session_result)

	// check if session is valid or not
	sessionTimeElapsed := time.Now().Unix() - session_result.CreatedAt.Unix()

	// check time elapsed
	if sessionTimeElapsed > 900 {
		c.Abort()
		c.JSON(400, gin.H{
			"msg":     "Session expired",
			"success": "failed",
		})
		return
	} else {
		c.Next()
	}

}
