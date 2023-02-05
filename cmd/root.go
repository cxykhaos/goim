package cmd

import (
	"fmt"
	"goim/internal/common"
	"goim/internal/server"
	"goim/pkg/db"
	"goim/pkg/log"
	"goim/router"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "goim",
	Short: "Run the goim server",
	Run: func(cmd *cobra.Command, args []string) {
		log.InitLogger("logs/chat.log", "debug")
		defer db.CloseMongo()
		defer db.CloseRedis()
		common.StaticPath = viper.GetString("file.savepath")
		port := viper.GetString("server.port")
		r := router.SetupRouter()
		go server.MyServer.Start()
		panic(r.Run(":" + port))

	},
}

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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
