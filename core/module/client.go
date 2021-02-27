// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"
	"time"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/core/util"
)

// ClientModule type
type ClientModule struct {
	db driver.Cassandra
}

// ClientModel struct
type ClientModel struct {
	ID        string   `json:"id"`
	APIKey    string   `json:"api_key"`
	NodeID    string   `json:"node_id"`
	Status    string   `json:"status"`
	Channels  []string `json:"channels"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

// GenerateClient generates a new client
func GenerateClient(channels []string) *ClientModel {
	now := time.Now().Unix()

	return &ClientModel{
		ID:        util.GenerateUUID4(),
		APIKey:    util.CreateHash(util.GenerateUUID4()),
		Channels:  channels,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// NewClient creates a client instance
func NewClient(db driver.Cassandra) *ClientModule {
	result := new(ClientModule)
	result.db = db

	return result
}

// CreateClient stores a client
func (c *ClientModule) CreateClient(client ClientModel) (bool, error) {
	return true, nil
}

// UpdateClientByID updates a client by ID
func (c *ClientModule) UpdateClientByID(client ClientModel) (bool, error) {
	return true, nil
}

// GetClientByID gets a client by ID
func (c *ClientModule) GetClientByID(ID string) (ClientModel, error) {
	var clientResult ClientModel

	return clientResult, fmt.Errorf(
		"Unable to find client %s",
		ID,
	)
}

// GetClientNode gets a client node
func (c *ClientModule) GetClientNode(ID string) (string, error) {
	return "", fmt.Errorf(
		"Unable to find the client node %s",
		ID,
	)
}

// DeleteClientByID deletes a client with ID
func (c *ClientModule) DeleteClientByID(ID string) (bool, error) {
	return false, nil
}

// Unsubscribe from channels
func (c *ClientModule) Unsubscribe(ID string, channels []string) (bool, error) {
	return true, nil
}

// Subscribe to channels
func (c *ClientModule) Subscribe(ID string, channels []string) (bool, error) {
	return true, nil
}

// Connect a client
func (c *ClientModule) Connect(clientID, node string) error {
	return nil
}

// Disconnect a client
func (c *ClientModule) Disconnect(clientID string) error {
	return nil
}

// addToChannel adds a client to a channel
func (c *ClientModule) addToChannel(ID string, channel string) (bool, error) {
	return true, nil
}

// removeFromChannel removes a client from a channel
func (c *ClientModule) removeFromChannel(ID string, channel string) (bool, error) {
	return false, nil
}
