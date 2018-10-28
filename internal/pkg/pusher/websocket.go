// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pusher

import (
	"fmt"
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gorilla/websocket"
	"net/http"
)

// Message Object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Channel  string `json:"channel"`
}

// Websocket Object
type Websocket struct {
	Clients   map[string]map[*websocket.Conn]bool
	Broadcast chan Message
	Upgrader  websocket.Upgrader
}

// Websocket Init
func (e *Websocket) Init() {
	e.Clients = make(map[string]map[*websocket.Conn]bool)
	e.Broadcast = make(chan Message)
	e.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}
}

// Websocket HandleMessages
func (e *Websocket) HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-e.Broadcast

		// Send it out to every client that is currently connected
		for client := range e.Clients[msg.Channel] {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("error: %v", err)
				client.Close()
				delete(e.Clients[msg.Channel], client)
			}
		}
	}
}

// Websocket HandleConnections
func (e *Websocket) HandleConnections(w http.ResponseWriter, r *http.Request, appName string) {
	// Upgrade initial GET request to a websocket
	ws, err := e.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		logger.Fatal(err)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	if _, ok := e.Clients[appName]; !ok {
		e.Clients[appName] = make(map[*websocket.Conn]bool)
	}

	fmt.Println(utils.GenerateUUID())

	// Register our new client
	e.Clients[appName][ws] = true

	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)

		if err != nil {
			fmt.Printf("error: %v", err)
			delete(e.Clients[appName], ws)
			break
		}

		// Send the newly received message to the broadcast channel
		e.Broadcast <- msg
	}
}
