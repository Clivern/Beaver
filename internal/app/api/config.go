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

// ConfigsHashPrefix is the hash prefix
const ConfigsHashPrefix string = "beaver.config"

// Config struct
type Config struct {
	Driver *driver.Redis
}

// ConfigResult struct
type ConfigResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// LoadFromJSON load object from json
func (c *ConfigResult) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON converts object to json
func (c *ConfigResult) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Init initialize the redis connection
func (c *Config) Init() bool {
	c.Driver = driver.NewRedisDriver()

	result, err := c.Driver.Connect()
	if !result {
		logger.Errorf("Error while connecting to redis: %s", err.Error())
		return false
	}
	return true
}

// CreateConfig creates a config
func (c *Config) CreateConfig(key string, value string) (bool, error) {
	exists, err := c.Driver.HExists(ConfigsHashPrefix, key)

	if err != nil {
		logger.Errorf("Error while creating config %s: %s", key, err.Error())
		return false, fmt.Errorf("Error while creating config %s", key)
	}

	if exists {
		logger.Warningf("Trying to create existent config %s", key)
		return false, fmt.Errorf("Trying to create existent config %s", key)
	}

	_, err = c.Driver.HSet(ConfigsHashPrefix, key, value)

	if err != nil {
		logger.Errorf("Error while creating config %s: %s", key, err.Error())
		return false, fmt.Errorf("Error while creating config %s", key)
	}

	return true, nil
}

// GetConfigByKey gets a config value with key
func (c *Config) GetConfigByKey(key string) (string, error) {

	exists, err := c.Driver.HExists(ConfigsHashPrefix, key)

	if err != nil {
		logger.Errorf("Error while getting config %s: %s", key, err.Error())
		return "", fmt.Errorf("Error while getting config %s", key)
	}

	if !exists {
		logger.Warningf("Trying to get non existent config %s", key)
		return "", fmt.Errorf("Trying to get non existent config %s", key)
	}

	value, err := c.Driver.HGet(ConfigsHashPrefix, key)

	if err != nil {
		logger.Errorf("Error while getting config %s: %s", key, err.Error())
		return "", fmt.Errorf("Error while getting config %s", key)
	}

	return value, nil
}

// UpdateConfigByKey updates a config with key
func (c *Config) UpdateConfigByKey(key string, value string) (bool, error) {

	exists, err := c.Driver.HExists(ConfigsHashPrefix, key)

	if err != nil {
		logger.Errorf("Error while updating config %s: %s", key, err.Error())
		return false, fmt.Errorf("Error while updating config %s", key)
	}

	if !exists {
		logger.Warningf("Trying to update non existent config %s", key)
		return false, fmt.Errorf("Trying to update non existent config %s", key)
	}

	_, err = c.Driver.HSet(ConfigsHashPrefix, key, value)

	if err != nil {
		logger.Errorf("Error while updating config %s: %s", key, err.Error())
		return false, fmt.Errorf("Error while updating config %s", key)
	}

	return true, nil

}

// DeleteConfigByKey deletes a config with key
func (c *Config) DeleteConfigByKey(key string) (bool, error) {

	deleted, err := c.Driver.HDel(ConfigsHashPrefix, key)

	if err != nil {
		logger.Errorf("Error while deleting config %s: %s", key, err.Error())
		return false, fmt.Errorf("Error while deleting config %s", key)
	}

	if deleted <= 0 {
		logger.Warningf("Trying to delete non existent config %s", key)
		return false, fmt.Errorf("Trying to delete non existent config %s", key)
	}

	return true, nil
}
