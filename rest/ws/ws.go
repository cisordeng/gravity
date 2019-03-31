package ws

import (
	"fmt"
	"github.com/cisordeng/beego/xenon"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Ws struct {
	xenon.RestResource
	websocket.Conn
}

func init () {
	xenon.RegisterResource(new(Ws))
}

func (this *Ws) Resource() string {
	return "ws.ws"
}

func (this *Ws) Params() map[string][]string {
	return map[string][]string{
		"GET": []string{"name"},
	}
}

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *Ws) Get() {
	name := this.GetString("name")
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer ws.Close()

	var clients = map[string]bool{}
	clients[name] = true
	/*
		for {
			time.Sleep(time.Second * 3)
			msg := models.Message{Message: "这是向页面发送的数据" + time.Now().Format("2006-01-02 15:04:05")}
			broadcast <- msg
		}
	*/
	for {
		time.Sleep(time.Second * 3)
		msg := xenon.Map{} // Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("页面可能断开啦 ws.ReadJSON error: %v", err)
			delete(clients, name)
			break
		} else {
			fmt.Println("接受到从页面上反馈回来的信息 ", msg)
			err := ws.WriteJSON(xenon.Map{
				"a": 1,
				"b": 2,
			})
			if err != nil {
				log.Printf("发送失败 %v", err)
			}
		}

		//broadcast <- msg
	}

	this.ReturnJSON(xenon.Map{})
}
