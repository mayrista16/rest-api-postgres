package user_controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mayrista16/rest-api-postgres/database"
	"github.com/mayrista16/rest-api-postgres/models"
	"github.com/mayrista16/rest-api-postgres/requests"
	"github.com/mayrista16/rest-api-postgres/responses"
	"golang.org/x/crypto/bcrypt"
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
		"message": "data received.",
		"data":    users,
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

	if user.ID == "" {
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

func Signup(c *gin.Context) {
	//Get form data email dan password
	var body struct {
		ID       string
		Name     string
		Address  string
		Date     time.Time
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed get form data",
		})
		return
	}

	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate password",
		})
		return
	}

	//Create User
	genId := uuid.New().String()

	user := models.User{ID: genId, Name: body.Name, Address: body.Address, Date: body.Date, Password: string(hash), Email: body.Email}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "data has been created",
	})
}

func Login(c *gin.Context) {
	// Get email and password form data
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get form data",
		})

		return
	}

	// Check requested User
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data not found",
		})

		return
	}

	// Compare password with pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email / Password",
		})

		return
	}

	// Gen jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := []byte("4lly0uRth1n6S8eLoN6T0u5")
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}
	// Send back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
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

	user.ID = genId
	user.Name = userReq.Name
	user.Address = userReq.Address
	user.Date = userReq.BornDate

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

	if user.ID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found in database",
		})
		return
	}

	user.Name = userReq.Name
	user.Address = userReq.Address
	user.Date = userReq.BornDate

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
