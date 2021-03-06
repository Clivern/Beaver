// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package schema

import (
	"context"
	"fmt"
	"strings"

	"github.com/clivern/beaver/core/driver"

	"github.com/spf13/viper"
)

// Migration type
type Migration struct {
	db *driver.Cassandra
}

// NewMigration creates a migration instance
func NewMigration(db *driver.Cassandra) *Migration {
	result := new(Migration)
	result.db = db

	return result
}

// Init inits the migration tables
func (m *Migration) Init(ctx context.Context) error {
	statements := []string{}

	statements = append(statements, fmt.Sprintf(
		"CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }",
		viper.GetString("app.database.cassandra.databaseName"),
	))

	statements = append(statements, fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s.migration (id VARCHAR, PRIMARY KEY(id))",
		viper.GetString("app.database.cassandra.databaseName"),
	))

	for _, statement := range statements {

		err := m.db.Query(ctx, strings.TrimSpace(statement)).Exec()

		if err != nil {
			return fmt.Errorf("Statement: %s and Error: %s", statement, err.Error())
		}
	}

	return nil
}

// Migrate migrates the database
func (m *Migration) Migrate(ctx context.Context) error {
	statements := []string{}

	ok, err := m.IsMigrated(ctx, "v2.0.0")

	if err != nil {
		return err
	}

	if !ok {
		// Trigger v2.0.0 migrations
		statements = append(statements, fmt.Sprintf(
			"CREATE TABLE IF NOT EXISTS %s.node (id UUID, hostname VARCHAR, address VARCHAR, status VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP, PRIMARY KEY (id, hostname));",
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"CREATE TABLE IF NOT EXISTS %s.client (id UUID, node_id UUID, status VARCHAR, api_key VARCHAR, channels set<text>, created_at TIMESTAMP, updated_at TIMESTAMP, PRIMARY KEY (id, node_id));",
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"CREATE TABLE IF NOT EXISTS %s.channel (id UUID, name VARCHAR, type VARCHAR, created_at TIMESTAMP, updated_at TIMESTAMP, PRIMARY KEY (name));",
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"CREATE TABLE IF NOT EXISTS %s.message (id UUID, message TEXT, from_client_id UUID, to_channel_id UUID, to_channel_name VARCHAR, to_client_id UUID, created_at TIMESTAMP, updated_at TIMESTAMP, PRIMARY KEY (id));",
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"CREATE INDEX %s_client_channels_idx ON %s.client ( channels );",
			viper.GetString("app.database.cassandra.databaseName"),
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"CREATE INDEX %s_message_from_client_id_idx ON %s.message ( from_client_id );",
			viper.GetString("app.database.cassandra.databaseName"),
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"CREATE INDEX %s_message_to_channel_id_idx ON %s.message ( to_channel_id );",
			viper.GetString("app.database.cassandra.databaseName"),
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"CREATE INDEX %s_message_to_channel_name_idx ON %s.message ( to_channel_name );",
			viper.GetString("app.database.cassandra.databaseName"),
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"CREATE INDEX %s_message_to_client_id_idx ON %s.message ( to_client_id );",
			viper.GetString("app.database.cassandra.databaseName"),
			viper.GetString("app.database.cassandra.databaseName"),
		))

		statements = append(statements, fmt.Sprintf(
			"INSERT INTO %s.migration (id) VALUES ('v2.0.0');",
			viper.GetString("app.database.cassandra.databaseName"),
		))

		for _, statement := range statements {

			err := m.db.Query(ctx, strings.TrimSpace(statement)).Exec()

			if err != nil {
				return fmt.Errorf("Statement: %s and Error: %s", statement, err.Error())
			}
		}
	}

	return nil
}

// IsMigrated check if migration run before
func (m *Migration) IsMigrated(ctx context.Context, id string) (bool, error) {
	var count int

	err := m.db.Query(
		ctx,
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s.migration WHERE id = '%s';",
			viper.GetString("app.database.cassandra.databaseName"),
			id,
		),
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
