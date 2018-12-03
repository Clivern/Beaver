// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"fmt"
	"github.com/clivern/beaver/internal/app/driver"
	"github.com/clivern/beaver/internal/pkg/logger"
)

// ChannelsHashPrefix is the hash prefix
const ChannelsHashPrefix string = "beaver.channel"

// Channel struct
type Channel struct {
	Driver        *driver.Redis
	CorrelationID string
}

// ChannelResult struct
type ChannelResult struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Listeners   int    `json:"listeners"`
	Subscribers int    `json:"subscribers"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
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
	c.Driver = driver.NewRedisDriver()

	result, err := c.Driver.Connect()
	if !result {
		logger.Errorf(
			`Error while connecting to redis: %s {"correlationId":"%s"}`,
			err.Error(),
			c.CorrelationID,
		)
		return false
	}
	return true
}

// CreateChannel creates a channel
func (c *Channel) CreateChannel(channel ChannelResult) (bool, error) {
	exists, err := c.Driver.HExists(ChannelsHashPrefix, channel.Name)

	if err != nil {
		logger.Errorf(
			`Error while creating channel %s: %s {"correlationId":"%s"}`,
			channel.Name,
			err.Error(),
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Error while creating channel %s`,
			channel.Name,
		)
	}

	if exists {
		logger.Warningf(
			`Trying to create existent channel %s {"correlationId":"%s"}`,
			channel.Name,
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Trying to create existent channel %s`,
			channel.Name,
		)
	}

	result, err := channel.ConvertToJSON()

	if err != nil {
		logger.Errorf(
			`Something wrong with channel %s data: %s {"correlationId":"%s"}`,
			channel.Name,
			err.Error(),
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Something wrong with channel %s data`,
			channel.Name,
		)
	}

	_, err = c.Driver.HSet(ChannelsHashPrefix, channel.Name, result)

	if err != nil {
		logger.Errorf(
			`Error while creating channel %s: %s {"correlationId":"%s"}`,
			channel.Name,
			err.Error(),
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Error while creating channel %s`,
			channel.Name,
		)
	}

	return true, nil
}

// GetChannelByName gets a channel by name
func (c *Channel) GetChannelByName(name string) (ChannelResult, error) {
	var channelResult ChannelResult

	exists, err := c.Driver.HExists(ChannelsHashPrefix, name)

	if err != nil {
		logger.Errorf(
			`Error while getting channel %s: %s {"correlationId":"%s"}`,
			name,
			err.Error(),
			c.CorrelationID,
		)
		return channelResult, fmt.Errorf(
			`Error while getting channel %s`,
			name,
		)
	}

	if !exists {
		logger.Warningf(
			`Trying to get non existent channel %s {"correlationId":"%s"}`,
			name,
			c.CorrelationID,
		)
		return channelResult, fmt.Errorf(
			`Trying to get non existent channel %s`,
			name,
		)
	}

	value, err := c.Driver.HGet(ChannelsHashPrefix, name)

	if err != nil {
		logger.Errorf(
			`Error while getting channel %s: %s {"correlationId":"%s"}`,
			name,
			err.Error(),
			c.CorrelationID,
		)
		return channelResult, fmt.Errorf(
			`Error while getting channel %s`,
			name,
		)
	}

	_, err = channelResult.LoadFromJSON([]byte(value))

	if err != nil {
		logger.Errorf(
			`Error while getting channel %s: %s {"correlationId":"%s"}`,
			name,
			err.Error(),
			c.CorrelationID,
		)
		return channelResult, fmt.Errorf(
			`Error while getting channel %s`,
			name,
		)
	}

	return channelResult, nil
}

// UpdateChannelByName updates a channel by name
func (c *Channel) UpdateChannelByName(channel ChannelResult) (bool, error) {
	exists, err := c.Driver.HExists(ChannelsHashPrefix, channel.Name)

	if err != nil {
		logger.Errorf(
			`Error while updating channel %s: %s {"correlationId":"%s"}`,
			channel.Name,
			err.Error(),
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Error while updating channel %s`,
			channel.Name,
		)
	}

	if !exists {
		logger.Warningf(
			`Trying to create non existent channel %s {"correlationId":"%s"}`,
			channel.Name,
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Trying to create non existent channel %s`,
			channel.Name,
		)
	}

	result, err := channel.ConvertToJSON()

	if err != nil {
		logger.Errorf(
			`Something wrong with channel %s data: %s {"correlationId":"%s"}`,
			channel.Name,
			err.Error(),
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Something wrong with channel %s data`,
			channel.Name,
		)
	}

	_, err = c.Driver.HSet(ChannelsHashPrefix, channel.Name, result)

	if err != nil {
		logger.Errorf(
			`Error while updating channel %s: %s {"correlationId":"%s"}`,
			channel.Name,
			err.Error(),
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Error while updating channel %s`,
			channel.Name,
		)
	}

	return true, nil
}

// DeleteChannelByName deletes a channel with name
func (c *Channel) DeleteChannelByName(name string) (bool, error) {
	deleted, err := c.Driver.HDel(ChannelsHashPrefix, name)

	if err != nil {
		logger.Errorf(
			`Error while deleting channel %s: %s {"correlationId":"%s"}`,
			name,
			err.Error(),
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Error while deleting channel %s`,
			name,
		)
	}

	if deleted <= 0 {
		logger.Warningf(
			`Trying to delete non existent channel %s {"correlationId":"%s"}`,
			name,
			c.CorrelationID,
		)
		return false, fmt.Errorf(
			`Trying to delete non existent channel %s`,
			name,
		)
	}

	c.Driver.HTruncate(fmt.Sprintf("%s.listeners", name))
	c.Driver.HTruncate(fmt.Sprintf("%s.subscribers", name))

	return true, nil
}

