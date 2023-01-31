package controller

import "github.com/simple-demo/common"

var DemoVideos = []common.Video{
	{
		Id:            1,                                                                      //测试视频发布者的ID
		Author:        DemoUser,                                                               //测试视频发布者
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",                             //测试视频所在的网址
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg", //测试视频覆盖的网址
		FavoriteCount: 0,                                                                      //测试视频点赞数
		CommentCount:  0,                                                                      //测试视频评论数
		IsFavorite:    false,                                                                  //是否点赞测试视频的状态
	},
}

var DemoComments = []common.Comment{
	{
		Id:         1,              //测试评论者的ID
		User:       DemoUser,       //测试评论者
		Content:    "Test Comment", //测试评论内容
		CreateDate: "05-01",        //测试评论时间
	},
}

var DemoUser = common.User{
	Id:            1,          //测试评论者的ID
	Name:          "TestUser", //测试评论者的用户名
	FollowCount:   0,          //测试评论者关注的人数
	FollowerCount: 0,          //关注测试评论者的人数
	IsFollow:      false,      //是否关注测试评论者的状态
}
