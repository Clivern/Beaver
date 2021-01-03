// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build integration

package cluster

import (
	"fmt"
	"testing"
	"time"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"github.com/spf13/viper"
)

// TestIntegrationNodeAlive
func TestIntegrationNodeAlive(t *testing.T) {
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

	node := NewNode(db)
	stats := NewStats(db)

	g := goblin.Goblin(t)

	g.Describe("#NodeAlive", func() {
		g.It("It should return zero count", func() {
			result, err := stats.GetTotalNodes()

			g.Assert(result).Equal(0)
			g.Assert(err).Equal(nil)
		})

		g.It("It should return nil", func() {
			g.Assert(node.Alive(1)).Equal(nil)
		})

		g.It("It should return 1", func() {
			result, err := stats.GetTotalNodes()

			g.Assert(result).Equal(1)
			g.Assert(err).Equal(nil)
		})

		g.It("It should return zero count", func() {
			// Wait till lease expire
			time.Sleep(3 * time.Second)

			result, err := stats.GetTotalNodes()

			g.Assert(result).Equal(0)
			g.Assert(err).Equal(nil)
		})
	})
}
