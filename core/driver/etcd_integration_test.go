// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build integration

package driver

import (
	"fmt"
	"testing"
	"time"

	"github.com/clivern/beaver/core/util"
	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"github.com/spf13/viper"
)

// TestIntegrationEtcd
func TestIntegrationEtcd(t *testing.T) {
	// Skip if -short flag exist
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	db := NewEtcdDriver()

	g.Describe("#IsConnected", func() {
		g.It("It should return false", func() {
			g.Assert(db.IsConnected()).Equal(false)
		})
	})

	g.Describe("#Connect", func() {
		g.It("It should connect to etcd server", func() {
			g.Assert(db.Connect()).Equal(nil)
		})
	})

	g.Describe("#IsConnected", func() {
		g.It("It should return true", func() {
			g.Assert(db.IsConnected()).Equal(true)
		})
	})

	defer db.Close()

	// Cleanup
	db.Delete(viper.GetString("app.database.etcd.databaseName"))

	g.Describe("#Exists", func() {
		g.It("The key should not exist", func() {
			result, err := db.Exists(fmt.Sprintf(
				"%s/exists_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			))
			g.Assert(result).Equal(false)
			g.Assert(err).Equal(nil)
		})

		g.It("The key should exist", func() {
			err := db.Put(fmt.Sprintf(
				"%s/exists_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			result, err := db.Exists(fmt.Sprintf(
				"%s/exists_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			))

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#Get/Put/Delete", func() {
		g.It("The key get_key_01 should not exist", func() {
			result, err := db.Get(fmt.Sprintf(
				"%s/get_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			))
			g.Assert(len(result)).Equal(0)
			g.Assert(err).Equal(nil)
		})

		g.It("The key get_key_01 not exist so delete will return zero", func() {
			result, err := db.Delete(fmt.Sprintf(
				"%s/get_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			))
			g.Assert(result).Equal(int64(0))
			g.Assert(err).Equal(nil)
		})

		g.It("Create the key get_key_01 and check if it exists", func() {
			err := db.Put(fmt.Sprintf(
				"%s/get_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			result, err := db.Exists(fmt.Sprintf(
				"%s/get_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			))

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)
		})

		g.It("Delete the key get_key_01 and get the deleted count", func() {
			result, err := db.Delete(fmt.Sprintf(
				"%s/get_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			))

			g.Assert(result).Equal(int64(1))
			g.Assert(err).Equal(nil)
		})

		g.It("Get the key get_key_01", func() {
			err := db.Put(fmt.Sprintf(
				"%s/get_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			result, err := db.Get(fmt.Sprintf(
				"%s/get_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			))
			g.Assert(len(result)).Equal(1)
			g.Assert(result[fmt.Sprintf(
				"%s/get_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			)]).Equal("#")
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#Leases", func() {
		g.It("Validates lease renew", func() {
			lease1, err := db.CreateLease(5)

			g.Assert(err).Equal(nil)

			err = db.PutWithLease(fmt.Sprintf(
				"%s/test_lease_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			), "#", lease1)

			// Renew lease1
			err = db.RenewLease(lease1)

			g.Assert(err).Equal(nil)

			result, err := db.Exists(fmt.Sprintf(
				"%s/test_lease_key_01",
				viper.GetString("app.database.etcd.databaseName"),
			))

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)
		})

		g.It("Validates lease expire", func() {
			lease, err := db.CreateLease(1)

			g.Assert(err).Equal(nil)

			err = db.PutWithLease(fmt.Sprintf(
				"%s/test_lease_key_02",
				viper.GetString("app.database.etcd.databaseName"),
			), "#", lease)

			g.Assert(err).Equal(nil)

			time.Sleep(4 * time.Second)

			result, err := db.Exists(fmt.Sprintf(
				"%s/test_lease_key_02",
				viper.GetString("app.database.etcd.databaseName"),
			))

			g.Assert(result).Equal(false)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#Keys", func() {
		g.It("Validate result keys", func() {
			err := db.Put(fmt.Sprintf(
				"%s/parent1/child1/info",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			err = db.Put(fmt.Sprintf(
				"%s/parent1/child1/id",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			err = db.Put(fmt.Sprintf(
				"%s/parent1/child2/info",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			err = db.Put(fmt.Sprintf(
				"%s/parent1/child2/id",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			err = db.Put(fmt.Sprintf(
				"%s/parent1/child3/info",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			err = db.Put(fmt.Sprintf(
				"%s/parent1/child3/id",
				viper.GetString("app.database.etcd.databaseName"),
			), "#")

			g.Assert(err).Equal(nil)

			result1, err := db.Get(fmt.Sprintf(
				"%s/parent1/child1",
				viper.GetString("app.database.etcd.databaseName"),
			))

			g.Assert(len(result1)).Equal(2)
			g.Assert(result1[fmt.Sprintf(
				"%s/parent1/child1/info",
				viper.GetString("app.database.etcd.databaseName"),
			)]).Equal("#")
			g.Assert(result1[fmt.Sprintf(
				"%s/parent1/child1/id",
				viper.GetString("app.database.etcd.databaseName"),
			)]).Equal("#")
			g.Assert(err).Equal(nil)

			result2, err := db.GetKeys(fmt.Sprintf(
				"%s/parent1/child2",
				viper.GetString("app.database.etcd.databaseName"),
			))

			g.Assert(util.InArray(fmt.Sprintf(
				"%s/parent1/child2/info",
				viper.GetString("app.database.etcd.databaseName"),
			), result2)).Equal(true)

			g.Assert(util.InArray(fmt.Sprintf(
				"%s/parent1/child2/id",
				viper.GetString("app.database.etcd.databaseName"),
			), result2)).Equal(true)

			g.Assert(len(result2)).Equal(2)
			g.Assert(err).Equal(nil)
		})
	})
}
