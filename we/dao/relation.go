package dao

// Relation 点赞
type Relation struct {
	ID     int64 // 自增主键
	FanId  int64 // 粉丝的用户ID
	UserId int64 // 被关注用户ID
	Status int8  // 点赞状态（1表示关注，2表示取关）
}

// FindRelationsByUserID 根据被关注用户ID查询点赞状态
func FindRelationsByUserID(userID int64) []Relation {
	var relations []Relation
	db.Where(&Relation{UserId: userID, Status: 1}).Find(&relations)
	return relations
}

// FindRelationsByFanID 根据视频ID查询点赞状态
func FindRelationsByFanID(fanID int64) []Relation {
	var relations []Relation
	db.Where(&Relation{FanId: fanID, Status: 1}).Find(&relations)
	return relations
}

// UpdateRelation 更新点赞状态
func UpdateRelation(userID int64, fanID int64, status int8) int64 {
	res := db.Model(&Relation{}).Where("user_id = ? AND fan_id = ?", userID, fanID).Update("Status", status)
	return res.RowsAffected
}

// InsertRelation 插入新的点赞数据
func InsertRelation(userID int64, fanID int64, status int8) error {
	res := db.Create(&Relation{UserId: userID, FanId: fanID, Status: status})
	return res.Error
}