// DecrementListeners decrement listeners
func (c *Channel) DecrementListeners(name string) bool {
	var channelResult ChannelResult

	exists, err := c.Driver.HExists(ChannelsHashPrefix, name)

	if err != nil || !exists {
		return false
	}

	value, err := c.Driver.HGet(ChannelsHashPrefix, name)

	if err != nil {
		return false
	}

	channelResult.LoadFromJSON([]byte(value))
	channelResult.Listeners--

	if channelResult.Listeners < 0 {
		channelResult.Listeners = 0
	}

	result, err := channelResult.ConvertToJSON()

	if err != nil {
		return false
	}

	_, err = c.Driver.HSet(ChannelsHashPrefix, name, result)

	if err != nil {
		return false
	}

	return true
}

// IncrementListeners increment listeners
func (c *Channel) IncrementListeners(name string) bool {
	var channelResult ChannelResult

	exists, err := c.Driver.HExists(ChannelsHashPrefix, name)

	if err != nil || !exists {
		return false
	}

	value, err := c.Driver.HGet(ChannelsHashPrefix, name)

	if err != nil {
		return false
	}

	channelResult.LoadFromJSON([]byte(value))
	channelResult.Listeners++

	if channelResult.Listeners < 0 {
		channelResult.Listeners = 0
	}

	result, err := channelResult.ConvertToJSON()

	if err != nil {
		return false
	}

	_, err = c.Driver.HSet(ChannelsHashPrefix, name, result)

	if err != nil {
		return false
	}

	return true
}

// DecrementSubscribers decrement subscribers
func (c *Channel) DecrementSubscribers(name string) bool {
	var channelResult ChannelResult

	exists, err := c.Driver.HExists(ChannelsHashPrefix, name)

	if err != nil || !exists {
		return false
	}

	value, err := c.Driver.HGet(ChannelsHashPrefix, name)

	if err != nil {
		return false
	}

	channelResult.LoadFromJSON([]byte(value))
	channelResult.Subscribers--

	if channelResult.Subscribers < 0 {
		channelResult.Subscribers = 0
	}

	result, err := channelResult.ConvertToJSON()

	if err != nil {
		return false
	}

	_, err = c.Driver.HSet(ChannelsHashPrefix, name, result)

	if err != nil {
		return false
	}

	return true
}

// IncrementSubscribers increment subscribers
func (c *Channel) IncrementSubscribers(name string) bool {
	var channelResult ChannelResult

	exists, err := c.Driver.HExists(ChannelsHashPrefix, name)

	if err != nil || !exists {
		return false
	}

	value, err := c.Driver.HGet(ChannelsHashPrefix, name)

	if err != nil {
		return false
	}

	channelResult.LoadFromJSON([]byte(value))
	channelResult.Subscribers++

	if channelResult.Subscribers < 0 {
		channelResult.Subscribers = 0
	}

	result, err := channelResult.ConvertToJSON()

	if err != nil {
		return false
	}

	_, err = c.Driver.HSet(ChannelsHashPrefix, name, result)

	if err != nil {
		return false
	}

	return true
}

// ResetSubscribers resets subscribers
func (c *Channel) ResetSubscribers(name string) bool {
	var channelResult ChannelResult

	exists, err := c.Driver.HExists(ChannelsHashPrefix, name)

	if err != nil || !exists {
		return false
	}

	value, err := c.Driver.HGet(ChannelsHashPrefix, name)

	if err != nil {
		return false
	}

	channelResult.LoadFromJSON([]byte(value))
	channelResult.Subscribers = 0

	result, err := channelResult.ConvertToJSON()

	if err != nil {
		return false
	}

	_, err = c.Driver.HSet(ChannelsHashPrefix, name, result)

	if err != nil {
		return false
	}

	return true
}

// ResetListeners resets listeners
func (c *Channel) ResetListeners(name string) bool {
	var channelResult ChannelResult

	exists, err := c.Driver.HExists(ChannelsHashPrefix, name)

	if err != nil || !exists {
		return false
	}

	value, err := c.Driver.HGet(ChannelsHashPrefix, name)

	if err != nil {
		return false
	}

	channelResult.LoadFromJSON([]byte(value))
	channelResult.Listeners = 0

	result, err := channelResult.ConvertToJSON()

	if err != nil {
		return false
	}

	_, err = c.Driver.HSet(ChannelsHashPrefix, name, result)

	if err != nil {
		return false
	}

	return true
}

// ChannelsExist checks if channels exist
func (c *Channel) ChannelsExist(channels []string) (bool, error) {
	for _, channel := range channels {
		exists, err := c.Driver.HExists(ChannelsHashPrefix, channel)

		if err != nil {
			logger.Errorf(
				`Error while getting channel %s: %s {"correlationId":"%s"}`,
				channel,
				err.Error(),
				c.CorrelationID,
			)
			return false, fmt.Errorf(
				`Error while getting channel %s`,
				channel,
			)
		}

		if !exists {
			logger.Infof(
				`Channel %s not exist {"correlationId":"%s"}`,
				channel,
				c.CorrelationID,
			)
			return false, fmt.Errorf(
				`Channel %s not exist`,
				channel,
			)
		}
	}

	return true, nil
}

// ChannelExist checks if channel exists
func (c *Channel) ChannelExist(channel string) (bool, error) {
	return c.ChannelsExist([]string{channel})
}
