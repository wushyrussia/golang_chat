package service

import (
	"flag"
	"fmt"
	"gitlab.com/wushyrussia/go_chat/server/model"
	"net"
	"os"
	"regexp"
	"strings"
)

var AppContext = model.ApplicationContext{}

const (
	portArgument            = "port"
	defaultPortArgument     = "8085"
	portArgumentDescription = "server port"

	roomsArgument            = "rooms"
	defaultRoomsArgument     = "undefined"
	roomsArgumentDescription = "server chat rooms. with use ',' as delimiter for several values"
)

func CmdArgumentsHandler() {
	serverPort := flag.String(portArgument, defaultPortArgument, portArgumentDescription)
	rooms := flag.String(roomsArgument, defaultRoomsArgument, roomsArgumentDescription)

	flag.Parse()

	serverPortInit(*serverPort)

	roomInit(*rooms)
}

func GetChatNameList() []string {
	var nameList []string

	for _, element := range AppContext.ChatList {
		nameList = append(nameList, element.Name)
	}

	return nameList
}

func AddUserToChatContext(chatName string, conn *net.Conn) {
	for index, element := range AppContext.ChatList {
		if element.Name == chatName {
			element.UserList = append(element.UserList, *conn)
			AppContext.ChatList = append(AppContext.ChatList[:index], AppContext.ChatList[index+1:]...)
			AppContext.ChatList = append(AppContext.ChatList, element)
			break
		}
	}
}

func RemoveUserFromChatContext(chatName string, conn net.Conn) {
	for chatIndex, chat := range AppContext.ChatList {
		if chat.Name == chatName {
			for index, connection := range chat.UserList {
				if connection == conn {
					AppContext.ChatList[chatIndex].UserList =
						append(AppContext.ChatList[chatIndex].UserList[:index],
							AppContext.ChatList[chatIndex].UserList[index+1:]...)
					break
				}
			}
		}
	}
}

func roomInit(rooms string) {

	AppContext.ChatList = []model.ChatRoom{
		{Name: "General chat", UserList: []net.Conn{}},
		{Name: "Next gen chat", UserList: []net.Conn{}},
		{Name: "Programmers chat", UserList: []net.Conn{}}}

	if rooms != "undefined" {
		for _, element := range strings.Split(rooms, ",") {
			room := model.ChatRoom{Name: element, UserList: []net.Conn{}}
			AppContext.ChatList = append(AppContext.ChatList, room)
		}
	}
}

func serverPortInit(port string) {
	matched, _ := regexp.MatchString(`^[0-9]+$`, port)
	if matched {
		AppContext.ServerPort = port
	} else {
		fmt.Printf("Incorrect server port: %s . Restart app with correct arguments or with out arguments", port)
		os.Exit(0)
	}
}
