package service

import (
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
)

// 函数的作用是修改点赞状态
func FavoriteAction(token string, videoId int64, actionType int64) common.Response {
	user, exist := UsersLoginInfo[token]
	if exist {
		if affectRows := dao.UpdateFavorite(user.Id, videoId, int8(actionType)); affectRows < 1 {
			err := dao.InsertFavorite(user.Id, videoId, int8(actionType))
			if err != nil {
				return common.Response{StatusCode: 2, StatusMsg: "InsertFavorite Error"}
			}
		}
		return common.Response{StatusCode: 0}
	} else {
		return common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"}
	}
}

// 函数的作用是获取点赞列表
func FavoriteList(token string) []common.Video {
	user, exist := UsersLoginInfo[token]
	if exist {
		favorites := dao.FindFavoriteByUserID(user.Id)
		var IDList []int64
		for _, favorite := range favorites {
			IDList = append(IDList, favorite.VideoId)
		}
		res := dao.FindVideosByIdList(IDList)
		return convertVideos(res, user.Id)
	} else {
		return []common.Video{}
	}
}
