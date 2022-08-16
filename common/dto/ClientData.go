package dto

type ClientData struct {
	UserName       string `json:"name"`
	Input          string `json:"input"`
	ActiveChatName string `json:"userActiveChat"`
}
