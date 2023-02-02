package model

type Relation struct {
	UserId      string `json:"userId" bson:"userId" `
	MessageType int64  `json:"messageType" bson:"messageType" `
	CreateAt    int64  `json:"createAt" bson:"createAt"`
	UpdateAt    int64  `json:"updateAt" bson:"updateAt"`
	DeleteAt    int64  `json:"deleteAt" bson:"deleteAt"`
	RelationId  string `json:"relationId" bson:"relationId"`
}
