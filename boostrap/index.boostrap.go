package boostrap

import (
	"github.com/gin-gonic/gin"
	"github.com/mayrista16/rest-api-postgres/configs/app_config"
	"github.com/mayrista16/rest-api-postgres/database"
	"github.com/mayrista16/rest-api-postgres/routes"
)

func BoostrapApp() {
	database.ConnectDatabase()
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
