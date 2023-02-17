package service

import (
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
)

// UsersLoginInfo 缓存用户的登录信息，当查询不到时到数据库中查询，服务器重启时清空缓存
// test data: username=zhanglei, password=douyin
var UsersLoginInfo = map[string]common.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

// UserLoginResponse 用户登录响应
type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// UserResponse 用户响应
type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

// Register 注册
func Register(username string, password string) UserLoginResponse {
	if _, err := dao.FindUserByName(username); err != nil { // 用户名不能重复
		ID, _ := dao.CreateUserByNameAndPassword(username, password)
		token, _ := CreateToken(int(ID), username)
		UsersLoginInfo[token] = common.User{
			Id:            ID,
			Name:          username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		return UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   ID,
			Token:    token,
		}
	} else {
		return UserLoginResponse{Response: common.Response{StatusCode: 1, StatusMsg: "User already exist"}}
	}
}

// Login 登录
func Login(username string, password string) UserLoginResponse {
	if ID, err := dao.FindUserByNameAndPassword(username, password); err != nil {
		return UserLoginResponse{Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"}}
	} else {
		token, _ := CreateToken(int(ID), username)
		follow := dao.FindRelationsByFanID(ID)
		fans := dao.FindRelationsByUserID(ID)
		UsersLoginInfo[token] = common.User{
			Id:            ID,
			Name:          username,
			FollowCount:   int64(len(follow)),
			FollowerCount: int64(len(fans)),
			IsFollow:      false, // TODO
		}
		return UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   ID,
			Token:    token,
		}
	}
}

// UserInfo 获取用户信息
func UserInfo(token string) UserResponse {
	if user, exist := UsersLoginInfo[token]; exist {
		return UserResponse{
			Response: common.Response{StatusCode: 0},
			User:     user,
		}
	} else {
		return UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		}
	}
}

// 格式转换，把数据库得到的Video格式转换为common.Video格式
func convertUsers(users []dao.User) []common.User {
	var res []common.User
	for _, user := range users {
		follow := dao.FindRelationsByFanID(user.ID)
		fans := dao.FindRelationsByUserID(user.ID)
		res = append(res, common.User{
			Id:            user.ID,
			Name:          user.Name,
			FollowCount:   int64(len(follow)),
			FollowerCount: int64(len(fans)),
			IsFollow:      true,
		})
	}
	return res
}
