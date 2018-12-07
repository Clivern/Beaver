// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"encoding/json"
	"fmt"
	"github.com/clivern/beaver/internal/app/api"
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// Message struct
type Message struct {
	Client string `json:"client"`
	Data   string `json:"data"`
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
	Clients   map[string]*websocket.Conn
	Broadcast chan Message
	Upgrader  websocket.Upgrader
}

// LoadFromJSON load object from json
func (c *BroadcastRequest) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON converts object to json
func (c *BroadcastRequest) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// LoadFromJSON load object from json
func (c *PublishRequest) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON converts object to json
func (c *PublishRequest) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(data), nil
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

// PushMessages push messages to all clients
func (e *Websocket) PushMessages(_ http.ResponseWriter, _ *http.Request) {
	msg := Message{
		Data: "HeyBuddy",
	}
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

// HandleConnections manage new clients
func (e *Websocket) HandleConnections(w http.ResponseWriter, r *http.Request, ID string, token string, correlationID string) {

	var clientResult api.ClientResult
	validate := utils.Validator{}

	// Validate client uuid & token
	if validate.IsEmpty(ID) || validate.IsEmpty(token) || !validate.IsUUID4(ID) {
		return
	}

	client := api.Client{
		CorrelationID: correlationID,
	}

	if !client.Init() {
		return
	}

	clientResult, err := client.GetClientByID(ID)

	if err != nil {
		return
	}

	// Ensure that client is alreay registered before
	if clientResult.ID != ID || clientResult.Token != token {
		return
	}

	ok, err := client.Connect(clientResult)

	if !ok || err != nil {
		return
	}

	// Upgrade initial GET request to a websocket
	ws, err := e.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		logger.Fatal(err)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	e.Clients[ID] = ws

	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)

		if err != nil {
			fmt.Printf("error: %v", err)
			delete(e.Clients, ID)
			break
		}

		msg.Client = ID

		// Send the newly received message to the broadcast channel
		e.Broadcast <- msg
	}
}

// HandleMessages send messages to a specific connected client
func (e *Websocket) HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-e.Broadcast

		// The client id
		ID := msg.Client

		// Push message to that client if it still connected
		// or remove from clients if we can't deliver messages to
		// it anymore
		if client, ok := e.Clients[ID]; ok {
			err := client.WriteJSON(msg)
			if err != nil {
				// Remove client from listeners list
				fmt.Printf("error: %v", err)
				client.Close()
				delete(e.Clients, ID)
			}
		}
	}
}

// BroadcastAction controller
func (e *Websocket) BroadcastAction(c *gin.Context, rawBody []byte) {

	var broadcastRequest BroadcastRequest
	var key string
	var msg Message

	validate := utils.Validator{}

	broadcastRequest.LoadFromJSON(rawBody)

	if !validate.IsSlugs(broadcastRequest.Channels, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Provided client channels are invalid.",
		})
		return
	}

	channel := api.Channel{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	if !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	ok, err := channel.ChannelsExist(broadcastRequest.Channels)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	for _, name := range broadcastRequest.Channels {

		// Push message to all subscribed clients
		iter := channel.ChannelScan(name).Iterator()

		for iter.Next() {
			key = iter.Val()
			if key != "" && validate.IsUUID4(key) {
				msg = Message{
					Client: key,
					Data:   broadcastRequest.Data,
				}

				e.Broadcast <- msg
			}
		}
	}

	c.Status(http.StatusOK)
}

// PublishAction controller
func (e *Websocket) PublishAction(c *gin.Context, rawBody []byte) {

	var publishRequest PublishRequest
	var key string
	var msg Message

	validate := utils.Validator{}

	publishRequest.LoadFromJSON(rawBody)

	if !validate.IsSlug(publishRequest.Channel, 3, 60) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Provided client channel is invalid.",
		})
		return
	}

	channel := api.Channel{
		CorrelationID: c.Request.Header.Get("X-Correlation-ID"),
	}

	if !channel.Init() {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  "Internal server error",
		})
		return
	}

	ok, err := channel.ChannelExist(publishRequest.Channel)

	if !ok || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// Push message to all subscribed clients
	iter := channel.ChannelScan(publishRequest.Channel).Iterator()

	for iter.Next() {
		key = iter.Val()
		if key != "" && validate.IsUUID4(key) {
			msg = Message{
				Client: key,
				Data:   publishRequest.Data,
			}

			e.Broadcast <- msg
		}
	}

	c.Status(http.StatusOK)
}
