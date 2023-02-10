package dao

type Favorite struct {
	ID      int64
	UserId  int64
	VideoId int64
	Status  int8
}

// 函数的作用是通过用户的ID找到喜欢的内容列表
func FindFavoriteByUserID(userID int64) []Favorite {
	var favorites []Favorite
	db.Where(&Favorite{UserId: userID}).Find(&favorites)
	return favorites
}

// 函数的作用是通过视频的ID找到喜欢的内容列表
func FindFavoriteByVideoID(videoID int64) []Favorite {
	var favorites []Favorite
	db.Where(&Favorite{VideoId: videoID}).Find(&favorites)
	return favorites
}

// 函数的作用是更新喜欢的内容列表
func UpdateFavorite(userID int64, videoID int64, status int8) int64 {
	res := db.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userID, videoID).Update("Status", status)
	return res.RowsAffected
}

// 函数的作用是增加喜欢的内容列表
func InsertFavorite(userID int64, videoID int64, status int8) error {
	res := db.Create(&Favorite{UserId: userID, VideoId: videoID, Status: status})
	return res.Error
}
