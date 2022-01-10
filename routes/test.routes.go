package routes

import (
	"github.com/gin-gonic/gin"
	_ "main/controllers/test"
)

func testRoutes(router *gin.Engine) {
	router.GET("/test", func(c *gin.Context) {

		c.Set("name", "kai")
		c.Next()
	}, func(c *gin.Context) {
		c.Next()
	}, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
			"result":  c.MustGet("name"),
		})
	},
	)
}
