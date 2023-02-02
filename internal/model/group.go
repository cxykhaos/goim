package model

import "goim/internal/response"

type Group struct {
	Id        string `json:"id" form:"id" binding:"required" bson:"groupId,omitempty"`
	GroupName string `json:"groupName" bson:"groupName"`
	HeadImg   string `json:"headImg" bson:"headImg"`
	OwnId     string `json:"ownid" bson:"ownId"`
	Type      int64  `json:"type" bson:"type"`
	CreateAt  int64  `json:"createAt" bson:"createAt"`
	UpdateAt  int64  `json:"updateAt" bson:"updateAt"`
	DeleteAt  int64  `json:"deleteAt" bson:"deleteAt"`
}

//db.group.insertOne({"groupId":"1","groupName":"我们的群聊","headImg":"avatar/hg.jpg","ownId":"1","type":1,"createAt":1674973643552670000})

//db.relation.insertOne({"userId":"1","messageType":2,"relationId":"1","createAt":1674984369376252400})
//db.relation.insertOne({"userId":"2","messageType":2,"relationId":"1","createAt":1674984369376252400})
//db.relation.insertOne({"userId":"3","messageType":2,"relationId":"1","createAt":1674984369376252400})

func (g Group) ChangeToDto() response.GroupInfo {
	return response.GroupInfo{Id: g.Id, GroupName: g.GroupName, HeadImg: g.HeadImg, OwnId: g.OwnId}
}
