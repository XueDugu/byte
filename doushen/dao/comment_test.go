package dao

import (
	"testing"

	"gorm.io/gorm/utils/tests"
)

// 函数的作用是测试添加评论
func TestInsertComment(t *testing.T) {
	Init_db()
	id, _, err := InsertComment(1, 1, "test ljp") //添加评论
	tests.AssertEqual(t, err, nil)
	tests.AssertEqual(t, id, 1)
}

// 函数的作用是测试删除评论
func TestDeleteComment(t *testing.T) {
	Init_db()
	err := DeleteComment(1) //删除评论
	tests.AssertEqual(t, err, nil)
}
