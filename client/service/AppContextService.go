package service

import (
	"flag"
	"fmt"
	"gitlab.com/wushyrussia/go_chat/client/model"
	"gitlab.com/wushyrussia/go_chat/client/util"
	"gitlab.com/wushyrussia/go_chat/common/consts"
	"net"
	"os"
)

func InitInputChanel() {
	inputChanel := make(chan string, 1)

	AppContext.InputChanel = inputChanel
}

func InitConnection() {
	serverAddress := flag.String("server", "127.0.0.1:8085", "server address with port")
	flag.Parse()

	conn, connectionError := net.Dial("tcp", *serverAddress)
	if connectionError != nil {
		fmt.Printf("Error at InitConnection: %s", connectionError)
		fmt.Println("application will closed")
		os.Exit(0)
	}

	AppContext.Connection = &conn
}

func InitUserName() {
	fmt.Print("Enter your name: ")

	select {
	case userName := <-AppContext.InputChanel:
		AppContext.UserName = &userName
	}
}

func InitRoomContext(roomName string) {
	chatChanel := make(chan string, 1)

	AppContext.ActiveChat = model.ChatRoom{
		RoomName:   roomName,
		ChatChanel: chatChanel,
		IsActive:   true}

	util.SendInput(consts.JoinRoom, &AppContext)

	serverData := util.GetServerData(&AppContext)

	printWelcomeRoomMsg(roomName)
	fmt.Println(serverData.ServerData.ActiveChatMsg.Input)

}

func ClearRoomContext() {
	AppContext.ActiveChat = model.ChatRoom{}
}

func CloseApp() {
	util.SendInput(consts.DisconnectClient, &AppContext)

	os.Exit(0)
}

func UpdateActiveChatList() {
	util.SendInput(consts.GetActiveChatList, &AppContext)

	serverData := util.GetServerData(&AppContext)
	AppContext.ActiveChatList = serverData.ServerData.ServerChatList
}
