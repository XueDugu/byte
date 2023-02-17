package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed 获取最新30个短视频
func Feed(c *gin.Context) {
	latestTime := time.Now().Unix()
	sLatestTime := c.Query("latest_time")
	if sLatestTime != "" {
		var err error
		latestTime, err = strconv.ParseInt(sLatestTime, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: common.Response{
					StatusCode: 1,
				},
				VideoList: []common.Video{},
				NextTime:  time.Now().Unix(),
			})
			return
		}
	}
	token := c.Query("token")
	fmt.Printf("debug# time: %v, %v", latestTime, time.Unix(latestTime, 0))
	videoList, nextTime := service.Feed(time.Unix(latestTime, 0), token)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime.Unix(),
	})
}
