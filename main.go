package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients map[*websocket.Conn]*websocket.Conn = make(map[*websocket.Conn]*websocket.Conn)

func main() {
	port := 8080
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/", handleConnection)
	fmt.Printf("Starting application on port :%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Error while starting to listen: ", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Up and running!"))
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	/**
	Custom check origin function to bypass gorilla check origin which prevents it from running on localhost
	*/
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket: ", err)
		return
	}

	clients[c] = c
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("Error reading data: ", err)
			delete(clients, c)
			return
		}
		fmt.Println("Message from client: ", message)
		for _, con := range clients {
			if con == c {
				continue
			}
			if err := con.WriteMessage(mt, message); err != nil {
				log.Printf("Writing error: %#v\n", err)
				break
			}
		}
	}
}
