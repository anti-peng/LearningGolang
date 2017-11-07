package toy1

import (
	"fmt"
	"net/http"

	wsocket "github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

func Demo1WSServer() {
	http.Handle("/", websocket.Handler(echoServer))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func echoServer(ws *websocket.Conn) {
	for {
		var reply string

		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Receive: " + reply)

		msg := "Received: " + reply + ", thanks."
		fmt.Println("Send: " + msg)

		if err := websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Cant send")
			break
		}
	}
}

// Demo2WSServer works with gorilla/websocket
func Demo2WSServer() {

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Receive: " + string(msg))

		msgTobeSent := "Server response: [" + string(msg) + "] received"
		if err := conn.WriteMessage(msgType, []byte(msgTobeSent)); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
