// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"encoding/json"
	"github.com/clivern/beaver/internal/app/api"
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
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
	Clients   utils.Map
	Broadcast chan Message
	Upgrader  websocket.Upgrader
}

// IsValid checks if message is valid
func (m *Message) IsValid() bool {
	validator := utils.Validator{}
	return validator.IsJSON(m.Data)
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
	e.Clients = utils.NewMap()
	e.Broadcast = make(chan Message)
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
		logger.Fatalf(
			`Error while upgrading the GET request to a websocket for client %s: %s {"correlationId":"%s"}`,
			ID,
			err.Error(),
			correlationID,
		)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	e.Clients.Set(ID, ws)

	logger.Infof(
		`Client %s connected {"correlationId":"%s"}`,
		ID,
		correlationID,
	)

	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)

		if err != nil {
			e.Clients.Delete(ID)
			client.Disconnect(clientResult)
			logger.Infof(
				`Client %s disconnected {"correlationId":"%s"}`,
				ID,
				correlationID,
			)
			break
		}

		msg.FromClient = ID

		if msg.IsValid() {
			// Send the newly received message to the broadcast channel
			e.Broadcast <- msg
		}
	}
}

// HandleMessages send messages to a specific connected client
func (e *Websocket) HandleMessages() {

	validate := utils.Validator{}

	for {
		// Grab the next message from the broadcast channel
		msg := <-e.Broadcast

		// Send to Client
		if msg.IsValid() && !validate.IsEmpty(msg.ToClient) && !validate.IsEmpty(msg.Channel) && validate.IsUUID4(msg.ToClient) {
			// Push message to that client if it still connected
			// or remove from clients if we can't deliver messages to
			// it anymore
			if client, ok := e.Clients.Get(msg.ToClient); ok {
				err := client.(*websocket.Conn).WriteJSON(msg)
				if err != nil {
					client.(*websocket.Conn).Close()
					e.Clients.Delete(msg.ToClient)
				}
			}
		}

		// Send to client Peers on a channel
		if msg.IsValid() && !validate.IsEmpty(msg.FromClient) && !validate.IsEmpty(msg.Channel) && validate.IsUUID4(msg.FromClient) {

			channel := api.Channel{}
			channel.Init()
			iter := channel.ChannelScan(msg.Channel).Iterator()

			for iter.Next() {

				if msg.FromClient == iter.Val() {
					continue
				}

				msg.ToClient = iter.Val()

				if msg.ToClient != "" && validate.IsUUID4(msg.ToClient) {
					if client, ok := e.Clients.Get(msg.ToClient); ok {
						err := client.(*websocket.Conn).WriteJSON(msg)
						if err != nil {
							client.(*websocket.Conn).Close()
							e.Clients.Delete(msg.ToClient)
						}
					}
				}
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

	if !validate.IsJSON(broadcastRequest.Data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Message data is invalid JSON",
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
					ToClient: key,
					Data:     broadcastRequest.Data,
					Channel:  name,
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

	if !validate.IsJSON(publishRequest.Data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Message data is invalid JSON",
		})
		return
	}

	// Push message to all subscribed clients
	iter := channel.ChannelScan(publishRequest.Channel).Iterator()

	for iter.Next() {
		key = iter.Val()
		if key != "" && validate.IsUUID4(key) {
			msg = Message{
				ToClient: key,
				Data:     publishRequest.Data,
				Channel:  publishRequest.Channel,
			}

			e.Broadcast <- msg
		}
	}

	c.Status(http.StatusOK)
}
