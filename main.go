package main

import (
	"goim/internal/common"
	"goim/internal/server"
	"goim/pkg/db"
	"goim/pkg/log"
	"goim/router"
	"os"

	"github.com/spf13/viper"
)

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	db.InitMongo()
	db.InitRedis()
}

func main() {
	log.InitLogger("logs/chat.log", "debug")
	defer db.CloseMongo()
	defer db.CloseRedis()
	common.StaticPath = viper.GetString("file.savepath")
	port := viper.GetString("server.port")
	r := router.SetupRouter()
	go server.MyServer.Start()
	panic(r.Run(":" + port))

}
