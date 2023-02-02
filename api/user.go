package api

import (
	"fmt"
	"goim/internal/common"
	"goim/internal/model"
	"goim/internal/response"
	"goim/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

type login struct {
	UserId   string `json:"userId"`
	PassWord string `json:"passWord"`
}

func UserLogin(ctx *gin.Context) {
	var lp login
	ctx.ShouldBindJSON(&lp)
	user := service.UserService.GetUserById(lp.UserId)
	token := common.RandString(256)
	service.UserService.UserLogin(user.Id, token)
	ctx.JSON(200, gin.H{"token": user.Id + "-" + token, "userId": user.Id})

}

func UserRegister(ctx *gin.Context) {
	var lp login
	ctx.ShouldBindJSON(&lp)
	user := service.UserService.GetUserById(lp.UserId)
	if user.CreateAt != 0 {
		response.Fail(ctx, 400, "账号名已被使用")
		return
	}
	newuser := model.User{HeadImg: "avatar/khaos.png",
		Address:   "暂未填写",
		Email:     "",
		Username:  common.RandString(5),
		Signature: "暂无签名",
		ChatBgImg: "background/Ran.jpg",
		Id:        lp.UserId, Password: lp.PassWord,
		CreateAt: time.Now().UnixNano(),
	}
	service.UserService.CreateUser(newuser)
	response.Success(ctx, 200, user, "success")

}

type GetInfoRequest struct {
	Type int64  `json:"type"`
	ToId string `json:"toId"`
}

func GetOtherUserInfo(ctx *gin.Context) {
	var getInfoRequest GetInfoRequest
	ctx.ShouldBindQuery(&getInfoRequest)
	if getInfoRequest.Type == 1 {
		Info := service.UserService.GetUserById(getInfoRequest.ToId).ChangeToDto()
		response.Response(ctx, 200, Info, "")
		return
	} else {
		var groupMember []response.UserInfo

		groupInfo := service.GroupService.GetGroup(getInfoRequest.ToId).ChangeToDto()

		r := service.RelationService.GetGroupAllMember(getInfoRequest.ToId)
		for _, v := range r {
			user := service.UserService.GetUserById(v.UserId).ChangeToDto()
			groupMember = append(groupMember, user)
		}
		groupInfo.GroupMembers = groupMember
		fmt.Println(groupInfo)
		response.Response(ctx, 200, groupInfo, "")

	}

}
func GetUserInfo(ctx *gin.Context) {
	user := ctx.MustGet("user").(model.User)
	response.Response(ctx, 200, user, "")

}

func Search(ctx *gin.Context) {
	searchWord := ctx.Query("searchWord")
	user, group := service.SearchService.SearchThing(searchWord)
	fmt.Println(user)
	fmt.Println(group)
	var relationresp []response.RelationResponse
	if user.CreateAt != 0 {
		relationresp = append(relationresp, response.RelationResponse{
			Name:        user.Username,
			HeadImgs:    user.HeadImg,
			RelationId:  user.Id,
			MessageType: 1,
		})
	}
	if group.CreateAt != 0 {

		relationresp = append(relationresp, response.RelationResponse{
			Name:        group.GroupName,
			HeadImgs:    group.HeadImg,
			RelationId:  group.Id,
			MessageType: 2,
		})
	}

	response.Response(ctx, 200, relationresp, "")
}
