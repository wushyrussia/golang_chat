package service

import (
	"fmt"
	"gitlab.com/wushyrussia/go_chat/client/util"
)

func StartMainMenu() {
	util.ClearTerminal(AppContext.OperationSystem)
	printMainMenu()
	fmt.Print("-> Enter command number: ")
	select {
	case command := <-AppContext.InputChanel:
		mainMenuHelper(command)
	}
}

func mainMenuHelper(command string) {
	switch command {
	case "1":
		selectAndJoinMenu()
	case "2":
		CloseApp()
	default:
		fmt.Println("-> Please, enter valid command number: ")
	}
}

func printMainMenu() {
	fmt.Println("********* Use the command numbers to navigate through the menu ********")
	fmt.Println("")
	fmt.Println("->  1 - Select a chat and join")
	fmt.Println("->  2 - CloseApp")
}
