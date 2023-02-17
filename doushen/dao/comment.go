package dao

import (
	"gorm.io/gorm"
	"time"
)

// Comment 评论
type Comment struct {
	ID        int64          // 自增主键ID
	UserId    int64          // 用户ID
	VideoId   int64          // 视频ID
	Text      string         // 评论内容
	CreatedAt time.Time      // 创建时间
	DeletedAt gorm.DeletedAt // 删除时间（软删除）
}

// InsertComment 插入评论
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

// DeleteComment 删除评论
func DeleteComment(id int64) error {
	res := db.Delete(&Comment{}, id)
	return res.Error
}

// FindCommentsByVideoID 根据视频ID查找评论
func FindCommentsByVideoID(videoID int64) []Comment {
	var comments []Comment
	db.Where(&Comment{VideoId: videoID}).Find(&comments)
	return comments
}
