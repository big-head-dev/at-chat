package main

import (
	"fmt"
	"log"
)

// Room object to manage users
type Room struct {
	users    map[string]*User //map of usernames to users
	join     chan *User
	leave    chan *User
	incoming chan ChatMessage
}

// NewRoom creates a Room object and returns the pointer for management of the chatroom
func newRoom() *Room {
	return &Room{
		users:    make(map[string]*User),
		join:     make(chan *User),
		leave:    make(chan *User),
		incoming: make(chan ChatMessage),
	}
}

// goroutine to handle incoming users and incoming messages
func (r Room) start() {
	for {
		select {
		case user := <-r.join:
			r.users[user.username] = user
			log.Println("Room.join - user added", user.username)
			r.broadcastChatMessage(newStatusMessage(fmt.Sprintf("%s has joined.", user.username), r.getUsers()))
		case user := <-r.leave:
			if u, ok := r.users[user.username]; ok {
				u.disconnect()
				delete(r.users, u.username)
				log.Println("Room.leave - user disconnected and removed", user.username)
				r.broadcastChatMessage(newStatusMessage(fmt.Sprintf("%s has left.", u.username), r.getUsers()))
			}
		case message := <-r.incoming:
			r.broadcastChatMessage(message)
		}
	}
}

// sends chat messages to all users
func (r Room) broadcastChatMessage(m ChatMessage) {
	for _, u := range r.users {
		u.outgoing <- m
	}
}

func (r Room) getUsers() []string {
	uns := make([]string, len(r.users))

	i := 0
	for un := range r.users {
		uns[i] = un
		i++
	}
	return uns
}
