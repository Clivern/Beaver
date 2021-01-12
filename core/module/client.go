// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/core/util"

	"github.com/spf13/viper"
)

// Client type
type Client struct {
	db driver.Database
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

// LoadFromJSON load object from json
func (c *ClientResult) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON converts object to json
func (c *ClientResult) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// NewClient creates a client instance
func NewClient(db driver.Database) *Client {
	result := new(Client)
	result.db = db

	return result
}

// CreateClient stores a client
func (c *Client) CreateClient(client ClientResult) (bool, error) {
	result, err := client.ConvertToJSON()

	if err != nil {
		return false, err
	}

	// store client info
	err = c.db.Put(fmt.Sprintf(
		"%s/client/%s/info",
		viper.GetString("app.database.etcd.databaseName"),
		client.ID,
	), result)

	if err != nil {
		return false, err
	}

	// store client status
	err = c.db.Put(fmt.Sprintf(
		"%s/client/%s/status",
		viper.GetString("app.database.etcd.databaseName"),
		client.ID,
	), "offline")

	if err != nil {
		return false, err
	}

	// store client node
	err = c.db.Put(fmt.Sprintf(
		"%s/client/%s/node",
		viper.GetString("app.database.etcd.databaseName"),
		client.ID,
	), "#")

	if err != nil {
		return false, err
	}

	for _, channel := range client.Channels {
		ok, err := c.addToChannel(client.ID, channel)
		if !ok || err != nil {
			return false, err
		}
	}

	return true, nil
}

// UpdateClientByID updates a client by ID
func (c *Client) UpdateClientByID(client ClientResult) (bool, error) {
	client.UpdatedAt = time.Now().Unix()

	result, err := client.ConvertToJSON()

	if err != nil {
		return false, err
	}

	// store client info
	err = c.db.Put(fmt.Sprintf(
		"%s/client/%s/info",
		viper.GetString("app.database.etcd.databaseName"),
		client.ID,
	), result)

	if err != nil {
		return false, err
	}

	return true, nil
}

// GetClientByID gets a client by ID
func (c *Client) GetClientByID(ID string) (ClientResult, error) {
	var clientResult ClientResult

	data, err := c.db.Get(fmt.Sprintf(
		"%s/client/%s/info",
		viper.GetString("app.database.etcd.databaseName"),
		ID,
	))

	if err != nil {
		return clientResult, err
	}

	for k, v := range data {
		// Check if it is the info key
		if strings.Contains(k, "/info") {
			_, err = clientResult.LoadFromJSON([]byte(v))

			if err != nil {
				return clientResult, err
			}

			return clientResult, nil
		}
	}

	return clientResult, fmt.Errorf(
		"Unable to find client %s",
		ID,
	)
}

// DeleteClientByID deletes a client with ID
func (c *Client) DeleteClientByID(ID string) (bool, error) {
	client, err := c.GetClientByID(ID)

	if err != nil {
		return false, err
	}

	for _, channel := range client.Channels {
		ok, err := c.removeFromChannel(ID, channel)
		if !ok || err != nil {
			return false, err
		}
	}

	count, err := c.db.Delete(fmt.Sprintf(
		"%s/client/%s",
		viper.GetString("app.database.etcd.databaseName"),
		ID,
	))

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Unsubscribe from channels
func (c *Client) Unsubscribe(ID string, channels []string) (bool, error) {
	validator := util.Validator{}
	clientResult, err := c.GetClientByID(ID)

	if err != nil {
		return false, err
	}

	for i, channel := range clientResult.Channels {
		if validator.IsIn(channel, channels) {
			ok, err := c.removeFromChannel(ID, channel)
			if !ok || err != nil {
				return false, err
			}
			clientResult.Channels = util.Unset(clientResult.Channels, i)
		}
	}

	return c.UpdateClientByID(clientResult)
}

// Subscribe to channels
func (c *Client) Subscribe(ID string, channels []string) (bool, error) {
	validator := util.Validator{}
	clientResult, err := c.GetClientByID(ID)

	if err != nil {
		return false, err
	}

	for _, channel := range channels {
		if !validator.IsIn(channel, clientResult.Channels) {
			ok, err := c.addToChannel(ID, channel)
			if !ok || err != nil {
				return false, err
			}
			clientResult.Channels = append(clientResult.Channels, channel)
		}
	}

	return c.UpdateClientByID(clientResult)
}

// Connect a client
func (c *Client) Connect(clientID, node string) error {
	// update client status
	err := c.db.Put(fmt.Sprintf(
		"%s/client/%s/status",
		viper.GetString("app.database.etcd.databaseName"),
		clientID,
	), "online")

	if err != nil {
		return err
	}

	err = c.db.Put(fmt.Sprintf(
		"%s/client/%s/node",
		viper.GetString("app.database.etcd.databaseName"),
		clientID,
	), node)

	return err
}

// Disconnect a client
func (c *Client) Disconnect(clientID string) error {
	// update client status
	err := c.db.Put(fmt.Sprintf(
		"%s/client/%s/status",
		viper.GetString("app.database.etcd.databaseName"),
		clientID,
	), "offline")

	if err != nil {
		return err
	}

	err = c.db.Put(fmt.Sprintf(
		"%s/client/%s/node",
		viper.GetString("app.database.etcd.databaseName"),
		clientID,
	), "#")

	return err
}

// addToChannel adds a client to a channel
func (c *Client) addToChannel(ID string, channel string) (bool, error) {
	// Add client from channel subscribers
	err := c.db.Put(fmt.Sprintf(
		"%s/channel/%s/subscriber/%s",
		viper.GetString("app.database.etcd.databaseName"),
		channel,
		ID,
	), "#")

	if err != nil {
		return false, err
	}

	return true, nil
}

// removeFromChannel removes a client from a channel
func (c *Client) removeFromChannel(ID string, channel string) (bool, error) {
	// Remove client from channel subscribers
	count, err := c.db.Delete(fmt.Sprintf(
		"%s/channel/%s/subscriber/%s",
		viper.GetString("app.database.etcd.databaseName"),
		channel,
		ID,
	))

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
