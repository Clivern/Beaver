// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/core/util"

	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

const (
	Online  = "online"
	Offline = "offline"
	Unknown = "unknown"
)

// ClientModule type
type ClientModule struct {
	db driver.Cassandra
}

// ClientModel struct
type ClientModel struct {
	ID        gocql.UUID `json:"id"`
	APIKey    string     `json:"api_key"`
	NodeID    gocql.UUID `json:"node_id"`
	Status    string     `json:"status"`
	Channels  []string   `json:"channels"`
	CreatedAt int64      `json:"created_at"`
	UpdatedAt int64      `json:"updated_at"`
}

// GenerateClient generates a new client
func GenerateClient(nodeID gocql.UUID, status string, channels []string) *ClientModel {
	now := time.Now().Unix()

	return &ClientModel{
		ID:        gocql.TimeUUID(),
		APIKey:    util.CreateHash(util.GenerateUUID4()),
		NodeID:    nodeID,
		Status:    status,
		Channels:  channels,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// NewClientModule creates a client instance
func NewClientModule(db driver.Cassandra) *ClientModule {
	result := new(ClientModule)
	result.db = db

	return result
}

// CreateClient stores a client
func (c *ClientModule) CreateClient(ctx context.Context, client ClientModel) error {
	channels := "{}"

	if len(client.Channels) > 0 {
		channels = fmt.Sprintf("{'%s'}", strings.Join(client.Channels, "', '"))
	}

	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"INSERT INTO %s.client (id, node_id, status, api_key, channels, created_at, updated_at) VALUES (%s, %s, '%s', '%s', %s, %d, %d);",
			viper.GetString("app.database.cassandra.databaseName"),
			client.ID,
			client.NodeID,
			client.Status,
			client.APIKey,
			channels,
			client.CreatedAt,
			client.UpdatedAt,
		),
	).Exec()
}

// UpdateClientByID updates a client by ID
func (c *ClientModule) UpdateClientByID(ctx context.Context, client ClientModel) error {
	channels := "{}"

	if len(client.Channels) > 0 {
		channels = fmt.Sprintf("{'%s'}", strings.Join(client.Channels, "', '"))
	}

	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"UPDATE %s.client SET node_id = %s, status = '%s', api_key = '%s', channels = %s, updated_at = %d WHERE id = %s IF EXISTS;",
			viper.GetString("app.database.cassandra.databaseName"),
			client.NodeID,
			client.Status,
			client.APIKey,
			channels,
			time.Now().Unix(),
			client.ID,
		),
	).Exec()
}

// GetClientByID gets a client by ID
func (c *ClientModule) GetClientByID(ctx context.Context, ID gocql.UUID) (ClientModel, error) {
	var apiKey string
	var nodeID gocql.UUID
	var status string
	var channels []string
	var createdAt int64
	var updatedAt int64
	var clientModel ClientModel

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT node_id, status, api_key, channels, created_at, updated_at FROM %s.client WHERE id = %s;",
			viper.GetString("app.database.cassandra.databaseName"),
			ID,
		),
	).Scan(&nodeID, &status, &apiKey, channels, &createdAt, &updatedAt)

	if err != nil {
		return clientModel, err
	}

	clientModel.ID = ID
	clientModel.APIKey = apiKey
	clientModel.NodeID = nodeID
	clientModel.Status = status
	clientModel.Channels = channels
	clientModel.CreatedAt = createdAt
	clientModel.UpdatedAt = updatedAt

	return clientModel, nil
}

// DeleteClientByID deletes a client with ID
func (c *ClientModule) DeleteClientByID(ctx context.Context, ID gocql.UUID) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"DELETE FROM %s.node WHERE id = %s",
			viper.GetString("app.database.cassandra.databaseName"),
			ID,
		),
	).Exec()
}

// Connect a client
func (c *ClientModule) Connect(ctx context.Context, clientID, nodeID gocql.UUID) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"UPDATE %s.client SET status = '%s', node_id = %s, updated_at = %d WHERE id = %s IF EXISTS;",
			viper.GetString("app.database.cassandra.databaseName"),
			Online,
			nodeID,
			time.Now().Unix(),
			clientID,
		),
	).Exec()
}

// Disconnect a client
func (c *ClientModule) Disconnect(ctx context.Context, clientID gocql.UUID) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"UPDATE %s.client SET status = '%s', node_id = NULL, updated_at = %d WHERE id = %s IF EXISTS;",
			viper.GetString("app.database.cassandra.databaseName"),
			Offline,
			time.Now().Unix(),
			clientID,
		),
	).Exec()
}

// Unsubscribe from channels
func (c *ClientModule) Unsubscribe(ctx context.Context, clientID gocql.UUID, channels []string) error {
	for _, channel := range channels {
		err := c.removeFromChannel(ctx, clientID, channel)

		if err != nil {
			return err
		}
	}

	return nil
}

// Subscribe to channels
func (c *ClientModule) Subscribe(ctx context.Context, clientID gocql.UUID, channels []string) error {
	for _, channel := range channels {
		err := c.addToChannel(ctx, clientID, channel)

		if err != nil {
			return err
		}
	}

	return nil
}

// addToChannel adds a client to a channel
func (c *ClientModule) addToChannel(ctx context.Context, clientID gocql.UUID, channel string) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"UPDATE %s.client SET channels = channels + {'%s'}, updated_at = %d WHERE id = %s IF EXISTS;",
			viper.GetString("app.database.cassandra.databaseName"),
			channel,
			time.Now().Unix(),
			clientID,
		),
	).Exec()
}

// removeFromChannel removes a client from a channel
func (c *ClientModule) removeFromChannel(ctx context.Context, clientID gocql.UUID, channel string) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"UPDATE %s.client SET channels = channels - {'%s'}, updated_at = %d WHERE id = %s IF EXISTS;",
			viper.GetString("app.database.cassandra.databaseName"),
			channel,
			time.Now().Unix(),
			clientID,
		),
	).Exec()
}
