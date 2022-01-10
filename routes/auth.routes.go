package routes

import (
	"github.com/gin-gonic/gin"
	auth "main/controllers/auth"
)

func authRoutes(router *gin.Engine) {
	router.POST("/auth/register", auth.RegisterUser)
	router.POST("/auth/login", auth.LoginUser)
}
