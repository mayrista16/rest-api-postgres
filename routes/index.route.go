package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mayrista16/rest-api-postgres/controllers/user_controller"
	"github.com/mayrista16/rest-api-postgres/middleware"
)

func InitRoute(app *gin.Engine) {
	route := app

	UserRoute(route)

	//Route Book
	route.POST("/signup", user_controller.Signup)
	route.POST("/login", user_controller.Login)
	route.GET("/validate", middleware.RequireAuth, user_controller.Validate)
}
