package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Coordinate struct {
	Long float64 `json:"longitude"`
	Lat  float64 `json:"latitude"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *Coordinate)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.HandleFunc("/coordinate", LongLatHandler).Methods("POST")
	router.HandleFunc("/ws", wsHandler)
	go boardcastEvents()

	log.Printf("Starting Coordinate Server at :%d ", 8444)
	log.Fatal(http.ListenAndServe(":8444", router))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func LongLatHandler(w http.ResponseWriter, r *http.Request) {
	var coordinate Coordinate
	if err := json.NewDecoder(r.Body).Decode(&coordinate); err != nil {
		log.Printf("Error %s\n", err)
		http.Error(w, "Bad Request", http.StatusTeapot)
		return
	}

	defer r.Body.Close()
	go writer(&coordinate)
}

func writer(coordinate *Coordinate) {
	broadcast <- coordinate
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	clients[ws] = true
}

func boardcastEvents() {
	val := <-broadcast
	coordinate := fmt.Sprintf("%f %f ", val.Long, val.Lat)
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(coordinate))
		if err != nil {
			log.Printf("Error writing to websocket %s", err)
			client.Close()
			delete(clients, client)
		}
	}
}
