package routes

import (
	"Caj2PdfServer/controllers"
	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) {
	// 创建控制器实例
	userController := new(controllers.UserController)
	// 定义路由组
	v1 := router.Group("/api/v1")
	{
		v1.POST("/upload", controllers.Upload)
		// 定义路由
		v1.GET("/users", userController.GetUserList)
		v1.GET("/users/:id", userController.GetUserByID)
		v1.POST("/users", userController.CreateUser)
		v1.PUT("/users/:id", userController.UpdateUser)
		v1.DELETE("/users/:id", userController.DeleteUser)
	}
}
