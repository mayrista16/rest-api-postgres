package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mayrista16/rest-api-postgres/controllers/user_controller"
)

func UserRoute(app *gin.Engine) {
	route := app

	//Route User
	route.GET("/users", user_controller.GetAllUser)
	route.POST("/user", user_controller.Store)
	route.GET("/user/:id", user_controller.GetById)
	route.PATCH("/user/:id", user_controller.Update)
	route.DELETE("/user/:id", user_controller.Delete)

}
