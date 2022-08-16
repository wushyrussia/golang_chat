package service

import (
	"fmt"
	"gitlab.com/wushyrussia/go_chat/client/util"
	"gitlab.com/wushyrussia/go_chat/common/consts"
)

const chatMsgFormat = "%s: %s \n"

func JoinRoom(roomIndex int) {
	roomName := AppContext.ActiveChatList[roomIndex]
	InitRoomContext(roomName)

	go chatBodyRunner()

chatCycle:
	for {
		select {
		case senderMsg := <-AppContext.InputChanel:
			util.SendInput(senderMsg, &AppContext)

			if senderMsg == consts.StopChat {
				ClearRoomContext()
				break chatCycle
			}

		case chatMsg := <-AppContext.ActiveChat.ChatChanel:
			if chatMsg != fmt.Sprintf(chatMsgFormat, *AppContext.UserName, AppContext.LatestSendMsg) {
				fmt.Printf(chatMsg)
			}
		}
	}

	StartMainMenu()
}

func printWelcomeRoomMsg(name string) {
	util.ClearTerminal(AppContext.OperationSystem)

	fmt.Printf("-> *** Welcome to %s room *** \n", name)
	fmt.Println("-> *** To leave the room enter 'stop_chat' ", name)
}

func chatBodyRunner() {
	for {
		serverData := util.GetServerData(&AppContext)
		activeChatMsg := serverData.ServerData.ActiveChatMsg
		AppContext.ActiveChat.ChatChanel <- fmt.Sprintf(chatMsgFormat, activeChatMsg.UserName, activeChatMsg.Input)
	}
}
