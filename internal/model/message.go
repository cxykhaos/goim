package model

import "goim/internal/response"

type Message struct {
	HasBeenSentId   string  `json:"hasBeenSentId,omitempty" bson:"messageId"`
	FromId          string  `json:"fromId" bson:"fromId"`
	ToId            string  `json:"toId" bson:"toId"`
	Content         string  `json:"content" bson:"content"`         //消息
	ContentType     int64   `json:"contentType" bson:"contentType"` //消息类型
	MessageType     int64   `json:"messageType" bson:"messageType"` //发送的类型 1 私聊， 2 群聊
	ContentDuration float64 `json:"contentDuration" bson:"contentDuration"`
	Pic             string  `json:"pic" bson:"pic"`
	Url             string  `json:"url" bson:"url"`
	CreateAt        int64   `json:"createAt" bson:"createAt"`
	UpdateAt        int64   `json:"updateAt" bson:"updateAt"`
}

func (m Message) ChangeToDto() response.Message {
	return response.Message{HasBeenSentId: m.HasBeenSentId, FromId: m.FromId, ToId: m.ToId, Content: m.Content, ContentType: m.ContentType, MessageType: m.MessageType, ContentDuration: m.ContentDuration, CreateAt: m.CreateAt}

}
