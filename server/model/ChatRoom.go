package model

import "net"

type ChatRoom struct {
	Name     string
	UserList []net.Conn
}
