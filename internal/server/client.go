package server

import (
	"fmt"
	"goim/internal/model"
	"goim/pkg/log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan model.Message
}

func (c *Client) Read() {
	defer func() {
		MyServer.Ungister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		msg := model.Message{}
		err := c.Conn.ReadJSON(&msg)
		msg.CreateAt = time.Now().UnixNano()
		if err != nil {
			log.Logger.Error("client read message error " + err.Error())
			MyServer.Ungister <- c
			c.Conn.Close()
			break
		}

		MyServer.Broadcast <- msg

	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		go c.setSend(message)
	}
}
func (c *Client) setSend(message model.Message) {
	// time.Sleep(8 * time.Second)
	fmt.Printf("%s, 发了条信息, 通知您：%s\n", message.FromId, c.Name)
	c.Conn.WriteJSON(message)

}
