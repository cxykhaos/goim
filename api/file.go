package api

import (
	"goim/internal/common"
	"goim/internal/model"
	"goim/internal/service"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetFile(c *gin.Context) {
	fileName := c.Param("name")
	data, _ := ioutil.ReadFile(common.StaticPath + fileName)
	c.Writer.Write(data)
}
func SaveFile(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	namePreffix := uuid.New().String()
	m, _ := c.MultipartForm()
	a := c.DefaultPostForm("type", "chat")
	files := m.File["file"]
	fileName := files[0].Filename
	index := strings.LastIndex(fileName, ".")
	suffix := fileName[index:]
	newFileName := a + "/" + namePreffix + suffix
	path := common.StaticPath + newFileName
	c.SaveUploadedFile(files[0], path)
	if a == "avatar" {
		user.HeadImg = newFileName
		service.UserService.UpdateInfo(user)
	}
	c.JSON(http.StatusOK, gin.H{"data": newFileName})
}
