package service

import (
	"fmt"
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
	"time"
)

// Feed 查询最新的30个视频
func Feed(t time.Time, token string) ([]common.Video, time.Time) {
	videos := dao.FindVideosByCreatedTime(t)
	if len(videos) < 1 {
		return []common.Video{}, time.Now()
	}
	var userID int64 = 0
	if user, exists := UsersLoginInfo[token]; exists {
		userID = user.Id
	}
	return convertVideos(videos, userID), videos[len(videos)-1].CreatedAt
}

// 格式转换，把数据库得到的Video格式转换为common.Video格式
func convertVideos(videos []dao.Video, userID int64) []common.Video {
	var res []common.Video
	for _, video := range videos {
		if name, err := dao.FindUserByID(video.Author); err == nil {
			favorites := dao.FindFavoriteByVideoID(video.ID)
			favoriteCount := len(favorites)
			isFavorite := false
			for _, favorite := range favorites {
				if favorite.UserId == userID {
					isFavorite = true
				}
			}
			follow := dao.FindRelationsByFanID(video.Author)
			fans := dao.FindRelationsByUserID(video.Author)
			isFollow := false
			for _, fan := range fans {
				if fan.FanId == userID {
					isFollow = true
					break
				}
			}
			res = append(res, common.Video{
				Id: video.ID,
				Author: common.User{
					Id:            video.Author,
					Name:          name,
					FollowCount:   int64(len(follow)),
					FollowerCount: int64(len(fans)),
					IsFollow:      isFollow,
				},
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				FavoriteCount: int64(favoriteCount),
				CommentCount:  0,
				IsFavorite:    isFavorite,
			})
		} else {
			fmt.Printf("GetPublishList Error: %v", err)
		}
	}
	return res
}
