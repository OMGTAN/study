package service

import (
	"ginchat/models"
	"ginchat/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDB(c *gin.Context) {
	utils.DB.AutoMigrate(&models.UserBasic{})
	c.JSON(http.StatusOK, gin.H{
		"message": "create db success",
	})
}

func CreateUser(c *gin.Context) {
	user := &models.UserBasic{}
	user.Name = "h1"
	utils.DB.Create(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "create user success",
	})
}

func ListUsers(c *gin.Context) {
	data := make([]models.UserBasic, 10)
	utils.DB.Find(&data)
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
