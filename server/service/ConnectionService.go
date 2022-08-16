package service

import (
	"encoding/json"
	"fmt"
	"gitlab.com/wushyrussia/go_chat/common/dto"
	"net"
)

func EncodeDtoAndSend(connection net.Conn, clientDto dto.Client) {
	encode := json.NewEncoder(connection)
	jsonEncodeError := encode.Encode(clientDto)
	if jsonEncodeError != nil {
		fmt.Printf("Error at EncodeDtoAndSend: %v ", jsonEncodeError)
	}
}

func GetClient(conn net.Conn) (dto.Client, error) {
	decode := json.NewDecoder(conn)

	var requestClient dto.Client
	err := decode.Decode(&requestClient)

	if err != nil {
		return requestClient, err
	}

	return requestClient, nil
}

func RemoveConnection(connectionList []net.Conn, conn net.Conn) []net.Conn {
	var indexForRemove = -1
	for index, element := range connectionList {
		if element == conn {
			indexForRemove = index
		}
	}

	if indexForRemove != -1 {
		return append(connectionList[:indexForRemove], connectionList[indexForRemove+1:]...)
	}
	return nil
}

func StartConnectionHandler(listener *net.Listener) {
	conn, connectionError := (*listener).Accept()

	if connectionError != nil {
		fmt.Printf("Error at StartConnectionHandler: %s \n", connectionError)
	}

	AppContext.ConnectionList = append(AppContext.ConnectionList, conn)

	go ClientHandler(conn)
}
