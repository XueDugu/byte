package service

import (
	"fmt"
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
)

func AddComment(user common.User, videoID int64, text string) (common.Response, common.Comment) {
	id, createDate, err := dao.InsertComment(user.Id, videoID, text)
	if err != nil {
		return common.Response{StatusCode: 2, StatusMsg: fmt.Sprint(err)}, common.Comment{}
	}
	return common.Response{StatusCode: 0}, common.Comment{
		Id:         id,
		User:       user,
		Content:    text,
		CreateDate: createDate.Format("2006-01-02 15:04:05"),
	}
}

func DeleteComment(commentID int64) (common.Response, common.Comment) {
	if err := dao.DeleteComment(commentID); err != nil {
		return common.Response{StatusCode: 2, StatusMsg: fmt.Sprint(err)}, common.Comment{}
	}
	return common.Response{StatusCode: 0}, common.Comment{}
}

func CommentList(videoID int64) []common.Comment {
	var res []common.Comment
	comments := dao.FindCommentsByVideoID(videoID)
	for _, comment := range comments {
		if username, err := dao.FindUserByID(comment.UserId); err == nil {
			res = append(res, common.Comment{
				Id: comment.ID,
				User: common.User{
					Id:            comment.UserId,
					Name:          username,
					FollowCount:   0,
					FollowerCount: 0,
					IsFollow:      false,
				},
				Content:    comment.Text,
				CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}

	}
	return res
}
