package service

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var bucket = "liujingping-doushen"
var accessKey = "kasbWDX_UW5oxkxDdGuGSNM6NMAM4GlN2LBVMoUj"
var secretKey = "pRxoCVUzRW5NXqfJ4jmstZJRi-qY7KI6fdybaSmk"

// 函数的作用是上传文件
func UploadFile(key string, file io.Reader) (string, string, string, error) {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(file); err != nil {
		return "", "", "", err
	}
	data := buf.Bytes()
	return Upload(key, data)
}

// 函数的作用是上传数据
func Upload(key string, data []byte) (string, string, string, error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadongZheJiang2
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	fmt.Println(ret.Key, ret.PersistentID)
	return ret.Key, ret.Hash, ret.PersistentID, err
}

// 函数的作用是获取网址
func GetPublicURL(key string) string {
	domain := "http://rp3814hyw.bkt.clouddn.com"
	publicAccessURL := storage.MakePublicURL(domain, key)
	return publicAccessURL
}
