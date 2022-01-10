package routes

import (
	"github.com/gin-gonic/gin"
	public "main/controllers/public"
)

func publicRoutes(router *gin.Engine) {

	// create comment
	router.POST("/create", public.CreateComment())

}
