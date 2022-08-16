package model

type ChatRoom struct {
	RoomName   string
	ChatChanel chan string
	IsActive   bool
}
