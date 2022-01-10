package routes

import "github.com/gin-gonic/gin"

// entire routes of application are defined here
func AllRoutes(router *gin.Engine) {
	// public routes
	publicRoutes(router)

	// private routes
	privateRoutes(router)

	// auth routes
	authRoutes(router)

	// test routes
	testRoutes(router)
}

func Run(PORT string) {
	router := gin.Default()
	AllRoutes(router)
	router.Run(":" + PORT)
}
