// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package schema

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	// database query var
	database = fmt.Sprintf(
		"CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };",
		viper.GetString("app.database.cassandra.databaseName"),
	)

	// client_table query var
	client_table = fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s.client (id UUID PRIMARY KEY, node_id UUID, status VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);",
		viper.GetString("app.database.cassandra.databaseName"),
	)

	// channel_table query var
	channel_table = fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s.channel (id UUID PRIMARY KEY, slug VARCHAR, name VARCHAR, type VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);",
		viper.GetString("app.database.cassandra.databaseName"),
	)

	// message_table query var
	message_table = fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s.message (id UUID PRIMARY KEY, message TEXT, from_client_id UUID, to_channel_id UUID, to_client_id UUID, created_at TIMESTAMP, updated_at TIMESTAMP);",
		viper.GetString("app.database.cassandra.databaseName"),
	)

	// node_table query var
	node_table = fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s.node (id UUID PRIMARY KEY, address VARCHAR, status VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);",
		viper.GetString("app.database.cassandra.databaseName"),
	)

	// client_channel_table query var
	client_channel_table = fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s.client_channel (id UUID PRIMARY KEY, client_id UUID, channel_id UUID);",
		viper.GetString("app.database.cassandra.databaseName"),
	)
)
