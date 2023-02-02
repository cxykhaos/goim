package router

import (
	"goim/internal/model"
	"goim/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunSocket(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &server.Client{
		Name: user.Id,
		Conn: ws,
		Send: make(chan model.Message, 100),
	}
	server.MyServer.Register <- client
	go client.Read()
	go client.Write()
}
