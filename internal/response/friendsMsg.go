package response

type FirendsMsg struct {
	Id       string  `json:"id" form:"id" binding:"required" bson:"_id,omitempty"`
	Username string  `json:"username" form:"userName" binding:"required"`
	HeadImg  string  `json:"headImg"`
	Address  string  `json:"address"`
	Email    string  `json:"email"`
	CreateAt int64   `json:"createAt" bson:"createAt"`
	UpdateAt int64   `json:"updateAt" bson:"updateAt"`
	NewMsg   Message `json:"newMsg"`
}
