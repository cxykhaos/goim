package db

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// RedisClient Redis缓存客户端单例
var RC *redis.Client

//func GetRedis() *redis.Client {
//	return RedisClient
//}

func InitRedis() *redis.Client {
	addr := viper.GetString("db.redis.addr")
	db := viper.GetInt("db.redis.db")
	password := viper.GetString("db.redis.password")
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	RC = client
	return RC
}
func CloseRedis() {
	RC.Close()
}
