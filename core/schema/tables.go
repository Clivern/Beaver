// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package schema

import (
	"strings"
)

var (
	// Database query var
	Database = "CREATE KEYSPACE IF NOT EXISTS [Database] WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };"

	// ClientTable query var
	ClientTable = "CREATE TABLE IF NOT EXISTS [Database].client (id UUID PRIMARY KEY, node_id UUID, status VARCHAR, api_key VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);"

	// ChannelTable query var
	ChannelTable = "CREATE TABLE IF NOT EXISTS [Database].channel (id UUID PRIMARY KEY, name VARCHAR, type VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);"

	// MessageTable query var
	MessageTable = "CREATE TABLE IF NOT EXISTS [Database].message (id UUID PRIMARY KEY, message TEXT, from_client_id UUID, to_channel_id UUID, to_channel_name VARCHAR, to_client_id UUID, created_at TIMESTAMP, updated_at TIMESTAMP);"

	// NodeTable query var
	NodeTable = "CREATE TABLE IF NOT EXISTS [Database].node (id UUID PRIMARY KEY, hostname VARCHAR, address VARCHAR, status VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP);"

	// ClientChannelTable query var
	ClientChannelTable = "CREATE TABLE IF NOT EXISTS [Database].client_channel (id UUID PRIMARY KEY, client_id UUID, channel_id UUID, channel_name VARCHAR, client_status VARCHAR);"
)

// SchemaWithDatabase gets the query with database
func SchemaWithDatabase(database, query string) string {
	return strings.Replace(query, "[Database]", database, -1)
}
