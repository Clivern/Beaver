// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/clivern/beaver/core/util"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Message struct
type Message struct {
	FromClient string `json:"from_client"`
	ToClient   string `json:"to_client"`
	Channel    string `json:"channel"`
	Data       string `json:"data"`
}

// BroadcastRequest struct
type BroadcastRequest struct {
	Channels []string `json:"channels"`
	Data     string   `json:"data"`
}

// PublishRequest struct
type PublishRequest struct {
	Channel string `json:"channel"`
	Data    string `json:"data"`
}

// Websocket Object
type Websocket struct {
	Clients          util.Map
	Broadcast        chan Message
	PersistBroadcast chan Message
	Upgrader         websocket.Upgrader
}

// IsValid checks if message is valid
func (m *Message) IsValid() bool {
	validator := util.Validator{}

	return validator.IsJSON(m.Data)
}

// Init initialize the websocket object
func (e *Websocket) Init() {
	e.Clients = util.NewMap()

	e.Broadcast = make(chan Message)
	e.PersistBroadcast = make(chan Message)

	e.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}
}

// HandleConnections manage new clients
func (e *Websocket) HandleConnections(w http.ResponseWriter, r *http.Request, ID string, token string, correlationID string) {
}

// HandleMessages send messages to a specific connected client
func (e *Websocket) HandleMessages() {
}

// BroadcastAction controller
func (e *Websocket) BroadcastAction(c *gin.Context, rawBody []byte) {
	c.Status(http.StatusOK)
}

// PublishAction controller
func (e *Websocket) PublishAction(c *gin.Context, rawBody []byte) {
	c.Status(http.StatusOK)
}

// HandleBroadcastedMessages
func (e *Websocket) HandleBroadcastedMessages() {
	// Handle incoming messages from RabbitMQ
}

// HandlePersistenceCallback send message to backend
func (e *Websocket) HandlePersistenceCallback() {
}
