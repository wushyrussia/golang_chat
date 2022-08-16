package service

import (
	"fmt"
	"gitlab.com/wushyrussia/go_chat/client/model"
	"gitlab.com/wushyrussia/go_chat/client/util"
	"runtime"
)

var AppContext = model.ApplicationContext{}

func ClientRun() {
	util.ClearTerminal(runtime.GOOS)
	AppContext.OperationSystem = runtime.GOOS

	fmt.Println("********* Welcome to Go chat! ********")

	go util.InputData(&AppContext)

	InitInputChanel()
	InitUserName()
	InitConnection()

	StartMainMenu()

	conn := *AppContext.Connection
	defer conn.Close()
}
