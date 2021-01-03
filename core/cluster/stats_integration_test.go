// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build integration

package cluster

import (
	"fmt"
	"testing"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"github.com/spf13/viper"
)

// TestIntegrationNodeStats
func TestIntegrationNodeStats(t *testing.T) {
	// Skip if -short flag exist
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	db := driver.NewEtcdDriver()
	db.Connect()
	defer db.Close()

	// Cleanup
	db.Delete(viper.GetString("app.database.etcd.databaseName"))

	stats := NewStats(db)

	g := goblin.Goblin(t)

	g.Describe("#GetTotalNodes", func() {
		g.It("It should return zero count", func() {
			result, err := stats.GetTotalNodes()

			g.Assert(result).Equal(0)
			g.Assert(err).Equal(nil)
		})

		g.It("It should return 2", func() {
			db.Put(fmt.Sprintf(
				"%s/node/node1/item1",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/node/node1/item2",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/node/node2/item1",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/node/node2/item2",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			result, err := stats.GetTotalNodes()

			g.Assert(result).Equal(2)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#GetTotalChannels", func() {
		g.It("It should return zero count", func() {
			result, err := stats.GetTotalChannels()

			g.Assert(result).Equal(0)
			g.Assert(err).Equal(nil)
		})

		g.It("It should return 2", func() {
			db.Put(fmt.Sprintf(
				"%s/channel/channel1/item1",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/channel/channel1/item2",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/channel/channel2/item1",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/channel/channel2/item2",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			result, err := stats.GetTotalChannels()

			g.Assert(result).Equal(2)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#GetTotalClients", func() {
		g.It("It should return zero count", func() {
			result, err := stats.GetTotalClients()

			g.Assert(result).Equal(0)
			g.Assert(err).Equal(nil)
		})

		g.It("It should return 2", func() {
			db.Put(fmt.Sprintf(
				"%s/client/client1/item1",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/client/client1/item2",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/client/client2/item1",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			db.Put(fmt.Sprintf(
				"%s/client/client2/item2",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			result, err := stats.GetTotalClients()

			g.Assert(result).Equal(2)
			g.Assert(err).Equal(nil)
		})
	})
}
