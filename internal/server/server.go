package server

import (
	"fmt"
	"goim/internal/model"
	"goim/internal/service"
	"goim/pkg/log"
	"sync"
)

var MyServer = NewServer()

type Server struct {
	Clients   map[string]*Client
	mutex     *sync.Mutex
	Broadcast chan model.Message
	Register  chan *Client
	Ungister  chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:     &sync.Mutex{},
		Clients:   make(map[string]*Client),
		Broadcast: make(chan model.Message, 100),
		Register:  make(chan *Client),
		Ungister:  make(chan *Client),
	}
}

func (s *Server) Start() {
	log.Logger.Info("start server", log.Any("start server", "start server..."))
	for {
		select {
		case conn := <-s.Register:
			log.Logger.Info("user: " + conn.Name + " login")
			s.Clients[conn.Name] = conn

		case conn := <-s.Ungister:
			if _, ok := s.Clients[conn.Name]; ok {
				log.Logger.Info(conn.Name + " loginout")
				close(conn.Send)
				delete(s.Clients, conn.Name)
			}
		case message := <-s.Broadcast:
			go s.SolveBroadcast(message)
		}

	}
}

func (s *Server) SolveBroadcast(message model.Message) {
	fmt.Printf("%s %v send to %s\n", message.FromId, message, message.ToId)
	service.MessageService.SendMessage(message)
	if message.MessageType == 1 {
		go s.SendMsgToChannel(message, message.ToId)
	} else if message.MessageType == 2 {
		relationgroup := service.RelationService.GetGroupAllMember(message.ToId)
		for _, v := range relationgroup {
			if v.UserId == message.FromId {
				continue
			}
			go s.SendMsgToChannel(message, v.UserId)
		}
	}
}
func (s *Server) SendMsgToChannel(message model.Message, ToId string) {
	client, ok := s.Clients[ToId]
	if ok {
		client.Send <- message
	}

}
