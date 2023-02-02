package response

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(200, gin.H{"code": code, "data": data, "msg": msg})
}
func Success(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(200, gin.H{"code": code, "data": data, "msg": msg})
}

func Fail(ctx *gin.Context, code int, msg string) {
	ctx.JSON(200, gin.H{"code": code, "data": nil, "msg": msg})
}
