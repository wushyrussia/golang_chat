package model

import "net"

type ApplicationContext struct {
	ChatList       []ChatRoom
	ConnectionList []net.Conn
	ServerPort     string
}
