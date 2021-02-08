package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// this will handle comms
func webSocketHandler(w http.ResponseWriter, r *http.Request) {

}

//TODO: api handler for "log in" name
func apiHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	port := flag.Int("port", 8080, "port to use, defaults to 8080")
	flag.Parse()

	log.Println("Starting up at-chat on port", *port)

	//serves the chat app website
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/ws", webSocketHandler)

	err := http.ListenAndServe(fmt.Sprint(":", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
