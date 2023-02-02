package request

type SendMessageRequest struct {
	Content     string `json:content`
	ContentType int64  `json:contentType`
	FromId      string `json:fromId`
	Id          int64  `json:hasBeenSentId`
	CreateTime  int64  `json:createTime`
}
