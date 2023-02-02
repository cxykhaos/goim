package api

import (
	"fmt"
	"goim/internal/model"
	"goim/internal/response"
	"goim/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRelation(ctx *gin.Context) {
	user := ctx.MustGet("user").(model.User)
	relation := model.Relation{
		UserId: user.Id,
	}
	relations := service.RelationService.GetAllRelationGroup(relation)
	resp := make([]response.RelationResponse, 0)
	for _, v := range relations {
		var relationresp response.RelationResponse
		if v.MessageType == 2 {
			group := service.GroupService.GetRelationGroup(v.RelationId)
			relationresp = response.RelationResponse{
				Name:        group.GroupName,
				HeadImgs:    group.HeadImg,
				RelationId:  group.Id,
				MessageType: 2,
				NewMsg:      service.MessageService.GetMessagesByGroupOne(v.RelationId).ChangeToDto(),
			}
		} else {
			// } else if v.ChatType == 0 {
			userFriend := service.UserService.GetUserById(v.RelationId)
			relationresp = response.RelationResponse{
				Name:        userFriend.Username,
				HeadImgs:    userFriend.HeadImg,
				RelationId:  userFriend.Id,
				MessageType: 1,
				NewMsg:      service.MessageService.GetMessagesByPrivateOne(relation.UserId, v.RelationId).ChangeToDto(),
			}
		}
		resp = append(resp, relationresp)
	}

	response.Response(ctx, 200, resp, "")
}

func AddNewFriend(ctx *gin.Context) {
	user := ctx.MustGet("user").(model.User)
	var temp model.Relation
	ctx.ShouldBindJSON(&temp)
	temp.UserId = user.Id
	if temp.RelationId == user.Id {
		response.Response(ctx, 400, nil, "不可添加自己为好友")

	}
	temp.CreateAt = time.Now().UnixNano()
	flag := service.RelationService.AddUserFriend(temp)
	if !flag {
		response.Response(ctx, 400, nil, "已添加该好友")
		return
	}
	var systemMsg model.Message
	fmt.Println(temp.MessageType)
	if temp.MessageType == 2 {
		systemMsg = model.Message{CreateAt: time.Now().UnixNano(), ToId: temp.RelationId, Content: user.Id + "已加入群聊", ContentType: 0, MessageType: 2}
		service.MessageService.SendMessage(systemMsg)
		fmt.Println("fasong")
	}

	response.Response(ctx, 200, nil, "success")

}
