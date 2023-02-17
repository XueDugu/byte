package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDb 初始化数据库
func InitDb() {
	// 本地数据库
	//dsn := "root:xxzj9911@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	// 云数据库
	url := "d617fe90ba4d.c.methodot.com:33729"
	dsn := "root:xxzj9911@tcp(" + url + ")/douyin?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 不能用 :=
	if err != nil {
		println(err)
		panic(err)
	}
	println("connect to database:" + db.Name())
}
