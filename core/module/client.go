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

// Client type
type Client struct {
	db driver.Cassandra
}

// ClientResult struct
type ClientResult struct {
	ID        string   `json:"id"`
	Token     string   `json:"token"`
	Channels  []string `json:"channels"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

// GenerateClient generates a new client
func GenerateClient(channels []string) *ClientResult {
	now := time.Now().Unix()

	return &ClientResult{
		ID:        util.GenerateUUID4(),
		Token:     util.CreateHash(util.GenerateUUID4()),
		Channels:  channels,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// NewClient creates a client instance
func NewClient(db driver.Cassandra) *Client {
	result := new(Client)
	result.db = db

	return result
}

// CreateClient stores a client
func (c *Client) CreateClient(client ClientResult) (bool, error) {
	return true, nil
}

// UpdateClientByID updates a client by ID
func (c *Client) UpdateClientByID(client ClientResult) (bool, error) {
	return true, nil
}

// GetClientByID gets a client by ID
func (c *Client) GetClientByID(ID string) (ClientResult, error) {
	var clientResult ClientResult

	return clientResult, fmt.Errorf(
		"Unable to find client %s",
		ID,
	)
}

// GetClientNode gets a client node
func (c *Client) GetClientNode(ID string) (string, error) {
	return "", fmt.Errorf(
		"Unable to find the client node %s",
		ID,
	)
}

// DeleteClientByID deletes a client with ID
func (c *Client) DeleteClientByID(ID string) (bool, error) {
	return false, nil
}

// Unsubscribe from channels
func (c *Client) Unsubscribe(ID string, channels []string) (bool, error) {
	return true, nil
}

// Subscribe to channels
func (c *Client) Subscribe(ID string, channels []string) (bool, error) {
	return true, nil
}

// Connect a client
func (c *Client) Connect(clientID, node string) error {
	return nil
}

// Disconnect a client
func (c *Client) Disconnect(clientID string) error {
	return nil
}

// addToChannel adds a client to a channel
func (c *Client) addToChannel(ID string, channel string) (bool, error) {
	return true, nil
}

// removeFromChannel removes a client from a channel
func (c *Client) removeFromChannel(ID string, channel string) (bool, error) {
	return false, nil
}
