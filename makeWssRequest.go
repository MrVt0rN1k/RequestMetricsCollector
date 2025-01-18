package main

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

func makeWssRequest(url string, jsonData string, wg *sync.WaitGroup) {
	defer wg.Done()
	dialer := websocket.DefaultDialer

	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Error on connection:", err)
	}

	err = conn.WriteJSON(jsonData)
	if err != nil {
		fmt.Println("Error on send request:", err)
		return
	}

	err = conn.ReadJSON(&jsonData)
	if err != nil {
		fmt.Println("Error on read answer:", err)
		return
	}
	fmt.Printf("%+v\n", jsonData)
	defer conn.Close()
}
