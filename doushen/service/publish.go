package service

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
)

const (
	Success           = 0
	FileOpenFailed    = 2
	FileUploadFailed  = 3
	InsertVideoFailed = 3
)

// 函数的作用是确认反应并上传文件
func Publish(token string, title string, data *multipart.FileHeader) common.Response {
	filename := filepath.Base(data.Filename)
	user := UsersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	file, err := data.Open()
	if err != nil { // TODO: StatusCode枚举类型
		fmt.Printf("Publish Error: %v", err)
		return common.Response{StatusCode: FileOpenFailed, StatusMsg: "File open failed"}
	}
	filename, _, _, err = UploadFile(filename, file)
	if err != nil {
		fmt.Printf("Publish Error: %v", err)
		return common.Response{StatusCode: FileUploadFailed, StatusMsg: "File upload failed"}
	}
	playURL := GetPublicURL(filename)
	coverURL := "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg" // TODO: 制作封面

	if err := dao.InsertVideo(title, user.Id, playURL, coverURL); err != nil {
		fmt.Printf("Publish Error: %v", err)
		return common.Response{
			StatusCode: InsertVideoFailed,
			StatusMsg:  "Insert video failed",
		}
	}

	return common.Response{
		StatusCode: Success,
		StatusMsg:  finalName + " uploaded successfully",
	}
}

// 函数的作用是找到用户所有的视频
func GetPublishList(userID int64) []common.Video {
	videos := dao.FindVideosByAuthor(userID)
	return convertVideos(videos, userID)
}
