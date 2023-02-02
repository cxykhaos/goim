package service

import (
	"context"
	"goim/internal/model"
	"goim/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
)

type searchService struct{}

var SearchService = new(searchService)

type Thing interface {
	model.User | model.Group
}

func (s *searchService) SearchThing(id string) (model.User, model.Group) {
	var user model.User
	db.M.Collection("user").FindOne(context.TODO(), bson.M{"userId": id}).Decode(&user)

	var group model.Group
	db.M.Collection("group").FindOne(context.TODO(), bson.M{"groupId": id}).Decode(&group)
	return user, group
}
