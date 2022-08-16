package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gitlab.com/wushyrussia/go_chat/client/model"
	"gitlab.com/wushyrussia/go_chat/common/dto"
	"os"
	"os/exec"
)

func InputData(applicationContext *model.ApplicationContext) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		enteredText := scanner.Text()
		applicationContext.InputChanel <- enteredText
	}
}

func SendInput(input string, applicationContext *model.ApplicationContext) {
	if input != "" {
		roomName := applicationContext.ActiveChat.RoomName
		clientData := dto.ClientData{UserName: *applicationContext.UserName, Input: input, ActiveChatName: roomName}
		var request = dto.Client{ClientData: clientData}

		encode := json.NewEncoder(*applicationContext.Connection)
		jsonEncodeError := encode.Encode(request)

		applicationContext.LatestSendMsg = input
		if jsonEncodeError != nil {
			fmt.Printf("Error is: %v ", jsonEncodeError)
		}
	}
}

func GetServerData(applicationContext *model.ApplicationContext) dto.Client {
	decode := json.NewDecoder(*applicationContext.Connection)

	var serverData dto.Client

	err := decode.Decode(&serverData)
	if err != nil {
		fmt.Printf("Error at GetServerData: %v \n", err)
	}

	return serverData
}

func ClearTerminal(operationSystem string) {
	switch operationSystem {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
