package routes

import (
	controller "go-lang-jwt/controller"
	"go-lang-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	/*
		In UserRoutes to access API we will be using middleware due to authentication
		You can observe in authRoute where we have login and Signup routes that are open but after login
		authentication token is generated which is used fo securing routing
	*/
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}
