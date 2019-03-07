package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"mango/business"
	"mango/rest"
	"net/http"
	"time"
)

type Ws struct {
	rest.RestResource
	websocket.Conn
}

func init () {
	rest.Resources = append(rest.Resources, new(Ws))
}

func (this *Ws) Resource() string {
	return "ws.ws"
}

func (this *Ws) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{},
	}
}

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = map[*websocket.Conn]bool{}

type Message struct {
	message string
	sent_at time.Time
}

func (this *Ws) Get() {
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer ws.Close()
	clients[ws] = true
	/*
		for {
			time.Sleep(time.Second * 3)
			msg := models.Message{Message: "这是向页面发送的数据" + time.Now().Format("2006-01-02 15:04:05")}
			broadcast <- msg
		}
	*/
	for {
		time.Sleep(time.Second * 3)

		var msg Message // Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("页面可能断开啦 ws.ReadJSON error: %v", err)
			delete(clients, ws)
			break
		} else {
			fmt.Println("接受到从页面上反馈回来的信息 ", msg.message)
		}
	}
	this.ReturnJSON(business.Map{}, nil, rest.BusinessError{})
}