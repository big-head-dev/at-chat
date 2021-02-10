package main

import "fmt"

// Room object to manage users
type Room struct {
	users map[string]*User //map of usernames to users
	join  chan *User
	leave chan *User
}

// NewRoom creates a Room object and returns the pointer for management of the chatroom
func newRoom() *Room {
	return &Room{
		users: make(map[string]*User),
		join:  make(chan *User),
		leave: make(chan *User),
	}
}

func (r Room) start() {
	for {
		select {
		case user := <-r.join:
			r.users[user.username] = user
			fmt.Println("Room.join - user added ", user.username)
			fmt.Println("Now hosting users ", len(r.users))
			// broadcast join message
			r.broadcast(newStatusMessage(fmt.Sprintf("%s has joined.", user.username)))
			//TODO: write previous x messages for new user
		case user := <-r.leave:
			if u, ok := r.users[user.username]; ok {
				u.disconnect()
				delete(r.users, u.username)
				fmt.Println("Room.leave - user disconnected and removed", user.username)
				// broadcast leave message
				r.broadcast(newStatusMessage(fmt.Sprintf("%s has left.", u.username)))
			}
		}
	}
}

func (r Room) broadcast(m interface{}) {
	fmt.Println(m)
}
