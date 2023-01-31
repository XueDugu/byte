package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
)

const (
	AddCommentMode                = "1" // 添加评论状态
	DeleteCommentMode             = "2" // 删除评论状态
	Success                       = 0   //成功
	UserDoesNotExist              = 1   //用户不存在
	FavoriteActionParseIntIDError = 2   //点赞转换错误
	CommentListParseIntError      = 3   //注释列表解析错误
	UserIDDoesNotMatchToken       = 4   //用户ID不匹配
	WrongActionType               = 5   //未知状态
	decimalism                    = 10  //十进制
	bitSize                       = 64  //64位大小
	nanoseconds                   = 0   //纳秒忽略
	AddOne                        = 1   //原子锁加一
)

type CommentListResponse struct {
	common.Response                  //回复
	CommentList     []common.Comment `json:"comment_list,omitempty"` //评论列表
}

type CommentActionResponse struct {
	common.Response                //回复
	Comment         common.Comment `json:"comment,omitempty"` //评论
}

// 函数的作用是对评论进行添加或删除操作
func CommentAction(c *gin.Context) {
	token := c.Query("token")            //接受字符串形式的反应
	actionType := c.Query("action_type") //接受字符串形式的反应状态

	if user, exist := service.UsersLoginInfo[token]; exist {
		if actionType == AddCommentMode { // 添加评论
			videoID := c.Query("video_id")                            //接受字符串形式的视频发布者的ID
			ID, err := strconv.ParseInt(videoID, decimalism, bitSize) //获得数字形式的视频发布者的ID
			if err != nil {                                           //错误处理转换ID错误
				c.JSON(http.StatusOK, CommentListResponse{
					Response: common.Response{
						StatusCode: CommentListParseIntError,     //回复的状态
						StatusMsg:  "CommentList ParseInt Error", //回复的状态信息
					},
					CommentList: []common.Comment{},
				})
				return //结束进程
			}
			text := c.Query("comment_text")                                               //接受评论信息
			res, comment := service.AddComment(user, ID, text)                            //获得回复和评论
			c.JSON(http.StatusOK, CommentActionResponse{Response: res, Comment: comment}) //JSON序列化回复
		} else if actionType == DeleteCommentMode { // 删除评论
			commentID := c.Query("comment_id")                          //接受字符串形式的评论发布者的ID
			ID, err := strconv.ParseInt(commentID, decimalism, bitSize) //获得数字形式的视频发布者的ID
			if err != nil {                                             //错误处理转换ID错误
				c.JSON(http.StatusOK, CommentListResponse{
					Response: common.Response{
						StatusCode: CommentListParseIntError,     //回复的状态
						StatusMsg:  "CommentList ParseInt Error", //回复的状态信息
					},
					CommentList: []common.Comment{},
				})
				return //结束进程
			}
			res, comment := service.DeleteComment(ID)                                     //删除评论
			c.JSON(http.StatusOK, CommentActionResponse{Response: res, Comment: comment}) //JSON序列化回复
		} else {
			c.JSON(http.StatusOK, common.Response{StatusCode: WrongActionType, StatusMsg: "Wrong action type"}) //错误处理未知评论状态
		}
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: UserDoesNotExist, StatusMsg: "User doesn't exist"}) //错误处理用户不存在
	}
}

// 函数的作用是获取评论列表
func CommentList(c *gin.Context) {
	token := c.Query("token") //接受反应
	if _, exist := service.UsersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: UserDoesNotExist, StatusMsg: "User doesn't exist"}) //错误处理用户不存在
		return
	}
	videoID := c.Query("video_id")                            //获得字符串形式的视频发布者的ID
	ID, err := strconv.ParseInt(videoID, decimalism, bitSize) //获得数字形式的视频发布者的ID
	if err != nil {                                           //错误处理转换ID错误
		c.JSON(http.StatusOK, CommentListResponse{
			Response: common.Response{
				StatusCode: CommentListParseIntError,     //回复的状态
				StatusMsg:  "CommentList ParseInt Error", //回复的状态信息
			},
			CommentList: []common.Comment{},
		})
		return //结束进程Success
	}
	c.JSON(http.StatusOK, CommentListResponse{ //JSON序列化回复成功
		Response:    common.Response{StatusCode: Success},
		CommentList: service.CommentList(ID),
	})
}
