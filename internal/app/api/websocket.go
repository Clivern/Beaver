// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gorilla/websocket"
	"net/http"
)

// Message struct
type Message struct {
	Client string `json:"client"`
	Data   string `json:"data"`
}

// Websocket Object
type Websocket struct {
	Clients   map[string]*websocket.Conn
	Broadcast chan Message
	Upgrader  websocket.Upgrader
}

// Init initialize the websocket object
func (e *Websocket) Init() {
	e.Clients = make(map[string]*websocket.Conn)
	e.Broadcast = make(chan Message)
	e.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}
}

// HandleMessages send messages to connected clients
func (e *Websocket) HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-e.Broadcast

		// Send it out to every client that is currently connected
		for id, client := range e.Clients {
			msg.Client = id
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("error: %v", err)
				client.Close()
				delete(e.Clients, id)
			}
		}
	}
}

// HandleConnections manage new clients
func (e *Websocket) HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := e.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		logger.Fatal(err)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	clientID := utils.GenerateUUID()

	// Register our new client
	e.Clients[clientID] = ws

	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)

		if err != nil {
			fmt.Printf("error: %v", err)
			delete(e.Clients, clientID)
			break
		}

		// Send the newly received message to the broadcast channel
		e.Broadcast <- msg
	}
}
