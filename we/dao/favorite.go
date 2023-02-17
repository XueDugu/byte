package dao

// Favorite 点赞
type Favorite struct {
	ID      int64 // 自增主键
	UserId  int64 // 用户ID
	VideoId int64 // 视频ID
	Status  int8  // 点赞状态（0表示未点赞，1表示已点赞，这使得用户可以取消点赞）
}

// FindFavoriteByUserID 根据用户ID查询点赞状态
func FindFavoriteByUserID(userID int64) []Favorite {
	var favorites []Favorite
	db.Where(&Favorite{UserId: userID}).Find(&favorites)
	return favorites
}

// FindFavoriteByVideoID 根据视频ID查询点赞状态
func FindFavoriteByVideoID(videoID int64) []Favorite {
	var favorites []Favorite
	db.Where(&Favorite{VideoId: videoID}).Find(&favorites)
	return favorites
}

// UpdateFavorite 更新点赞状态
func UpdateFavorite(userID int64, videoID int64, status int8) int64 {
	res := db.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userID, videoID).Update("Status", status)
	return res.RowsAffected
}

// InsertFavorite 插入新的点赞数据
func InsertFavorite(userID int64, videoID int64, status int8) error {
	res := db.Create(&Favorite{UserId: userID, VideoId: videoID, Status: status})
	return res.Error
}
