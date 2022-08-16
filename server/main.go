package main

import (
	"gitlab.com/wushyrussia/go_chat/server/service"
	"net"
)
import "fmt"

func main() {
	service.CmdArgumentsHandler()
	fmt.Println("Server started")
	printActiveChat()

	listener, _ := net.Listen("tcp", fmt.Sprintf(":%s", service.AppContext.ServerPort))

	for {
		service.StartConnectionHandler(&listener)
	}
}

func printActiveChat() {
	fmt.Print("Active chat:")
	for _, element := range service.AppContext.ChatList {
		fmt.Printf("'%s' ", element.Name)
	}
	fmt.Println()
}
