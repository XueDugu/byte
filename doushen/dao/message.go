package dao

import (
	"time"
)

// Message 消息
type Message struct {
	ID        int64     // 自增主键ID
	FromId    int64     // 发送消息的用户ID
	ToId      int64     // 接收消息的用户ID
	Content   string    // 消息内容
	CreatedAt time.Time // 创建时间
}

// InsertMessage 插入消息
func InsertMessage(fromID, toID int64, content string) (int64, time.Time, error) {
	message := Message{
		FromId:    fromID,
		ToId:      toID,
		Content:   content,
		CreatedAt: time.Now(),
	}
	res := db.Create(&message)
	return message.ID, message.CreatedAt, res.Error
}

// FindMessageByFromID 根据发送人查询
func FindMessageByFromID(fromID int64) []Message {
	var messages []Message
	db.Where(&Message{FromId: fromID}).Order("created_at desc").Find(&messages)
	return messages
}

// FindMessageByToID 根据接收人查询
func FindMessageByToID(toID int64) []Message {
	var messages []Message
	db.Where(&Message{ToId: toID}).Order("created_at desc").Find(&messages)
	return messages
}

// FindMessageByFromToID 根据发送人和接收人查询
func FindMessageByFromToID(fromID, toID int64) []Message {
	var messages []Message
	db.Where(&Message{FromId: fromID, ToId: toID}).Order("created_at desc").Find(&messages)
	return messages
}

// FindMessageByTwoID 查询两个用户之间的消息，不区分接收和发送
func FindMessageByTwoID(fromID, toID int64) []Message {
	var messages []Message
	db.Where(&Message{FromId: fromID, ToId: toID}).Or(&Message{FromId: toID, ToId: fromID}).Order("created_at desc").Find(&messages)
	return messages
}
