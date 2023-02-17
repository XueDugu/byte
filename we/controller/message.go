package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/simple-demo/common"
	"github.com/simple-demo/service"
	"net/http"
	"strconv"
)

//var tempChat = map[string][]common.Message{}
//
//var messageIdSequence = int64(1)

type ChatResponse struct {
	common.Response
	MessageList []common.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	if user, exist := service.UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.ParseInt(toUserId, 10, 64)
		if err := service.MessageAction(user.Id, userIdB, content); err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 2, StatusMsg: err.Error()})
		} else {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	if user, exist := service.UsersLoginInfo[token]; exist {
		userIdB, _ := strconv.ParseInt(toUserId, 10, 64)
		println("user id", user.Id)
		println("userIdB", userIdB)
		c.JSON(http.StatusOK, ChatResponse{Response: common.Response{StatusCode: 0}, MessageList: service.MessageChat(user.Id, userIdB)})
		//c.JSON(http.StatusOK, ChatResponse{Response: common.Response{StatusCode: 0}, MessageList: []common.Message{}})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

//func genChatKey(userIdA int64, userIdB int64) string {
//	if userIdA > userIdB {
//		return fmt.Sprintf("%d_%d", userIdB, userIdA)
//	}
//	return fmt.Sprintf("%d_%d", userIdA, userIdB)
//}
