package dao

import (
	"fmt"
)

type User struct {
	//gorm.Model
	ID       int64
	Name     string
	Password string
}

// 函数的作用是显示所有用户
func findAllUsers() {
	var users []User
	// 获取全部记录
	result := db.Find(&users)
	// SELECT * FROM users;
	println(users[0].Name)
	println(result.RowsAffected) // 返回找到的记录数，相当于 `len(users)`
	println(result.Error)        // returns error
}

// 函数的作用是通过用户的ID找到用户
func FindUserByID(id int64) (username string, err error) {
	var users []User
	db.Where([]int64{id}).Find(&users)
	if len(users) < 1 {
		return "", fmt.Errorf("FindUserByID not found")
	}
	return users[0].Name, nil
}

// 函数的作用是通过用户名和密码找到用户
func FindUserByNameAndPassword(username string, password string) (ID int64, err error) {
	var users []User
	db.Where(&User{Name: username, Password: password}).Find(&users)
	if len(users) < 1 {
		return 0, fmt.Errorf("UserByNameAndPassword not found")
	}
	return users[0].ID, nil
}

// 函数的作用是通过用户名找到用户
func FindUserByName(username string) (ID int64, err error) {
	var users []User
	db.Where(&User{Name: username}).Find(&users)
	if len(users) < 1 {
		return 0, fmt.Errorf("UserByNameAndPassword not found")
	}
	return users[0].ID, nil
}

func CreateUserByNameAndPassword(username string, password string) (ID int64, err error) {
	user := User{Name: username, Password: password}
	result := db.Create(&user) // 通过数据的指针来创建
	return user.ID, result.Error
}
