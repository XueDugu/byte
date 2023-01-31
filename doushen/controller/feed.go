package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
)

type FeedResponse struct {
	common.Response                //回复
	VideoList       []common.Video `json:"video_list,omitempty"` //视频列表
	NextTime        int64          `json:"next_time,omitempty"`  //时间
}

// 函数的作用是获取视频列表和观看时间
func Feed(c *gin.Context) {
	latestTime := time.Now().Unix()       //获得执行到此处的时间
	sLatestTime := c.Query("latest_time") //接受字符串形式的持续时间
	if sLatestTime != "" {
		var err error
		latestTime, err = strconv.ParseInt(sLatestTime, decimalism, bitSize) //获得数字形式的持续时间
		if err != nil {                                                      //错误处理用户不存在
			c.JSON(http.StatusOK, FeedResponse{
				Response: common.Response{
					StatusCode: UserDoesNotExist,
				},
				VideoList: []common.Video{},
				NextTime:  time.Now().Unix(), //获得执行到此处的时间
			})
			return //结束进程
		}
	}
	token := c.Query("token")                                                         //接受字符串形式的反应
	fmt.Printf("debug# time: %v, %v", latestTime, time.Unix(latestTime, nanoseconds)) //打印时间
	videoList, nextTime := service.Feed(time.Unix(latestTime, nanoseconds), token)    //获得视频列表
	c.JSON(http.StatusOK, FeedResponse{                                               //JSON序列化成功
		Response:  common.Response{StatusCode: Success},
		VideoList: videoList,
		NextTime:  nextTime.Unix(),
	})
}
