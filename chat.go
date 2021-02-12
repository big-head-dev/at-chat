package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var port = flag.Int("port", 8080, "port to use, defaults to 8080")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	// parse arguments
	flag.Parse()

	//init chat room
	room := newRoom()
	go room.start()

	//serves the chat app website
	http.Handle("/", http.FileServer(http.Dir("./web")))

	//sock requests
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {

		//get username from cookie in request, no reason to upgrade without proper auth
		cookie, err := r.Cookie("atchat-username")
		if err != nil {
			log.Println("No 'atchat-username' cookie found ", err)
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Username required."))
			return
		}

		//upgrade connection
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrader failed ", err)
			return
		}

		// deny already taken usernames
		if _, ok := room.users[cookie.Value]; ok {
			log.Println("Username already connected", cookie.Value)
			conn.WriteJSON(newStatusMessage("A user with that name is already connected.", nil))
			conn.Close()
			return
		}

		//create user
		user := &User{cookie.Value, conn, room, make(chan ChatMessage)}
		go user.start()

		//add to room
		room.join <- user
	})

	//start
	log.Println("Starting on port", *port)
	if err := http.ListenAndServe(fmt.Sprint(":", *port), nil); err != nil {
		log.Fatal("http.ListenAndServe ", err)
	}
}
