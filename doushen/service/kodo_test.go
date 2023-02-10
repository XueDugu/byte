package service

import "testing"

//函数的作用是测试上传功能
func TestUpload(t *testing.T) {
	filename := "test.txt"
	data := []byte("hello, this is qiniu cloud")
	key, hash, ID, err := Upload(filename, data)
	if err != nil {
		println(err)
		return
	}
	println(key, hash, ID)
}

//函数的作用是测试获取网址功能
func TestGetPublicURL(t *testing.T) {
	filename := "test.txt"
	GetPublicURL(filename)
}
