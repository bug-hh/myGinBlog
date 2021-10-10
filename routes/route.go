package routes

import (
	"github.com/gin-gonic/gin"
	v1 "myGinBlog/api/v1"
	"myGinBlog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
	}

	r.Run(utils.HttpPort)
	
}
