package response

type UserInfo struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Signature string `json:"signature"`
	HeadImg   string `json:"headImg"`
	ChatBgImg string `json:"chatBgImg"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	CreateAt  int64  `json:"createAt"`
}
