package service

import (
	"context"
	"fmt"
	"goim/internal/model"
	"goim/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type relationService struct{}

var RelationService = new(relationService)

func (g relationService) AddUserFriend(relation model.Relation) bool {
	var mem model.Relation
	db.M.Collection("relation").FindOne(context.Background(), bson.M{"userId": relation.UserId, "relationId": relation.RelationId, "messageType": relation.MessageType}).Decode(&mem)
	fmt.Println(mem)
	if mem.CreateAt != 0 {
		return false
	}
	db.M.Collection("relation").InsertOne(context.Background(), relation)
	return true
}
func (g relationService) GetAllRelationGroup(member model.Relation) []*model.Relation {
	var relations []*model.Relation
	sort := bson.D{{"createAt", -1}}
	O := options.Find()
	O.SetSort(sort)
	O.SetLimit(30)
	F := bson.M{"userId": member.UserId}
	cur, _ := db.M.Collection("relation").Find(context.TODO(), F, O)
	for cur.Next(context.Background()) {
		m := model.Relation{}
		err := cur.Decode(&m)
		if err != nil {
			continue
		}
		relations = append(relations, &m)
	}
	return relations

}

func (g relationService) GetGroupAllMember(str string) []*model.Relation {
	var relations []*model.Relation
	O := options.Find()
	F := bson.M{"relationId": str, "messageType": 2}
	cur, _ := db.M.Collection("relation").Find(context.TODO(), F, O)
	for cur.Next(context.Background()) {
		m := model.Relation{}
		err := cur.Decode(&m)
		if err != nil {
			continue
		}
		relations = append(relations, &m)
	}
	return relations

}
