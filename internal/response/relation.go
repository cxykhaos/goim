package response

type RelationResponse struct {
	RelationId  string  `json:"relationId"`
	Name        string  `json:"name"`
	HeadImgs    string  `json:"headImg"`
	MessageType int64   `json:"messageType"`
	CreateAt    int64   `json:"createAt" bson:"createAt"`
	UpdateAt    int64   `json:"updateAt" bson:"updateAt"`
	NewMsg      Message `json:"newMsg"`
}
