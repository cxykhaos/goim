package service

import (
	"context"
	"fmt"
	"goim/internal/model"
	"goim/pkg/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type userService struct {
}

var UserService = userService{}

func (u *userService) CreateUser(user model.User) {
	db.M.Collection("user").InsertOne(context.Background(), user)

}

func (u *userService) UpdateInfo(user model.User) {
	db.M.Collection("user").UpdateOne(context.Background(), // 获取上下文参数
		bson.D{ // 设置查询条件， 查询item=paper的文档
			{"userId", user.Id},
		},
		bson.D{ // 设置更新表达式
			{"$set", bson.D{ // 使用$set操作符设置更新的字段内容
				{"headImg", user.HeadImg},
				{"updateAt", time.Now().UnixNano()},
			}},
		})
}
func (u *userService) GetUsersById(roomIdentity string, skip, limit *int64) []*model.User {
	var user []*model.User
	cur, _ := db.M.Collection("user").Find(context.TODO(), bson.M{})
	for cur.Next(context.Background()) {
		m := model.User{}
		err := cur.Decode(&m)
		if err != nil {
			continue
		}
		user = append(user, &m)
	}
	return user

}

func (u *userService) UserLogin(id, token string) {
	key := "login:" + id
	db.RC.Set(key, token, 24*time.Hour)
}

func (u *userService) GetUserById(id string) model.User {
	var user model.User
	err := db.M.Collection("user").FindOne(context.TODO(), bson.M{"userId": id}).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user

}

func (u *userService) SearchUsersByGroup(groupId string) {

}

// func (u *userService) UpdateUserById(user *model.User) {
// 	db.M.Collection("user").UpdateOne(context.TODO(), bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"avatar": "213123w"}})

// }
