package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// User object to manage connection
type User struct {
	username string
	ws       *websocket.Conn
}

// Closes socks connection
func (u User) disconnect() {
	u.ws.Close()
}

func (u User) receiveStatusMessage(m StatusMessage) {
	if err := u.ws.WriteJSON(m); err != nil {
		log.Panicln("receiveStatusMessage user", u.username, err)
	}
}

func (u User) receiveChatMessage(m ChatMessage) {
	if err := u.ws.WriteJSON(m); err != nil {
		log.Panicln("receiveChatMessage user", u.username, err)
	}
}
