package service

import "testing"

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

func TestGetPublicURL(t *testing.T) {
	filename := "test.txt"
	GetPublicURL(filename)
}
