package api

import (
	"goim/internal/model"
	"goim/internal/request"
	"goim/internal/response"
	"goim/internal/service"

	"github.com/gin-gonic/gin"
)

func GetMessage(ctx *gin.Context) {
	var p request.GetMessageRequest
	ctx.BindQuery(&p)
	var messages []*model.Message
	user := ctx.MustGet("user").(model.User)
	p.FromId = user.Id
	if p.MessageType == 1 {
		messages = service.MessageService.GetMessagesByPrivate(p.FromId, p.ToId, p.Offset, p.MessageType)
	} else if p.MessageType == 2 {
		messages = service.MessageService.GetMessagesByGroup(p.ToId, p.Offset, p.MessageType)
	}
	response.Response(ctx, 200, messages, "")
}

func MemberFromChat(ctx *gin.Context) {
	var p request.GetMessageRequest
	ctx.BindQuery(&p)
	var Users []response.UserInfo
	if p.MessageType == 1 {
		user := service.UserService.GetUserById(p.ToId).ChangeToDto()
		Users = append(Users, user)
	} else {
		r := service.RelationService.GetGroupAllMember(p.ToId)
		for _, v := range r {
			user := service.UserService.GetUserById(v.UserId).ChangeToDto()
			Users = append(Users, user)
		}
	}

	response.Response(ctx, 200, Users, "")

}
