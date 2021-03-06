// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"context"
	"fmt"

	"github.com/clivern/beaver/core/driver"

	"github.com/gocql/gocql"
	"github.com/spf13/viper"
)

// NodeModule type
type NodeModule struct {
	db *driver.Cassandra
}

// NodeModel struct
type NodeModel struct {
	ID        gocql.UUID `json:"id"`
	Address   string     `json:"address"`
	Status    string     `json:"status"`
	Hostname  string     `json:"hostname"`
	CreatedAt int64      `json:"created_at"`
	UpdatedAt int64      `json:"updated_at"`
}

// NewNode creates a client instance
func NewNodeModule(db *driver.Cassandra) *NodeModule {
	result := new(NodeModule)
	result.db = db

	return result
}

// Insert ...
func (c *NodeModule) Insert(ctx context.Context, node NodeModel) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"INSERT INTO %s.node (id, hostname, address, status, created_at, updated_at) VALUES (%s, '%s', '%s', '%s', %d, %d);",
			viper.GetString("app.database.cassandra.databaseName"),
			node.Hostname,
			node.ID,
			node.Address,
			node.Status,
			node.CreatedAt,
			node.UpdatedAt,
		),
	).Exec()
}

// DeleteById ...
func (c *NodeModule) DeleteById(ctx context.Context, id gocql.UUID) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"DELETE FROM %s.node WHERE id = %s",
			viper.GetString("app.database.cassandra.databaseName"),
			id,
		),
	).Exec()
}

// UpdateById ...
func (c *NodeModule) UpdateById(ctx context.Context, node NodeModel) error {
	return c.db.Query(
		ctx,
		fmt.Sprintf(
			"UPDATE %s.node SET hostname = '%s', address = '%s', status = '%s', created_at = %d, updated_at = %d WHERE id = %s IF EXISTS;",
			viper.GetString("app.database.cassandra.databaseName"),
			node.Hostname,
			node.Address,
			node.Status,
			node.CreatedAt,
			node.UpdatedAt,
			node.ID,
		),
	).Exec()
}

// GetById ...
func (c *NodeModule) GetById(ctx context.Context, id gocql.UUID) (NodeModel, error) {
	var address string
	var status string
	var hostname string
	var createdAt int64
	var updatedAt int64
	var nodeModel NodeModel

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT hostname, address, status, created_at, updated_at FROM %s.node WHERE id = %s;",
			viper.GetString("app.database.cassandra.databaseName"),
			id,
		),
	).Scan(&hostname, &address, &status, &createdAt, &updatedAt)

	if err != nil {
		return nodeModel, err
	}

	nodeModel.ID = id
	nodeModel.Hostname = hostname
	nodeModel.Address = address
	nodeModel.Status = status
	nodeModel.CreatedAt = createdAt
	nodeModel.UpdatedAt = updatedAt

	return nodeModel, nil
}

// Exists ...
func (c *NodeModule) Exists(ctx context.Context, id gocql.UUID) (bool, error) {
	var count int

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.node WHERE id = %s;",
			viper.GetString("app.database.cassandra.databaseName"),
			id,
		),
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return (count > 0), nil
}

// Count ...
func (c *NodeModule) Count(ctx context.Context) (int, error) {
	var count int

	err := c.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.node;",
			viper.GetString("app.database.cassandra.databaseName"),
		),
	).Scan(&count)

	if err != nil {
		return count, err
	}

	return count, nil
}
