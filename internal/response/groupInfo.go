package response

type GroupInfo struct {
	Id           string     `json:"id" form:"id" binding:"required" bson:"groupId,omitempty"`
	GroupName    string     `json:"groupName" bson:"groupName"`
	HeadImg      string     `json:"headImg" bson:"headImg"`
	OwnId        string     `json:"ownId" bson:"ownId"`
	Type         int64      `json:"type" bson:"type"`
	CreateAt     int64      `json:"createAt" bson:"createAt"`
	GroupMembers []UserInfo `json:"groupMembers"`
}
