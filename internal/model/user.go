package model

import "goim/internal/response"

type User struct {
	Id        string `json:"id" binding:"required" bson:"userId,omitempty"`
	Username  string `json:"username" bson:"userName" binding:"required"`
	Signature string `json:"signature" bson:"signature" `
	HeadImg   string `json:"headImg" bson:"headImg" `
	ChatBgImg string `json:"chatBgImg" bson:"chatBgImg" `
	Address   string `json:"address" bson:"address" `
	Email     string `json:"email" bson:"email" `
	Password  string `json:"password" bson:"passWord" binding:"required"`
	CreateAt  int64  `json:"createAt" bson:"createAt"`
	UpdateAt  int64  `json:"updateAt" bson:"updateAt" `
	DeleteAt  int64  `json:"deleteAt" bson:"deleteAt" `
}

func (u User) ChangeToDto() response.UserInfo {
	return response.UserInfo{Id: u.Id, Username: u.Username, HeadImg: u.HeadImg, Address: u.Address}
}
