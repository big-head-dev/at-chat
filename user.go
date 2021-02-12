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
	outgoing chan ChatMessage
}

func (u User) start() {
	defer u.disconnect()
	go u.listen()
	for {
		t, p, err := u.ws.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error ", u.username, t, p, err)
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

func (u User) listen() {
	for {
		select {
		case message := <-u.outgoing:
			log.Println(message)
			if err := u.ws.WriteJSON(message); err != nil {
				log.Println("Outgoing message error ", u.username, message, err)
			}
		}
	}
}

// Closes socks connection
func (u User) disconnect() {
	u.ws.Close()
}
