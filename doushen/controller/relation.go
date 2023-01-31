package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
)

type UserListResponse struct {
	common.Response               //用户回复列表
	UserList        []common.User `json:"user_list"` //用户列表
}

// 函数的作用是判断用户是否注册成功
func RelationAction(c *gin.Context) {
	token := c.Query("token") //接受字符串形式的反应

	if _, exist := service.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: Success})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: UserDoesNotExist, StatusMsg: "User doesn't exist"}) //错误处理用户注册错误
	}
}

// 函数的作用是查看自己关注的人的列表
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{ //JSON序列化关注列表
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: []common.User{DemoUser},
	})
}

// 函数的作用是查看关注自己的人的列表
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{ //JSON序列化关注他的人
		Response: common.Response{
			StatusCode: Success,
		},
		UserList: []common.User{DemoUser},
	})
}

// 函数的作用是查看好友列表
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{ //JSON序列化用户好友列表
		Response: common.Response{
			StatusCode: Success,
		},
		UserList: []common.User{DemoUser},
	})
}
