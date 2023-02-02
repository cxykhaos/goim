package service

import (
	"context"
	"goim/internal/model"
	"goim/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
)

type groupService struct{}

var GroupService = new(groupService)

func (g groupService) CreateGroup(group model.Group) {
	db.M.Collection("group").InsertOne(context.Background(), group)

}

func (g groupService) CreateGroupMember(group model.Relation) {
	db.M.Collection("relation").InsertOne(context.Background(), group)

}

func (g groupService) GetGroup(R string) model.Group {
	var mem model.Group
	db.M.Collection("group").FindOne(context.Background(), bson.D{{"groupId", R}}).Decode(&mem)
	return mem
}

func (g groupService) GetRelationGroup(R string) model.Group {
	var mem model.Group
	db.M.Collection("group").FindOne(context.Background(), bson.D{{"groupId", R}}).Decode(&mem)
	return mem
}

func (g groupService) GetGroupMember(R string) model.Group {
	var mem model.Group
	db.M.Collection("group").FindOne(context.Background(), bson.D{{"groupId", R}}).Decode(&mem)
	return mem
}
