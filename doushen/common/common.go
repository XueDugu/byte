package common

type Response struct {
	StatusCode int32  `json:"status_code"`          //回复的状态码
	StatusMsg  string `json:"status_msg,omitempty"` //回复的状态信息
}

type Video struct {
	Id            int64  `json:"id,omitempty"`             //视频作者的ID
	Author        User   `json:"author"`                   //视频的作者
	PlayUrl       string `json:"play_url,omitempty"`       //展示视频的网址
	CoverUrl      string `json:"cover_url,omitempty"`      //视频覆盖的网址
	FavoriteCount int64  `json:"favorite_count,omitempty"` //收藏视频的人数
	CommentCount  int64  `json:"comment_count,omitempty"`  //评论视频的人数
	IsFavorite    bool   `json:"is_favorite,omitempty"`    //是否收藏的状态
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`          //评论者的ID
	User       User   `json:"user"`                  //评论的人
	Content    string `json:"content,omitempty"`     //评论的内容
	CreateDate string `json:"create_date,omitempty"` //评论的时间
}

type User struct {
	Id            int64  `json:"id,omitempty"`             //用户的ID
	Name          string `json:"name,omitempty"`           //用户名
	FollowCount   int64  `json:"follow_count,omitempty"`   //用户关注的人数
	FollowerCount int64  `json:"follower_count,omitempty"` //关注用户的人数
	IsFollow      bool   `json:"is_follow,omitempty"`      //是否关注的状态
}

type Message struct {
	Id         int64  `json:"id,omitempty"`          //信息发送者的ID
	Content    string `json:"content,omitempty"`     //信息的内容
	CreateTime string `json:"create_time,omitempty"` //信息创建的时间
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`     //信息发送者的ID
	ToUserId   int64  `json:"to_user_id,omitempty"`  //信息接收者的ID
	MsgContent string `json:"msg_content,omitempty"` //信息的内容
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`     //信息发送者的ID
	MsgContent string `json:"msg_content,omitempty"` //信息的内容
}
