package main

import (
	"fmt"
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
			fmt.Println("Room.join - user added ", user.username)
			fmt.Println("Now hosting users ", len(r.users))
			r.broadcastStatus(newStatusMessage(fmt.Sprintf("%s has joined.", user.username)))
			r.sendPreviousMessagesToUser(user.username)
		case user := <-r.leave:
			if u, ok := r.users[user.username]; ok {
				u.disconnect()
				delete(r.users, u.username)
				fmt.Println("Room.leave - user disconnected and removed", user.username)
				r.broadcastStatus(newStatusMessage(fmt.Sprintf("%s has left.", u.username)))
			}
		case message := <-r.incoming:
			r.broadcastChatMessage(message)
		}
	}
}

// sends status messages to all users
func (r Room) broadcastStatus(m StatusMessage) {
	for _, u := range r.users {
		u.receiveStatusMessage(m)
	}
}

// sends chat messages to all users
func (r Room) broadcastChatMessage(m ChatMessage) {
	for _, u := range r.users {
		u.receiveChatMessage(m)
	}
}

// TODO: sends saved chat messages to newly connected users
func (r Room) sendPreviousMessagesToUser(username string) {
	fmt.Println("Send previous messages to user", username)
}
