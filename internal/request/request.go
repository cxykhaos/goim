package request

type GetMessageRequest struct {
	Offset          int64   `json:"offset,omitempty"`
	HasBeenSentId   string  `json:"hasBeenSentId,omitempty"`
	FromId          string  `json:"fromId"`
	ToId            string  `json:"toId"`
	Content         string  `json:"content"`
	ContentType     int64   `json:"contentType"`
	MessageType     int64   `json:"messageType"`
	ContentDuration float64 `json:"contentDuration"`
	Pic             string  `json:"pic"`
	Url             string  `json:"url" `
	CreatedAt       int64   `json:"createAt"`
	UpdatedAt       int64   `json:"updateAt"`
}
