// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"

	"github.com/clivern/beaver/core/driver"
)

// Channel type
type Channel struct {
	db driver.Cassandra
}

// ChannelResult struct
type ChannelResult struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// NewChannel creates a channel instance
func NewChannel(db driver.Cassandra) *Channel {
	result := new(Channel)
	result.db = db

	return result
}

// ChannelsExist checks if channels exist
func (c *Channel) ChannelsExist(channels []string) (bool, error) {
	return true, nil
}

// ChannelExist checks if a channel exists
func (c *Channel) ChannelExist(name string) (bool, error) {
	return false, nil
}

// DeleteChannelByName deletes a channel with name
func (c *Channel) DeleteChannelByName(name string) (bool, error) {
	return true, nil
}

// GetSubscribers gets a list of subscribers with channel name (all subscribers)
func (c *Channel) GetSubscribers(name string) ([]string, error) {
	return []string{}, nil
}

// CountSubscribers counts channel subscribers (all subscribers)
func (c *Channel) CountSubscribers(name string) (int, error) {
	return 0, nil
}

// GetListeners gets a list of listeners with channel name (online subscribers)
func (c *Channel) GetListeners(name string) ([]string, error) {
	return []string{}, nil
}

// CountListeners counts channel listeners (online subscribers)
func (c *Channel) CountListeners(name string) (int, error) {
	return 0, nil
}

// isSubscriberOnline checks if subscriber is online
func (c *Channel) isSubscriberOnline(uuid string) (bool, error) {
	return false, nil
}

// CreateChannel creates a channel
func (c *Channel) CreateChannel(channel ChannelResult) (bool, error) {
	return true, nil
}

// GetChannelByName gets a channel by name
func (c *Channel) GetChannelByName(name string) (ChannelResult, error) {
	var channelResult ChannelResult

	return channelResult, fmt.Errorf(
		"Unable to find channel %s",
		name,
	)
}

// UpdateChannelByName updates a channel by name
func (c *Channel) UpdateChannelByName(channel ChannelResult) (bool, error) {
	return true, nil
}
