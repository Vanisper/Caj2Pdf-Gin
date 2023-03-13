package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController 是用户控制器
type UserController struct{}

// GetUserList 处理获取用户列表的请求
func (u *UserController) GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "获取用户列表成功",
	})
}

// GetUserByID 处理获取指定用户的请求
func (u *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "获取用户" + id + "的信息成功",
	})
}

// CreateUser 处理创建用户的请求
func (u *UserController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "创建用户成功",
	})
}

// UpdateUser 处理更新用户的请求
func (u *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "更新用户" + id + "的信息成功",
	})
}

// DeleteUser 处理删除用户的请求
func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "删除用户" + id + "成功",
	})
}
