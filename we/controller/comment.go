package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
	"net/http"
	"strconv"
)

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

// CommentActionResponse 评论行为响应
type CommentActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

// CommentAction 发表评论
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := service.UsersLoginInfo[token]; exist {
		if actionType == "1" { // 添加评论
			videoID := c.Query("video_id")
			ID, err := strconv.ParseInt(videoID, 10, 64)
			if err != nil {
				c.JSON(http.StatusOK, CommentListResponse{
					Response: common.Response{
						StatusCode: 3,
						StatusMsg:  "CommentList ParseInt Error",
					},
					CommentList: []common.Comment{},
				})
				return
			}
			text := c.Query("comment_text")
			res, comment := service.AddComment(user, ID, text)
			c.JSON(http.StatusOK, CommentActionResponse{Response: res, Comment: comment})
		} else if actionType == "2" { // 删除评论
			commentID := c.Query("comment_id")
			ID, err := strconv.ParseInt(commentID, 10, 64)
			if err != nil {
				c.JSON(http.StatusOK, CommentListResponse{
					Response: common.Response{
						StatusCode: 3,
						StatusMsg:  "CommentList ParseInt Error",
					},
					CommentList: []common.Comment{},
				})
				return
			}
			res, comment := service.DeleteComment(ID)
			c.JSON(http.StatusOK, CommentActionResponse{Response: res, Comment: comment})
		} else {
			c.JSON(http.StatusOK, common.Response{StatusCode: 5, StatusMsg: "Wrong action type"})
		}
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	token := c.Query("token")
	user, exist := service.UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	videoID := c.Query("video_id")
	ID, err := strconv.ParseInt(videoID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: common.Response{
				StatusCode: 3,
				StatusMsg:  "CommentList ParseInt Error",
			},
			CommentList: []common.Comment{},
		})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: service.CommentList(ID, user.Id),
	})
}
