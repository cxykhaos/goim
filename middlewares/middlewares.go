package middlewares

import (
	"goim/internal/service"
	"goim/pkg/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		GetUserInfo(ctx)
	}
}
func GetToken(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("token")
	if token == "" {
		token = ctx.Query("token")
	}
	return token
}

func GetUserInfo(ctx *gin.Context) {
	token := GetToken(ctx)
	if userId, ok := CheckUserStatus(token); ok {
		user := service.UserService.GetUserById(userId)
		ctx.Set("user", user)
		ctx.Next()
		return
	}
	ctx.JSON(http.StatusUnauthorized, "权限不够1")
	ctx.Abort()

}

func CheckUserStatus(token string) (string, bool) {
	sp := strings.Split(token, "-")
	if len(sp) < 2 {
		return "", false
	}
	key := "login:" + sp[0]
	res, _ := db.RC.Get(key).Result()
	return sp[0], res == sp[1]
}
