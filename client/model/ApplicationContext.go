package model

import "net"

type ApplicationContext struct {
	UserName        *string
	Connection      *net.Conn
	ActiveChat      ChatRoom
	InputChanel     chan string
	LatestSendMsg   string
	ActiveChatList  []string
	OperationSystem string
	ServerAddress   string
}
