package service

import (
	"fmt"
	"gitlab.com/wushyrussia/go_chat/client/util"
	"regexp"
	"strconv"
)

func selectAndJoinMenu() {
	var isReturn bool
	var inputValue string

	UpdateActiveChatList()

	util.ClearTerminal(AppContext.OperationSystem)
	printMenu()

	processInput(&isReturn, &inputValue)

	if isReturn {
		StartMainMenu()
	}

	number, _ := strconv.Atoi(inputValue)
	JoinRoom(number - 1)
}

func processInput(isReturn *bool, inputValue *string) {
	for {
		*inputValue = startInput()

		*isReturn = isReturnCommand(*inputValue)
		if *isReturn {
			break
		}

		isValidNumber := isValidNumber(*inputValue)
		if isValidNumber {
			break
		}

		fmt.Print("-> Enter the valid chat number to join or 'return' command: ")
	}
}

func startInput() string {
	for {
		select {
		case command := <-AppContext.InputChanel:
			return command
			break
		}
	}
}

func isValidNumber(command string) bool {
	matched, _ := regexp.MatchString(`^[1-9]+$`, command)
	if matched {
		number, _ := strconv.Atoi(command)
		return number-1 < len(AppContext.ActiveChatList)
	}

	return false
}

func isReturnCommand(command string) bool {
	if command == "return" {
		return true
	}

	return false
}

func printMenu() {
	fmt.Println("********* Active chat list ********")
	chatList := AppContext.ActiveChatList

	for index, element := range chatList {
		fmt.Printf("->  %s - %s \n", strconv.Itoa(index+1), element)
	}

	fmt.Println()
	fmt.Print("-> Enter the chat number to join or 'return' to return to the main menu: ")
}
