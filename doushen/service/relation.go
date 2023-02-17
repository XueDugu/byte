package service

import (
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
)

// RelationAction 首先尝试更新数据库，如果没有数据，则插入新的数据
func RelationAction(fan common.User, toUserId int64, actionType int8) error {
	if affectRows := dao.UpdateRelation(fan.Id, toUserId, actionType); affectRows < 1 {
		err := dao.InsertFavorite(fan.Id, toUserId, actionType)
		return err
	}
	return nil
}

// FollowList 查询关注列表
func FollowList(user common.User) []common.User {
	follows := dao.FindRelationsByFanID(user.Id)
	var IDList []int64
	for _, follow := range follows {
		IDList = append(IDList, follow.UserId)
	}
	res := dao.FindUsersByIDList(IDList)
	return convertUsers(res)
}

// FollowerList 查询粉丝列表
func FollowerList(user common.User) []common.User {
	fans := dao.FindRelationsByUserID(user.Id)
	var IDList []int64
	for _, fan := range fans {
		IDList = append(IDList, fan.FanId)
	}
	res := dao.FindUsersByIDList(IDList)
	return convertUsers(res)
}

// FriendList 查询粉丝列表，然后对每个粉丝，查询消息
func FriendList(user common.User) []common.FriendUser {
	fans := dao.FindRelationsByUserID(user.Id)
	var IDList []int64
	for _, fan := range fans {
		IDList = append(IDList, fan.FanId)
	}
	res := dao.FindUsersByIDList(IDList)
	users := convertUsers(res)
	friends := make([]common.FriendUser, len(users))
	for _, u := range users {
		messages := dao.FindMessageByTwoID(u.Id, user.Id)
		content := ""
		var msgType int64 = 0
		if len(messages) > 0 {
			content = messages[0].Content
			if messages[0].FromId == user.Id {
				msgType = 1
			}
		}
		friends = append(friends, common.FriendUser{
			Id:            u.Id,
			Name:          u.Name,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      u.IsFollow,
			Avatar:        "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg", // TODO: 头像URL
			Message:       content,
			MsgType:       msgType,
		})
	}
	return friends
}
