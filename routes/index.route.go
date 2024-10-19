package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mayrista16/rest-api-postgres/controllers/book_controller"
)

func InitRoute(app *gin.Engine) {
	route := app

	UserRoute(route)

	//Route Book
	route.GET("/", book_controller.GetAllBook)
}
