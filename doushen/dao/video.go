package dao

import "time"

// Video 视频
type Video struct {
	//gorm.Model
	ID        int64     // 自增主键
	Title     string    // 视频标题
	Author    int64     // 视频作者
	PlayUrl   string    // 视频播放URL
	CoverUrl  string    // 视频封面URL
	CreatedAt time.Time // 创建时间
}

// FindVideosByIdList 根据ID查找视频
func FindVideosByIdList(idList []int64) []Video {
	var videos []Video
	db.Find(&videos, idList)
	return videos
}

// FindVideosByCreatedTime 更具创建时间查找视频
func FindVideosByCreatedTime(t time.Time) []Video {
	var videos []Video
	// 获取全部记录（最多30条）
	db.Where("created_at < ?", t).Limit(30).Order("created_at desc").Find(&videos)
	return videos
}

// FindVideosByAuthor 根据作者查找视频
func FindVideosByAuthor(authorID int64) []Video {
	var videos []Video
	// 获取全部记录
	db.Where(&Video{Author: authorID}).Order("created_at desc").Find(&videos)
	return videos
}

// InsertVideo 插入新的视频
func InsertVideo(title string, author int64, playUrl string, coverUrl string) error {
	video := Video{
		Title:     title,
		Author:    author,
		PlayUrl:   playUrl,
		CoverUrl:  coverUrl,
		CreatedAt: time.Now(),
	}
	result := db.Create(&video) // 通过数据的指针来创建

	return result.Error // 返回 error
}
