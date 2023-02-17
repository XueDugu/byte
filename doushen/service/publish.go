package service

import (
	"fmt"
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
	"mime/multipart"
	"path/filepath"
)

// Publish 检查token鉴权，然后发布标题名为title的视频
func Publish(token string, title string, data *multipart.FileHeader) common.Response {
	filename := filepath.Base(data.Filename)
	user := UsersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	file, err := data.Open()
	if err != nil { // TODO: StatusCode枚举类型
		fmt.Printf("Publish Error: %v", err)
		return common.Response{StatusCode: 2, StatusMsg: "File open failed"}
	}
	println("debug# before UploadFile")
	filename, _, _, err = UploadFile(filename, file)
	if err != nil {
		fmt.Printf("Publish Error: %v", err)
		return common.Response{StatusCode: 3, StatusMsg: "File upload failed"}
	}
	playURL := GetPublicURL(filename)
	coverURL := "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg" // TODO: 制作封面
	println("debug# before InsertVideo")
	if err := dao.InsertVideo(title, user.Id, playURL, coverURL); err != nil {
		fmt.Printf("Publish Error: %v", err)
		return common.Response{
			StatusCode: 3,
			StatusMsg:  "Insert video failed",
		}
	}
	return common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	}
}

// GetPublishList 获取userID用户的发布视频的列表
func GetPublishList(userID int64) []common.Video {
	videos := dao.FindVideosByAuthor(userID)
	return convertVideos(videos, userID)
}
