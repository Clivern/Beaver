// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/core/util"

	"github.com/spf13/viper"
)

// Channel type
type Channel struct {
	db driver.Database
}

// ChannelResult struct
type ChannelResult struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// LoadFromJSON load object from json
func (c *ChannelResult) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON converts object to json
func (c *ChannelResult) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// NewChannel creates a channel instance
func NewChannel(db driver.Database) *Channel {
	result := new(Channel)
	result.db = db

	return result
}

// ChannelsExist checks if channels exist
func (c *Channel) ChannelsExist(channels []string) (bool, error) {
	for _, channel := range channels {
		exists, err := c.ChannelExist(channel)

		if err != nil {
			return false, fmt.Errorf(
				`Error while getting channel %s`,
				channel,
			)
		}

		if !exists {
			return false, nil
		}
	}

	return true, nil
}

// ChannelExist checks if a channel exists
func (c *Channel) ChannelExist(name string) (bool, error) {
	return c.db.Exists(fmt.Sprintf(
		"%s/channel/%s",
		viper.GetString("app.database.etcd.databaseName"),
		name,
	))
}

// DeleteChannelByName deletes a channel with name
func (c *Channel) DeleteChannelByName(name string) (bool, error) {
	deleted, err := c.db.Delete(fmt.Sprintf(
		"%s/channel/%s",
		viper.GetString("app.database.etcd.databaseName"),
		name,
	))

	return deleted > 0, err
}

// GetSubscribers gets a list of subscribers with channel name (all subscribers)
func (c *Channel) GetSubscribers(name string) ([]string, error) {
	result := []string{}

	subscribers, err := c.db.GetKeys(fmt.Sprintf(
		"%s/channel/%s/subscriber",
		viper.GetString("app.database.etcd.databaseName"),
		name,
	))

	for _, subscriber := range subscribers {
		// Get the subscriber uuid only
		items := strings.Split(util.RemoveTrailingSlash(subscriber), "/")
		result = append(result, items[len(items)-1])
	}

	return result, err
}

// CountSubscribers counts channel subscribers (all subscribers)
func (c *Channel) CountSubscribers(name string) (int, error) {
	subscribers, err := c.GetSubscribers(name)

	return len(subscribers), err
}

// GetListeners gets a list of listeners with channel name (online subscribers)
func (c *Channel) GetListeners(name string) ([]string, error) {
	listeners := []string{}

	subscribers, err := c.GetSubscribers(name)

	// filter out offline subscribers
	for _, subscriber := range subscribers {
		items := strings.Split(subscriber, "/")
		listener := items[len(items)-1]

		online, err := c.isSubscriberOnline(listener)

		if err != nil {
			return listeners, err
		}

		if online {
			listeners = append(listeners, listener)
		}
	}

	return listeners, err
}

// CountListeners counts channel listeners (online subscribers)
func (c *Channel) CountListeners(name string) (int, error) {
	listeners, err := c.GetListeners(name)

	return len(listeners), err
}

// isSubscriberOnline checks if subscriber is online
func (c *Channel) isSubscriberOnline(uuid string) (bool, error) {
	subscriber, err := c.db.Get(fmt.Sprintf(
		"%s/client/%s",
		viper.GetString("app.database.etcd.databaseName"),
		uuid,
	))

	if err != nil {
		return false, err
	}

	for k, v := range subscriber {
		// Check if it is the status key
		if strings.Contains(k, "/status") {
			if v == "online" {
				return true, nil
			} else {
				return false, nil
			}
		}
	}

	return false, nil
}

// CreateChannel creates a channel
func (c *Channel) CreateChannel(channel ChannelResult) (bool, error) {
	result, err := channel.ConvertToJSON()

	if err != nil {
		return false, nil
	}

	err = c.db.Put(fmt.Sprintf(
		"%s/channel/%s/info",
		viper.GetString("app.database.etcd.databaseName"),
		channel.Name,
	), result)

	if err != nil {
		return false, nil
	}

	return true, nil
}

// GetChannelByName gets a channel by name
func (c *Channel) GetChannelByName(name string) (ChannelResult, error) {
	var channelResult ChannelResult

	data, err := c.db.Get(fmt.Sprintf(
		"%s/channel/%s/info",
		viper.GetString("app.database.etcd.databaseName"),
		name,
	))

	if err != nil {
		return channelResult, err
	}

	for k, v := range data {
		// Check if it is the info key
		if strings.Contains(k, "/info") {
			_, err = channelResult.LoadFromJSON([]byte(v))

			if err != nil {
				return channelResult, err
			}

			return channelResult, nil
		}
	}

	return channelResult, fmt.Errorf(
		"Unable to find channel %s",
		name,
	)
}

// UpdateChannelByName updates a channel by name
func (c *Channel) UpdateChannelByName(channel ChannelResult) (bool, error) {
	result, err := channel.ConvertToJSON()

	if err != nil {
		return false, nil
	}

	err = c.db.Put(fmt.Sprintf(
		"%s/channel/%s/info",
		viper.GetString("app.database.etcd.databaseName"),
		channel.Name,
	), result)

	if err != nil {
		return false, nil
	}

	return true, nil
}
