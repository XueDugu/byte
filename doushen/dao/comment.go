package dao

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64
	UserId    int64
	VideoId   int64
	Text      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// 函数的作用是添加评论
func InsertComment(userID, videoID int64, text string) (int64, time.Time, error) {
	comment := Comment{
		UserId:    userID,
		VideoId:   videoID,
		Text:      text,
		CreatedAt: time.Now(),
	}
	res := db.Create(&comment)
	return comment.ID, comment.CreatedAt, res.Error
}

// 函数的作用是删除评论
func DeleteComment(id int64) error {
	res := db.Delete(&Comment{}, id)
	return res.Error
}

// 函数的作用是通过视频的ID找到评论列表
func FindCommentsByVideoID(videoID int64) []Comment {
	var comments []Comment
	db.Where(&Comment{VideoId: videoID}).Find(&comments)
	return comments
}
