package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	common.Response
	UserList []common.User `json:"user_list"`
}

type FriendListResponse struct {
	common.Response
	UserList []common.FriendUser `json:"user_list"`
}

// RelationAction 关注或取关
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	userID, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 2, StatusMsg: "RelationAction ParseInt ID Error"})
		return
	}
	action, err := strconv.ParseInt(actionType, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 2, StatusMsg: "RelationAction ParseInt action Error"})
		return
	}

	if user, exist := service.UsersLoginInfo[token]; exist {
		err := service.RelationAction(user, userID, int8(action))
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 2, StatusMsg: fmt.Sprint(err)})
		} else {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList 查询关注列表
func FollowList(c *gin.Context) {
	token := c.Query("token")
	if user, exist := service.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{
				StatusCode: 0,
			},
			UserList: service.FollowList(user),
		})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowerList 查询粉丝列表
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	if user, exist := service.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserListResponse{
			Response: common.Response{
				StatusCode: 0,
			},
			UserList: service.FollowerList(user),
		})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FriendList TODO 暂时实现为粉丝列表
func FriendList(c *gin.Context) {
	token := c.Query("token")
	if user, exist := service.UsersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, FriendListResponse{
			Response: common.Response{
				StatusCode: 0,
			},
			UserList: service.FriendList(user),
		})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
