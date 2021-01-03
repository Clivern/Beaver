// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build integration

package module

import (
	"fmt"
	"testing"
	"time"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"github.com/spf13/viper"
)

// TestIntegrationChannel
func TestIntegrationChannel(t *testing.T) {
	// Skip if -short flag exist
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	db := driver.NewEtcdDriver()
	db.Connect()
	defer db.Close()

	channel := NewChannel(db)

	// Cleanup
	db.Delete(viper.GetString("app.database.etcd.databaseName"))

	g := goblin.Goblin(t)

	g.Describe("#CreateChannel", func() {
		g.It("It should return true & create channel", func() {
			result, err := channel.CreateChannel(ChannelResult{
				Name:      "found",
				Type:      "public",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			})

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#ChannelExist", func() {
		g.It("It should return false", func() {
			result, err := channel.ChannelExist("not_found")
			g.Assert(result).Equal(false)
			g.Assert(err).Equal(nil)
		})

		g.It("It should return true", func() {
			result, err := channel.ChannelExist("found")
			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#ChannelsExist", func() {
		g.It("It should return false", func() {
			result, err := channel.ChannelsExist([]string{"not_found"})

			g.Assert(result).Equal(false)
			g.Assert(err).Equal(nil)
		})

		g.It("It should return true", func() {
			result, err := channel.ChannelsExist([]string{"found"})

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#GetChannelByName", func() {
		g.It("It should return channel by name", func() {
			result, err := channel.GetChannelByName("found")

			g.Assert(result.Name).Equal("found")
			g.Assert(err).Equal(nil)
		})

		g.It("It should return empty channel & error", func() {
			result, err := channel.GetChannelByName("not_found")

			g.Assert(result.Name).Equal("")
			g.Assert(err != nil).Equal(true)
		})
	})

	g.Describe("#UpdateChannelByName", func() {
		g.It("It should update the channel type by name", func() {
			result1, err := channel.UpdateChannelByName(ChannelResult{
				Name:      "found",
				Type:      "private",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			})

			g.Assert(result1).Equal(true)
			g.Assert(err).Equal(nil)

			result2, err := channel.GetChannelByName("found")

			g.Assert(result2.Type).Equal("private")
			g.Assert(err).Equal(nil)
		})
	})
}
