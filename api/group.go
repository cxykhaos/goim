package api

import (
	"goim/internal/model"
	"goim/internal/response"
	"goim/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRelationGroup(ctx *gin.Context) {
	relation := model.Relation{
		UserId:      "1",
		MessageType: 1,
	}
	relations := service.RelationService.GetAllRelationGroup(relation)
	resp := make([]response.RelationResponse, 0)
	for _, v := range relations {
		group := service.GroupService.GetRelationGroup(v.RelationId)
		relationresp := response.RelationResponse{
			Name:        group.GroupName,
			HeadImgs:    group.HeadImg,
			MessageType: 1,
		}
		resp = append(resp, relationresp)
	}

	ctx.JSON(200, gin.H{
		"data": resp,
	})
}

func CreateGroupMember(ctx *gin.Context) {

	relation := model.Relation{
		UserId:      "1",
		MessageType: 1,
		RelationId:  "63c671f64fccf79056837fe4",
		CreateAt:    time.Now().UnixNano(),
	}
	service.GroupService.CreateGroupMember(relation)
	// relations := service.RelationService.GetRelationGroup(relation)
}
