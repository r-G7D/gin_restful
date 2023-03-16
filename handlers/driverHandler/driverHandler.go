package driverHandler

import (
	"net/http"
	"r-G7D/go_gin_restful/app"
	"r-G7D/go_gin_restful/domains"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var drivers []domains.Driver

	app.DB.Find(&drivers)
	c.JSON(http.StatusOK, gin.H{"drivers": drivers})
}

func Show(c *gin.Context) {
	var driver domains.Driver
	id := c.Param("id")

	if err := app.DB.First(&driver, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"driver": driver})
}

func Create(c *gin.Context) {
	var driver domains.Driver
	if err := c.BindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
		return
	}

	var existingDriver domains.Driver
	if err := app.DB.Where("name == ? and email == ?", driver.Name, driver.Email).First(&existingDriver).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		}
	}

	if existingDriver.Id != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Record already exists!"})
		return
	}

	if err := app.DB.Create(&driver).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"driver": driver})
}

func Update(c *gin.Context) {
	var driver domains.Driver
	id := c.Param("id")

	if err := c.BindJSON(&driver); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
		return
	}

	if err := app.DB.Where("id = ?", id).Updates(&driver).Error; err != nil {
		// if existingDriver := app.DB.Where("name == ? and email == ?", driver.Name, driver.Email); existingDriver != nil {
		// 	c.JSON(http.StatusConflict, gin.H{"error": "Record already exists!"})
		// 	return
		// }
		//!doesn't work
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"driver": driver, "message": "Record updated successfully!"})
}

func Delete(c *gin.Context) {
	var driver domains.Driver
	id := c.Param("id")

	if err := app.DB.First(&driver, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
			return
		}
	}

	if err := app.DB.Delete(&driver).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Record deleted successfully!",
	})
}
