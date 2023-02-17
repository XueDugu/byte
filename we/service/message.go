package service

import (
	"encoding/json"
	"fmt"
	"github.com/simple-demo/common"
	"github.com/simple-demo/dao"
	"io"
	"net"
	"sync"
)

var chatConnMap = sync.Map{}

func RunMessageServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		var event = common.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Printf("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			fmt.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		pushEvent := common.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(fromID, toID int64, content string) error {
	_, _, err := dao.InsertMessage(fromID, toID, content)
	return err
}

// MessageChat all users have same follow list
func MessageChat(fromID, toID int64) []common.Message {
	messages := dao.FindMessageByTwoID(fromID, toID)
	res := make([]common.Message, len(messages))
	for _, message := range messages {
		println(message.Content)
		res = append(res, common.Message{
			Id:         message.ID,
			ToUserId:   message.ToId,
			FromUserId: message.FromId,
			Content:    message.Content,
			CreateTime: message.CreatedAt.Unix(),
		})
	}
	return res
}
