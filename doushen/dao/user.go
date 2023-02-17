package dao

import (
	"fmt"
)

// User 用户
type User struct {
	//gorm.Model
	ID       int64  // 自增主键
	Name     string // 用户名
	Password string // 密码
}

// 查询所有的用户（用于测试）
func findAllUsers() {
	var users []User
	// 获取全部记录
	result := db.Find(&users)
	// SELECT * FROM users;
	println(users[0].Name)
	println(result.RowsAffected) // 返回找到的记录数，相当于 `len(users)`
	println(result.Error)        // returns error
}

// FindUserByID 根据ID查询用户
func FindUserByID(id int64) (username string, err error) {
	var users []User
	db.Where([]int64{id}).Find(&users)
	if len(users) < 1 {
		return "", fmt.Errorf("FindUserByID not found")
	}
	return users[0].Name, nil
}

// FindUsersByIDList 根据ID列表查询若干用户
func FindUsersByIDList(idList []int64) []User {
	var users []User
	db.Find(&users, idList)
	return users
}

// FindUserByNameAndPassword 根据用户名和密码查询用户
func FindUserByNameAndPassword(username string, password string) (ID int64, err error) {
	var users []User
	db.Where(&User{Name: username, Password: password}).Find(&users)
	if len(users) < 1 {
		return 0, fmt.Errorf("UserByNameAndPassword not found")
	}
	return users[0].ID, nil
}

// FindUserByName 根据用户名查询用户
func FindUserByName(username string) (ID int64, err error) {
	var users []User
	db.Where(&User{Name: username}).Find(&users)
	if len(users) < 1 {
		return 0, fmt.Errorf("UserByNameAndPassword not found")
	}
	return users[0].ID, nil
}

// CreateUserByNameAndPassword 使用用户名和密码创建用户
func CreateUserByNameAndPassword(username string, password string) (ID int64, err error) {
	user := User{Name: username, Password: password}
	result := db.Create(&user) // 通过数据的指针来创建
	return user.ID, result.Error
}
