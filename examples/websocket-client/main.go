package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8444", "http service address")

func main() {
	flag.Parse()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to websocket %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error: ", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("read error: ", err)
				return
			}
			log.Printf("recv: %s\n", message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Received Interrupt stopping")
			return
		}
	}
}
