package response

type Message struct {
	HasBeenSentId   string  `json:"hasBeenSentId,omitempty"`
	FromId          string  `json:"fromId"`
	ToId            string  `json:"toId"`
	Content         string  `json:"content"`     //消息
	ContentType     int64   `json:"contentType"` //消息类型
	MessageType     int64   `json:"messageType"` //发送的类型 1 私聊， 2 群聊
	ContentDuration float64 `json:"contentDuration" bson:"contentDuration"`
	Pic             string  `json:"pic" bson:"pic"`
	Url             string  `json:"url" bson:"url"`
	CreateAt        int64   `json:"createAt"`
}
