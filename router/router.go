package router

import (
	"goim/api"
	"goim/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	socket := RunSocket
	r := gin.Default()
	r.Use(Cors())
	user := r.Group("user")
	{
		user.POST("login", api.UserLogin)
		user.GET("info", middlewares.AuthMiddleware(), api.GetUserInfo)
		user.GET("otherInfo", middlewares.AuthMiddleware(), api.GetOtherUserInfo)
		user.POST("register", api.UserRegister)
		user.GET("fromMessage", middlewares.AuthMiddleware(), api.MemberFromChat)
		// user.GET("friend/list", api.GetUserFriendsAndMsg)
	}
	relation := r.Group("relation")
	{
		relation.POST("add", middlewares.AuthMiddleware(), api.AddNewFriend)
		// relation.GET("list", api.GetRelationGroup)
		relation.GET("list", middlewares.AuthMiddleware(), api.GetRelation)
	}
	message := r.Group("message")
	{
		message.GET("list", middlewares.AuthMiddleware(), api.GetMessage)
	}
	search := r.Group("search")
	{
		search.GET("list", middlewares.AuthMiddleware(), api.Search)
	}
	file := r.Group("file")
	{
		file.Static("static", "files/")
		file.GET("get/:fileName", api.GetFile)
		file.POST("upload", middlewares.AuthMiddleware(), api.SaveFile)
	}
	r.GET("/socket", middlewares.AuthMiddleware(), socket)
	return r

}
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
