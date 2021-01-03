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

// TestIntegrationCreateClient
func TestIntegrationCreateClient(t *testing.T) {
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
	client := NewClient(db)

	// Cleanup
	db.Delete(viper.GetString("app.database.etcd.databaseName"))

	g := goblin.Goblin(t)

	g.Describe("#CreateClient", func() {
		g.It("It should create client and add channel messages", func() {
			result, err := channel.CreateChannel(ChannelResult{
				Name:      "messages",
				Type:      "public",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			})

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			result, err = channel.CreateChannel(ChannelResult{
				Name:      "messages_01",
				Type:      "private",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			})

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			result, err = channel.CreateChannel(ChannelResult{
				Name:      "messages_02",
				Type:      "presence",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			})

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			newClient := GenerateClient([]string{"messages"})

			result, err = client.CreateClient(*newClient)

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			clientData, err := client.GetClientByID(newClient.ID)

			g.Assert(err).Equal(nil)
			g.Assert(clientData.Channels[0]).Equal("messages")
			g.Assert(len(clientData.Channels)).Equal(1)
			g.Assert(clientData.Token).Equal(newClient.Token)
		})
	})

	g.Describe("#UpdateClientByID", func() {
		g.It("It should update client by id", func() {
			newClient := GenerateClient([]string{"messages"})

			result, err := client.CreateClient(*newClient)

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			newClient.Token = "newToken"

			result, err = client.UpdateClientByID(*newClient)

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			clientData, err := client.GetClientByID(newClient.ID)

			g.Assert(err).Equal(nil)
			g.Assert(clientData.Token).Equal("newToken")
		})
	})

	g.Describe("#DeleteClientByID", func() {
		g.It("It should delete client by id", func() {
			newClient := GenerateClient([]string{
				"messages",
				"messages_01",
				"messages_02",
			})

			result, err := client.CreateClient(*newClient)
			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			result, err = client.DeleteClientByID(newClient.ID)
			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			clientData, err := client.GetClientByID(newClient.ID)
			g.Assert(err != nil).Equal(true)
			g.Assert(clientData.ID).Equal("")
		})
	})

	g.Describe("#Unsubscribe", func() {
		g.It("It should unsubscribe from channel", func() {
			newClient := GenerateClient([]string{
				"messages",
				"messages_01",
				"messages_02",
			})

			result, err := client.CreateClient(*newClient)
			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			result, err = client.Unsubscribe(newClient.ID, []string{
				"messages_01",
				"messages_02",
			})

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			clientData, err := client.GetClientByID(newClient.ID)

			g.Assert(err).Equal(nil)
			g.Assert(len(clientData.Channels)).Equal(2)
		})
	})

	g.Describe("#Subscribe", func() {
		g.It("It should subscribe into channels", func() {
			newClient := GenerateClient([]string{
				"messages",
			})

			result, err := client.CreateClient(*newClient)
			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			result, err = client.Subscribe(newClient.ID, []string{
				"messages",
				"messages_01",
				"messages_02",
			})

			g.Assert(result).Equal(true)
			g.Assert(err).Equal(nil)

			clientData, err := client.GetClientByID(newClient.ID)

			g.Assert(err).Equal(nil)
			g.Assert(len(clientData.Channels)).Equal(3)
		})
	})

	g.Describe("#Connect", func() {
		g.It("It should store client as online", func() {
			newClient := GenerateClient([]string{
				"messages",
			})

			g.Assert(client.Connect(newClient.ID)).Equal(nil)
		})
	})

	g.Describe("#Disconnect", func() {
		g.It("It should store client as offline", func() {
			newClient := GenerateClient([]string{
				"messages",
			})

			g.Assert(client.Disconnect(newClient.ID)).Equal(nil)
		})
	})
}
