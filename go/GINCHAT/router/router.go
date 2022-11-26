package router

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/listUser", service.ListUsers)
	r.GET("/createDb", service.CreateDB)
	r.GET("/createUser", service.CreateUser)
	return r
}
