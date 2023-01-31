package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
)

// 函数的作用是判断有无点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")                                 //接受字符串形式的反应
	videoID := c.Query("video_id")                            //接受字符串形式的视频发布者的ID
	actionType := c.Query("action_type")                      //获得字符串形式的反应状态
	ID, err := strconv.ParseInt(videoID, decimalism, bitSize) //获得数字形式的视频发布者的ID
	if err != nil {                                           //错误处理转换ID错误
		c.JSON(http.StatusOK, common.Response{StatusCode: FavoriteActionParseIntIDError, StatusMsg: "FavoriteAction ParseInt ID Error"})
		return
	}
	action, err := strconv.ParseInt(actionType, decimalism, bitSize) //获得数字形式的反应状态
	if err != nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: FavoriteActionParseIntIDError, StatusMsg: "FavoriteAction ParseInt action Error"})
		return
	}

	c.JSON(http.StatusOK, service.FavoriteAction(token, ID, action))
}

// 函数的作用是获取点赞列表
func FavoriteList(c *gin.Context) {
	token := c.Query("token") //接受字符串形式的反应
	if _, exist := service.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: UserDoesNotExist,
				StatusMsg:  "User doesn't exist",
			},
			VideoList: []common.Video{},
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: Success,
			},
			VideoList: service.FavoriteList(token), //接受视频点赞列表
		})
	}
}
