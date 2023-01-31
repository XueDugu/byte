package dao

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        int64
	UserId    int64
	VideoId   int64
	Text      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

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

func DeleteComment(id int64) error {
	res := db.Delete(&Comment{}, id)
	return res.Error
}

func FindCommentsByVideoID(videoID int64) []Comment {
	var comments []Comment
	db.Where(&Comment{VideoId: videoID}).Find(&comments)
	return comments
}
