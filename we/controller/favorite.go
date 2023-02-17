package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
	"net/http"
	"strconv"
)

// FavoriteAction 点赞或取消点赞
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoID := c.Query("video_id")
	actionType := c.Query("action_type")
	ID, err := strconv.ParseInt(videoID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 2, StatusMsg: "FavoriteAction ParseInt ID Error"})
		return
	}
	action, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 2, StatusMsg: "FavoriteAction ParseInt action Error"})
		return
	}

	c.JSON(http.StatusOK, service.FavoriteAction(token, ID, action))
}

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if _, exist := service.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "User doesn't exist",
			},
			VideoList: []common.Video{},
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 0,
			},
			VideoList: service.FavoriteList(token),
		})
	}
}
