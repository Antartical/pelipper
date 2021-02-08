package routes

import (
	"pelipper/controllers"
	"pelipper/services"

	"github.com/gin-gonic/gin"
)

/*
Routes -> resgister backend routes in the given router
*/
func Routes(router *gin.Engine) {
	controllers.RegisterPingRoutes(router)
	controllers.RegisterUserRoutes(router, services.NewSMTPEmailService())
}
