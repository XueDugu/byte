package service

import (
	"fmt"
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
	"time"
)

// Feed same demo video list for every request
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
			res = append(res, common.Video{
				Id: video.ID,
				Author: common.User{
					Id:            video.Author,
					Name:          name,
					FollowCount:   0,
					FollowerCount: 0,
					IsFollow:      false,
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