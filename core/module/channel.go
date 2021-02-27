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

	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

// ChannelModule type
type ChannelModule struct {
	db driver.Cassandra
}

// ChannelModel struct
type ChannelModel struct {
	ID        gocql.UUID `json:"id"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	CreatedAt int64      `json:"created_at"`
	UpdatedAt int64      `json:"updated_at"`
}

// NewChannel creates a channel instance
func NewChannel(db driver.Cassandra) *ChannelModule {
	result := new(ChannelModule)
	result.db = db

	return result
}

// ChannelsExist checks if channels exist
func (c *ChannelModule) ChannelsExist(ctx context.Context, channels []string) (bool, error) {
	var count int

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.channel WHERE name IN ('%s');",
			viper.GetString("app.database.cassandra.databaseName"),
			strings.Join(channels, "', '"),
		),
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return (count == len(channels)), nil
}

// ChannelExist checks if a channel exists
func (c *ChannelModule) ChannelExist(ctx context.Context, name string) (bool, error) {
	var count int

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.channel WHERE name = %s;",
			viper.GetString("app.database.cassandra.databaseName"),
			name,
		),
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return (count > 0), nil
}

// DeleteChannelByName deletes a channel with name
func (c *ChannelModule) DeleteChannelByName(ctx context.Context, name string) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			`BEGIN BATCH
				DELETE FROM %s.channel WHERE name = '%s';
				DELETE FROM %s.message WHERE to_channel_name = '%s';
				DELETE FROM %s.client_channel WHERE channel_name = '%s';
			APPLY BATCH;`,
			viper.GetString("app.database.cassandra.databaseName"),
			name,
			viper.GetString("app.database.cassandra.databaseName"),
			name,
			viper.GetString("app.database.cassandra.databaseName"),
			name,
		),
	).Exec()
}

// GetSubscribers gets a list of subscribers with channel name (all subscribers)
func (c *ChannelModule) GetSubscribers(name string) []ClientModel {
	var clientIDs []gocql.UUID
	var clients []ClientModel

	iter := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT client_id FROM %s.client_channel WHERE channel_name = '%s';",
			viper.GetString("app.database.cassandra.databaseName"),
			name,
		),
	).Iter()

	for _, columnInfo := range iter.Columns() {
		clientIDs = append(clientIDs, columnInfo.client_id)
		clients = append(clients, ClientModel{
			ID: columnInfo.client_id,
		})
	}

	return clients
}

// CountSubscribers counts channel subscribers (all subscribers)
func (c *ChannelModule) CountSubscribers(ctx context.Context, name string) (int, error) {
	var count int

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.client_channel WHERE channel_name = '%s';",
			viper.GetString("app.database.cassandra.databaseName"),
			name,
		),
	).Scan(&count)

	if err != nil {
		return count, err
	}

	return count, nil
}

// GetListeners gets a list of listeners with channel name (online subscribers)
func (c *ChannelModule) GetListeners(name string) []ClientModel {
	var clientIDs []gocql.UUID
	var clients []ClientModel

	iter := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT client_id FROM %s.client_channel WHERE channel_name = '%s' AND client_status = '%s';",
			viper.GetString("app.database.cassandra.databaseName"),
			name,
			Online,
		),
	).Iter()

	for _, columnInfo := range iter.Columns() {
		clientIDs = append(clientIDs, columnInfo.client_id)
		clients = append(clients, ClientModel{
			ID: columnInfo.client_id,
		})
	}

	return clients
}

// CountListeners counts channel listeners (online subscribers)
func (c *ChannelModule) CountListeners(ctx context.Context, name string) (int, error) {
	var count int

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.client_channel WHERE channel_name = '%s' AND client_status = '%s';",
			viper.GetString("app.database.cassandra.databaseName"),
			name,
			Online,
		),
	).Scan(&count)

	if err != nil {
		return count, err
	}

	return count, nil
}

// isSubscriberOnline checks if subscriber is online
func (c *ChannelModule) isSubscriberOnline(ctx context.Context, id gocql.UUID) (bool, error) {
	var count int

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.client WHERE id = %s AND status = '%s';",
			viper.GetString("app.database.cassandra.databaseName"),
			id,
			Online,
		),
	).Scan(&count)

	if err != nil {
		return count > 0, err
	}

	return count > 0, nil
}

// CreateChannel creates a channel
func (c *ChannelModule) CreateChannel(ctx context.Context, channel ChannelModel) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"INSERT INTO %s.channel (id, name, type, created_at, updated_at) VALUES (%s, '%s', '%s', %d, %d);",
			viper.GetString("app.database.cassandra.databaseName"),
			channel.ID,
			channel.Name,
			channel.Type,
			channel.CreatedAt,
			channel.UpdatedAt,
		),
	).Exec()
}

// GetChannelByName gets a channel by name
func (c *ChannelModule) GetChannelByName(ctx context.Context, name string) (ChannelModel, error) {
	var id gocql.UUID
	var ch_type string
	var createdAt int64
	var updatedAt int64
	var channelModel ChannelModel

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT id, type, created_at, updated_at FROM %s.channel WHERE name = '%s';",
			viper.GetString("app.database.cassandra.databaseName"),
			name,
		),
	).Scan(&id, &ch_type, &createdAt, &updatedAt)

	if err != nil {
		return channelModel, err
	}

	channelModel.ID = id
	channelModel.Name = name
	channelModel.Type = ch_type
	channelModel.CreatedAt = createdAt
	channelModel.UpdatedAt = updatedAt

	return channelModel, nil
}

// UpdateChannelByName updates a channel by name
func (c *ChannelModule) UpdateChannelByName(ctx context.Context, channel_name, channel_type string) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"UPDATE %s.channel SET type = '%s', updated_at = %d WHERE name = '%s' IF EXISTS;",
			viper.GetString("app.database.cassandra.databaseName"),
			channel_type,
			time.Now().Unix(),
			channel_name,
		),
	).Exec()
}

// Count ...
func (c *ChannelModule) Count(ctx context.Context) (int, error) {
	var count int

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.channel;",
			viper.GetString("app.database.cassandra.databaseName"),
		),
	).Scan(&count)

	if err != nil {
		return count, err
	}

	return count, nil
}
