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
			//TODO: send better denied response??
			return
		}

		//upgrade connection
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Panicln("Upgrader failed ", err)
			return
		}
		fmt.Println("User connected ", conn.RemoteAddr())

		//create user
		user := &User{cookie.Value, conn}

		//add to room
		room.join <- user
	})

	//start
	log.Println("Starting on port", *port)
	if err := http.ListenAndServe(fmt.Sprint(":", *port), nil); err != nil {
		log.Fatal("http.ListenAndServe ", err)
	}
}

// handleWebSocketRequests listens for sock requests and upgrades the connection
func handleWebSocketRequests(w http.ResponseWriter, r *http.Request) {

	// joinedMessage := StatusMessage{"_", fmt.Sprintf("%s has entered the chat room.", user.username)}
	// b, err := joinedMessage.toJSON()
	// if err != nil {
	// 	log.Panicln("Could not marshal ChatMessage ", err)
	// 	return
	// }
	// if err = user.ws.WriteMessage(websocket.TextMessage, b); err != nil {
	// 	log.Panicln("Error sending message ", err)
	// } else {
	// 	log.Println("message sent ", joinedMessage)
	// }

	// user.ws.ReadMessage()

}
