package controllers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/websocket"

	"github.com/astaxie/beego"
)

type MyWebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Messages struct {
	Apiname string            `json:"apiname,omitempty"`
	Apidata map[string]string `json:"apidata,omitempty"`
}

// type Message struct {
// 	Message string `json:"message"`
// }

func (c *MyWebSocketController) Websocketmsg() {
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	clients[ws] = true
	//defer ws.Close()
	for {
		time.Sleep(time.Second * 3)
		var msg Usersapi // Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		fmt.Println(err)
		if err != nil {
			log.Printf("页面可能断开啦 ws.ReadJSON error: %v", err)
			delete(clients, ws)
			break
		} else {
			fmt.Println("接受到从页面上反馈回来的信息 ", msg)
		}
		for key, v := range msg {
			Funcuser := reflect.ValueOf(v)
			methodValue := Funcuser.MethodByName(key).Call(nil)
			broadcast <- methodValue[0].Interface()
		}
	}
}
