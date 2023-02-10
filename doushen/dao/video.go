package dao

import "time"

type Video struct {
	//gorm.Model
	ID        int64
	Title     string
	Author    int64
	PlayUrl   string
	CoverUrl  string
	CreatedAt time.Time
}

// 函数的作用通过ID列表找到视频
func FindVideosByIdList(id_list []int64) []Video {
	var videos []Video
	db.Find(&videos, id_list)
	return videos
}

// 函数的作用是通过创建时间找到视频
func FindVideosByCreatedTime(t time.Time) []Video {
	var videos []Video
	// 获取全部记录（最多30条）
	db.Where("created_at < ?", t).Limit(30).Order("created_at desc").Find(&videos)
	return videos
}

// 函数的作用是通过作者找到视频
func FindVideosByAuthor(authorID int64) []Video {
	var videos []Video
	// 获取全部记录
	db.Where(&Video{Author: authorID}).Order("created_at desc").Find(&videos)
	return videos
}

// 函数的作用是添加视频
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
