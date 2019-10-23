package controllers

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Message struct {
	Message string `json:"message"`
}

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan interface{})
)

func init() {
	go handleMessages()
}

//广播发送至页面
func handleMessages() {
	fmt.Println("===")
	for {
		fmt.Println("clients len(broadcast)", len(broadcast))
		msg := <-broadcast
		fmt.Println("clients msg ", msg)
		for client := range clients {
			err := client.WriteJSON(msg)
			fmt.Println("clients err ", err)
			// if err != nil {
			// 	log.Printf("client.WriteJSON error: %v", err)
			// 	client.Close()
			// 	delete(clients, client)
			// }
		}
	}
}
