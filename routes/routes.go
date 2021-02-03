package routes

import (
	controllers "pelipper/controllers"

	"github.com/gin-gonic/gin"
)

/*
Routes -> resgister backend routes in the given router
*/
func Routes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)

	router.POST("/emails/users/verify", controllers.EmailUserVerification)
}
