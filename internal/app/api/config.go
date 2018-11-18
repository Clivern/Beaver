// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

// ConfigHash is the hash prefix
const ConfigsHashPrefix string = "beaver.config"

// Config struct
type Config struct {
	Key   string
	Value string
}

// CreateConfig creates a config
func (c *Config) CreateConfig(key string, value string) (string, error) {
	return "", nil
}

// GetConfigByKey gets a config value with key
func (c *Config) GetConfigByKey(key string) (string, error) {
	return "", nil
}

// UpdateConfigByKey updates a config with key
func (c *Config) UpdateConfigByKey(key string, value string) (bool, error) {
	return true, nil
}

// DeleteConfigByKey deletes a config with key
func (c *Config) DeleteConfigByKey(key string) (bool, error) {
	return true, nil
}
