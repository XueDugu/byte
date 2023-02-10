package service

import (
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
)

// UsersLoginInfo use map to store user info, and key is token (username+password for demo)
// user data will be cleared every time the server starts
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

type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

// 函数的作用是注册用户
func Register(username string, password string) UserLoginResponse {
	if _, err := dao.FindUserByName(username); err != nil { // 用户名不能重复
		ID, _ := dao.CreateUserByNameAndPassword(username, password)
		token, _ := CreateToken(int(ID), username)
		UsersLoginInfo[token] = common.User{
			Id:            ID,
			Name:          username,
			FollowCount:   0,     // TODO
			FollowerCount: 0,     // TODO
			IsFollow:      false, // TODO
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

// 函数的作用是登录用户
func Login(username string, password string) UserLoginResponse {
	if ID, err := dao.FindUserByNameAndPassword(username, password); err != nil {
		return UserLoginResponse{Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"}}
	} else {
		token, _ := CreateToken(int(ID), username)
		UsersLoginInfo[token] = common.User{
			Id:            ID,
			Name:          username,
			FollowCount:   0,     // TODO
			FollowerCount: 0,     // TODO
			IsFollow:      false, // TODO
		}
		return UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   ID,
			Token:    token,
		}
	}
}

// 函数的作用是获得用户信息
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
