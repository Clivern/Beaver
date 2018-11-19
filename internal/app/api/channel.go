// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"fmt"
	"github.com/clivern/beaver/internal/app/driver"
	"github.com/clivern/beaver/internal/pkg/logger"
	"os"
	"strconv"
)

// ChannelsHashPrefix is the hash prefix
const ChannelsHashPrefix string = "beaver.channel"

// Channel struct
type Channel struct {
	Driver *driver.Redis
}

// ChannelResult struct
type ChannelResult struct {
	Name string `json:"name"`
	Type string `json:"type"`
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

// Init initialize the redis connection
func (c *Channel) Init() bool {
	DB, _ := strconv.Atoi(os.Getenv("RedisDB"))

	c.Driver = &driver.Redis{
		Addr:     os.Getenv("RedisAddr"),
		Password: os.Getenv("RedisPassword"),
		DB:       DB,
	}

	result, err := c.Driver.Connect()
	if !result {
		logger.Errorf("Error while connecting to redis: %s", err.Error())
		return false
	}
	return true
}

// CreateChannel creates a channel
func (c *Channel) CreateChannel(channel ChannelResult) (bool, error) {
	return true, nil
}

// GetChannelByName gets a channel by name
func (c *Channel) GetChannelByName(name string) (ChannelResult, error) {
	var channelResult ChannelResult

	return channelResult, nil
}

// UpdateChannelByName updates a channel by name
func (c *Channel) UpdateChannelByName(channel ChannelResult) (bool, error) {
	return true, nil
}

// DeleteChannelByName deletes a channel with name
func (c *Channel) DeleteChannelByName(name string) (bool, error) {

	deleted, err := c.Driver.HDel(ChannelsHashPrefix, name)

	if err != nil {
		logger.Errorf("Error while deleting channel %s: %s", name, err.Error())
		return false, fmt.Errorf("Error while deleting channel %s", name)
	}

	if deleted <= 0 {
		logger.Warningf("Trying to delete non existent channel %s", name)
		return false, fmt.Errorf("Trying to delete non existent channel %s", name)
	}

	c.Driver.HTruncate(fmt.Sprintf("%s.listeners", name))
	c.Driver.HTruncate(fmt.Sprintf("%s.subscribers", name))

	return true, nil
}
