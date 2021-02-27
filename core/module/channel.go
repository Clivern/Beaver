// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package module

import (
	"fmt"

	"github.com/clivern/beaver/core/driver"
)

// ChannelModule type
type ChannelModule struct {
	db driver.Cassandra
}

// ChannelModel struct
type ChannelModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// NewChannel creates a channel instance
func NewChannel(db driver.Cassandra) *ChannelModule {
	result := new(ChannelModule)
	result.db = db

	return result
}

// ChannelsExist checks if channels exist
func (c *ChannelModule) ChannelsExist(channels []string) (bool, error) {
	return true, nil
}

// ChannelExist checks if a channel exists
func (c *ChannelModule) ChannelExist(name string) (bool, error) {
	return false, nil
}

// DeleteChannelByName deletes a channel with name
func (c *ChannelModule) DeleteChannelByName(name string) (bool, error) {
	return true, nil
}

// GetSubscribers gets a list of subscribers with channel name (all subscribers)
func (c *ChannelModule) GetSubscribers(name string) ([]string, error) {
	return []string{}, nil
}

// CountSubscribers counts channel subscribers (all subscribers)
func (c *ChannelModule) CountSubscribers(name string) (int, error) {
	return 0, nil
}

// GetListeners gets a list of listeners with channel name (online subscribers)
func (c *ChannelModule) GetListeners(name string) ([]string, error) {
	return []string{}, nil
}

// CountListeners counts channel listeners (online subscribers)
func (c *ChannelModule) CountListeners(name string) (int, error) {
	return 0, nil
}

// isSubscriberOnline checks if subscriber is online
func (c *ChannelModule) isSubscriberOnline(uuid string) (bool, error) {
	return false, nil
}

// CreateChannel creates a channel
func (c *ChannelModule) CreateChannel(channel ChannelModel) (bool, error) {
	return true, nil
}

// GetChannelByName gets a channel by name
func (c *ChannelModule) GetChannelByName(name string) (ChannelModel, error) {
	var channelResult ChannelModel

	return channelResult, fmt.Errorf(
		"Unable to find channel %s",
		name,
	)
}

// UpdateChannelByName updates a channel by name
func (c *ChannelModule) UpdateChannelByName(channel ChannelModel) (bool, error) {
	return true, nil
}
