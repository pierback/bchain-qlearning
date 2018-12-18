package websocket

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func wsInit() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8089", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(res http.ResponseWriter, req *http.Request) {
	// Create log file and logger.
	logFile, err := os.Create("pong.log")
	if err != nil {
		fmt.Println("Failed to create pong.log")
		return
	}
	defer logFile.Close()
	log := log.New(logFile, "", log.Lmicroseconds)

	// Perform WebSocket upgrade.
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println("Upgrade Error: ", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connection initiated.")

	for {
		msgType, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Error: ", err)
			break
		}

		// We don't recognize any message that is not "ping".
		if msg := string(bytes[:]); msgType != websocket.TextMessage && msg != "ping" {
			log.Println("Unrecognized message received.")
			continue
		} else {
			log.Println("Received: ping.")
		}
		fmt.Print("Websocket Server started")
		ticker := time.NewTicker(10 * time.Second)
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:
					log.Println("Sending: pong.")
					err = conn.WriteMessage(websocket.TextMessage, []byte("pong"))
					if err != nil {
						log.Println("Write Error: ", err)
						break
					}
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
	}
	log.Println("WebSocket connection terminated.")
}
