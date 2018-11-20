package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Message struct {
	Type    int
	Content string
}
type ServiceHandler struct {
}

func (h ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		t, b, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(t)
		fmt.Println(string(b))
		conn.WriteJSON(&Message{
			Type:    t,
			Content: string(b),
		})
	}
}

func main() {
	handler := ServiceHandler{}
	http.ListenAndServe(":9001", handler)
}
