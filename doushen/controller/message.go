package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
)

var tempChat = map[string][]common.Message{} //创建信息整体

var messageIdSequence = int64(1) //第一个信息的ID

type ChatResponse struct {
	common.Response                  //信息的回复
	MessageList     []common.Message `json:"message_list"` //信息列表
}

// 函数的作用是检查反应是否有效并更新信息整体
func MessageAction(c *gin.Context) {
	token := c.Query("token")         //接受字符串形式的反应
	toUserId := c.Query("to_user_id") //接受字符串形式的信息接收者的ID
	content := c.Query("content")     //接受字符串形式的信息的内容

	if user, exist := service.UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)           //获得数字形式的信息接收者的ID
		chatKey := genChatKey(user.Id, int64(userIdB)) //获得字符串形式的信息接收者的ID和信息发送者的ID结合

		atomic.AddInt64(&messageIdSequence, AddOne) //赋予信息ID
		curMessage := common.Message{               //信息初始化
			Id:         messageIdSequence,
			Content:    content,
			CreateTime: time.Now().Format(time.Kitchen),
		}

		if messages, exist := tempChat[chatKey]; exist {
			tempChat[chatKey] = append(messages, curMessage) //将新信息加到原信息后面构成信息整体
		} else {
			tempChat[chatKey] = []common.Message{curMessage} //将新信息作为信息整体
		}
		c.JSON(http.StatusOK, common.Response{StatusCode: Success}) //JSON序列化信息成功
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: UserDoesNotExist, StatusMsg: "User doesn't exist"}) //错误处理用户不存在
	}
}

// 函数的作用是将信息收发关系整合进整体信息
func MessageChat(c *gin.Context) {
	token := c.Query("token")         //接受字符串形式的反应
	toUserId := c.Query("to_user_id") //接受字符串形式的信息接收者的ID

	if user, exist := service.UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)           //获得数字形式的信息接收者的ID
		chatKey := genChatKey(user.Id, int64(userIdB)) //获得字符串形式的信息接收者的ID和信息发送者的ID结合

		c.JSON(http.StatusOK, ChatResponse{Response: common.Response{StatusCode: Success}, MessageList: tempChat[chatKey]}) //JSON序列化成功
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: UserDoesNotExist, StatusMsg: "User doesn't exist"}) //错误处理用户不存在
	}
}

// 函数的作用是用下划线将ID从小到大连接起来
func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
