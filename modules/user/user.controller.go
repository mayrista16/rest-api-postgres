package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mayrista16/rest-api-postgres/database"
	"github.com/mayrista16/rest-api-postgres/models"
	"github.com/mayrista16/rest-api-postgres/requests"
	"github.com/mayrista16/rest-api-postgres/responses"
)

func GetAllUser(ctx *gin.Context) {
	users := new([]models.User)
	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(responses.UserResponse)

	err := database.DB.Table("users").Where("id = ?", id).Find(&user).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found in database",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data transmitted.",
		"data":    user,
	})
}

func Store(ctx *gin.Context) {
	userReq := new(requests.UserRequest)

	if err := ctx.ShouldBind(&userReq); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	genId := uuid.New().String()

	user := new(models.User)

	user.ID = &genId
	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Date = &userReq.BornDate

	err := database.DB.Table("users").Create(&user).Error
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "data failed to store.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data has been stored.",
		"data":    user,
	})
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	user := new(models.User)
	userReq := new(requests.UserRequest)

	if errReq := ctx.ShouldBind(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error.",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found in database",
		})
		return
	}

	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Date = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "update data failed.",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "data updated succesfully.",
		"data":    user,
	})
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&models.User{})

	ctx.JSON(200, gin.H{
		"message": "data has been deleted",
	})
}
