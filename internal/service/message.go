package service

import (
	"context"
	"goim/internal/model"
	"goim/pkg/db"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type messageService struct{}

var MessageService = new(messageService)

func (m *messageService) GetMessagesByPrivateOne(FromId, ToId string) model.Message {
	sort := bson.D{{"createAt", -1}}
	O := options.Find()
	O.SetSort(sort)
	O.SetLimit(1)
	var F bson.M
	F = bson.M{"$or": []bson.M{bson.M{"fromId": FromId, "toId": ToId}, bson.M{"fromId": ToId, "toId": FromId}}}
	cur, _ := db.M.Collection("message").Find(context.TODO(), F, O)
	for cur.Next(context.Background()) {
		m := model.Message{}
		cur.Decode(&m)
		return m
	}
	return model.Message{}
}

func (m *messageService) GetMessagesByGroupOne(ToId string) model.Message {
	sort := bson.D{{"createAt", -1}}
	O := options.Find()
	O.SetSort(sort)
	O.SetLimit(1)
	var F bson.M
	F = bson.M{"messageType": 2, "toId": ToId}
	cur, _ := db.M.Collection("message").Find(context.TODO(), F, O)
	for cur.Next(context.Background()) {
		m := model.Message{}
		cur.Decode(&m)
		return m
	}
	return model.Message{}
}
func (m *messageService) GetMessagesByGroup(ToId string, offset, messagetype int64) []*model.Message {
	var messages []*model.Message
	sort := bson.D{{"createAt", -1}}
	O := options.Find()
	O.SetSort(sort)
	O.SetLimit(30)
	O.SetSkip(offset)
	var F bson.M
	F = bson.M{"messageType": 2, "toId": ToId}
	cur, _ := db.M.Collection("message").Find(context.TODO(), F, O)
	for cur.Next(context.Background()) {
		m := model.Message{}
		err := cur.Decode(&m)
		if err != nil {
			continue
		}
		messages = append(messages, &m)
	}
	return messages
}

func (m *messageService) GetMessagesByPrivate(FromId, ToId string, offset, messagetype int64) []*model.Message {
	var messages []*model.Message
	sort := bson.D{{"createAt", -1}}
	O := options.Find()
	O.SetSort(sort)
	O.SetLimit(30)
	O.SetSkip(offset)
	var F bson.M
	F = bson.M{"messageType": messagetype, "$or": []bson.M{bson.M{"fromId": FromId, "toId": ToId}, bson.M{"fromId": ToId, "toId": FromId}}}
	cur, _ := db.M.Collection("message").Find(context.TODO(), F, O)
	for cur.Next(context.Background()) {
		m := model.Message{}
		err := cur.Decode(&m)
		if err != nil {
			continue
		}
		messages = append(messages, &m)
	}
	return messages
}

func (m *messageService) SendMessage(message model.Message) {
	message.HasBeenSentId = string(uuid.New().String())
	db.M.Collection("message").InsertOne(context.Background(), message)

}
