package main

import (
	"github.com/gorilla/websocket"
)

// User object to manage connection
type User struct {
	username string
	ws       *websocket.Conn
}

func (u User) disconnect() {
	u.ws.Close()
}
