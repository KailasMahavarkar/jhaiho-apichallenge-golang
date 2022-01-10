package routes

import (
	private "main/controllers/private"

	"github.com/gin-gonic/gin"

	// "main/helpers"
	mw "main/middlewares"
)

func privateRoutes(router *gin.Engine) {

	// view comment
	router.POST("/private/view", mw.AuthSession, private.ViewComment)

	// // update comment
	router.POST("/private/update", mw.AuthSession, private.UpdateComment)

	// // list comment
	router.GET("/private/list", mw.AuthSession, private.ListComment)

}
