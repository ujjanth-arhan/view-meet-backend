package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	port := 8080
	http.HandleFunc("/", handleConnection)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println("Error while starting to listen: ", err)
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	/**
	Custom check origin function to bypass gorilla check origin which prevents it from running on local host
	*/
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to websocket: ", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("Error reading data: ", err)
			break
		}

		fmt.Println("Message from client: ", message)
		if err := c.WriteMessage(mt, message); err != nil {
			log.Printf("Writing error: %#v\n", err)
			break
		}
	}
}
