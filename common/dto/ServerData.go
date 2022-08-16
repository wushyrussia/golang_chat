package dto

type ServerData struct {
	ServerChatList       []string `json:"serverChatList"`
	ServerActiveUserList []string `json:"serverActiveUserList"`
	ActiveChatMsg        ClientData
}
