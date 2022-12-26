package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	// v1 := r.Group("/")
	// {
	// 	eg := v1.Group("/test")
	// 	{
	// 		eg.GET("/listUsers", service.ListUsers)
	// 	}
	// }
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/index", service.GetIndex)
	r.GET("/createDb", service.CreateDB)
	r.POST("/createUser", service.CreateUser)
	r.POST("/updateUser", service.UpdateUser)
	r.GET("/deleteUser", service.DeleteUser)
	r.GET("/listUsers", service.ListUsers)

	r.GET("/sendMsg", service.SendMsg)
	return r
}
