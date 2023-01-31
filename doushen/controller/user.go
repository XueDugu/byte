package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simple-demo/service"
)

// 函数的作用是注册
func Register(c *gin.Context) {
	username := c.Query("username")                             //获得字符串形式的用户名
	password := c.Query("password")                             //获得字符串形式的密码
	c.JSON(http.StatusOK, service.Register(username, password)) //JSON序列化用户名和密码
}

// 函数的作用是登录
func Login(c *gin.Context) {
	username := c.Query("username")                          //获得字符串形式的用户名
	password := c.Query("password")                          //获得字符串形式的密码
	c.JSON(http.StatusOK, service.Login(username, password)) //JSON序列化用户名和密码
}

// 函数的作用是用户信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")                      //接受字符串形式的反应
	c.JSON(http.StatusOK, service.UserInfo(token)) //JSON序列化反应
}
