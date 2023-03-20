package authHandler

import (
	"net/http"
	"r-G7D/go_gin_restful/app"
	"r-G7D/go_gin_restful/domains"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

func RegisterHandler(c *gin.Context) {
	var driver domains.RegisterResponse
	if err := c.Bind(&driver); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
		return
	}

	if err := validate.Struct(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data!"})
		return
	}

	var existingDriver domains.RegisterResponse
	// if err := app.DB.Where("email = ?", driver.Email).First(&existingDriver).Error; err != nil {
	// 	if err != gorm.ErrRecordNotFound {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
	// 	}
	// }

	if existingDriver.Id != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists!"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(driver.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		return
	}

	driver.Password = string(hashedPassword)

	if err := app.DB.Create(&driver).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func LoginHandler(c *gin.Context) {
	var login domains.LoginResponse
	if err := c.Bind(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
		return
	}

	if err := validate.Struct(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data!"})
		return
	}

	var existingDriver domains.RegisterResponse
	if err := app.DB.Where("email = ?", login.Email).First(&existingDriver).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		}
	}

	if existingDriver.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingDriver.Password), []byte(login.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials!"})
		return
	}

	// TODO: add jwt

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}

func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful!"})
}
