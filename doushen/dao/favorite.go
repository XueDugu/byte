package dao

type Favorite struct {
	ID      int64
	UserId  int64
	VideoId int64
	Status  int8
}

func FindFavoriteByUserID(userID int64) []Favorite {
	var favorites []Favorite
	db.Where(&Favorite{UserId: userID}).Find(&favorites)
	return favorites
}

func FindFavoriteByVideoID(videoID int64) []Favorite {
	var favorites []Favorite
	db.Where(&Favorite{VideoId: videoID}).Find(&favorites)
	return favorites
}

func UpdateFavorite(userID int64, videoID int64, status int8) int64 {
	res := db.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userID, videoID).Update("Status", status)
	return res.RowsAffected
}

func InsertFavorite(userID int64, videoID int64, status int8) error {
	res := db.Create(&Favorite{UserId: userID, VideoId: videoID, Status: status})
	return res.Error
}
