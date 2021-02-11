package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// User object to manage connection
type User struct {
	username string
	ws       *websocket.Conn
	r        *Room
}

func (u User) start() {
	for {
		t, p, err := u.ws.ReadMessage()
		log.Println("ReadMessage ", u.username, t, p, err)
		if err != nil {
			log.Println("ReadMessage error, disconnecting user")
			u.r.leave <- &u
			return
		}

		if t == websocket.TextMessage {
			message := string(p)
			cm := newChatMessage(u.username, message)
			// send message to room for broadcasting
			u.r.incoming <- cm
		}
	}
}

// Closes socks connection
func (u User) disconnect() {
	u.ws.Close()
}

func (u User) receiveStatusMessage(m StatusMessage) {
	if err := u.ws.WriteJSON(m); err != nil {
		log.Println("User.receiveStatusMessage error ", u.username, err)
	}
}

func (u User) receiveChatMessage(m ChatMessage) {
	if err := u.ws.WriteJSON(m); err != nil {
		log.Println("User.receiveChatMessage error ", u.username, err)
	}
}
