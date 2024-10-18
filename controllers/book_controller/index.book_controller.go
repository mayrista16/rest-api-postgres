package book_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mayrista16/rest-api-postgres/database"
	"github.com/mayrista16/rest-api-postgres/models"
)

func GetAllBook(ctx *gin.Context) {
	books := new([]models.User)
	err := database.DB.Table("books").Find(&books).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": books,
	})
}
