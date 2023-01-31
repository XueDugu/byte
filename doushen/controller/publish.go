package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
)

type VideoListResponse struct {
	common.Response                //视频回复列表回复
	VideoList       []common.Video `json:"video_list"` //视频列表
}

// 函数的作用是
func Publish(c *gin.Context) {
	token := c.PostForm("token") //获得反应
	title := c.PostForm("title") //获得标题

	if _, exist := service.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: UserDoesNotExist, StatusMsg: "User doesn't exist"}) //错误处理用户不存在
		return                                                                                                //结束进程
	}

	data, err := c.FormFile("data") //获得第一个文件
	if err != nil {                 //错误处理用户不存在
		c.JSON(http.StatusOK, common.Response{
			StatusCode: UserDoesNotExist,
			StatusMsg:  err.Error(),
		})
		return //结束进程
	}

	c.JSON(http.StatusOK, service.Publish(token, title, data)) //JSON序列化
}

// 函数的作用是获取登录用户的发布列表
func PublishList(c *gin.Context) {
	userID := c.Query("user_id") //获得字符串形式的用户的ID
	ID, err := strconv.ParseInt(userID, decimalism, bitSize)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: Success,
			},
			VideoList: []common.Video{},
		})
		return //结束进程
	}
	token := c.Query("token")                    //接受字符串形式的反应
	user, exist := service.UsersLoginInfo[token] //获得发布列表
	if !exist {                                  //错误处理用户不存在
		c.JSON(http.StatusOK, common.Response{StatusCode: UserDoesNotExist, StatusMsg: "User doesn't exist"})
		return
	}
	if user.Id != ID {
		c.JSON(http.StatusOK, common.Response{StatusCode: UserIDDoesNotMatchToken, StatusMsg: "User ID doesn't match token"}) //错误处理用户ID不匹配
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{ //JSON序列化发布列表
		Response: common.Response{
			StatusCode: Success,
		},
		VideoList: service.GetPublishList(ID),
	})
}

//// Publish check token then save upload file to public directory
//func PublishDemo(c *gin.Context) {
//	token := c.PostForm("token")
//
//	if _, exist := service.UsersLoginInfo[token]; !exist {
//		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//		return
//	}
//
//	data, err := c.FormFile("data")
//	if err != nil {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//
//	filename := filepath.Base(data.Filename)
//	user := service.UsersLoginInfo[token]
//	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
//	saveFile := filepath.Join("./public/", finalName)
//	if err := c.SaveUploadedFile(data, saveFile); err != nil {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, common.Response{
//		StatusCode: 0,
//		StatusMsg:  finalName + " uploaded successfully",
//	})
//}
//
//// PublishList all users have same publish video list
//func PublishListDemo(c *gin.Context) {
//	c.JSON(http.StatusOK, VideoListResponse{
//		Response: common.Response{
//			StatusCode: 0,
//		},
//		VideoList: DemoVideos,
//	})
//}
