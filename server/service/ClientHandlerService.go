package service

import (
	"fmt"
	"gitlab.com/wushyrussia/go_chat/common/consts"
	"gitlab.com/wushyrussia/go_chat/common/dto"
	"net"
)

func ClientHandler(conn net.Conn) {
	defer conn.Close()

	for {
		requestClient, err := GetClient(conn)

		if err != nil {
			connections := RemoveConnection(AppContext.ConnectionList, conn)
			if connections != nil {
				AppContext.ConnectionList = connections
			}
			conn.Close()
			break
		}

		clientData := requestClient.ClientData
		serverData := dto.ServerData{ActiveChatMsg: clientData}
		client := dto.Client{ServerData: serverData}

		if clientData.UserName != "" {
			switch clientData.Input {
			case consts.StopChat:
				RemoveUserFromChatContext(client.ServerData.ActiveChatMsg.ActiveChatName, conn)
			case consts.GetActiveChatList:
				sendChatList(conn)
			case consts.JoinRoom:
				joinRoom(client, &conn)
			case consts.DisconnectClient:
				connections := RemoveConnection(AppContext.ConnectionList, conn)
				if connections != nil {
					AppContext.ConnectionList = connections
				}
				conn.Close()
			default:
				publishMsg(client)
			}
		}
	}
}

func sendChatList(connection net.Conn) {
	serverData := dto.ServerData{ServerChatList: GetChatNameList()}
	clientDto := dto.Client{ServerData: serverData}

	EncodeDtoAndSend(connection, clientDto)
}

func joinRoom(data dto.Client, conn *net.Conn) {
	chatName := data.ServerData.ActiveChatMsg.ActiveChatName
	userName := data.ServerData.ActiveChatMsg.UserName
	data.ServerData.ActiveChatMsg.Input = fmt.Sprintf("-> %s connected", userName)
	data.ServerData.ActiveChatMsg.UserName = ""

	AddUserToChatContext(chatName, conn)
	publishMsg(data)
}

func publishMsg(client dto.Client) {
	for _, chat := range AppContext.ChatList {
		if chat.Name == client.ServerData.ActiveChatMsg.ActiveChatName {
			for _, connection := range chat.UserList {
				EncodeDtoAndSend(connection, client)
			}
			break
		}
	}
}
