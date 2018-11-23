// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"github.com/clivern/beaver/internal/app/driver"
	"github.com/clivern/beaver/internal/pkg/logger"
)

// ClientsHashPrefix is the hash prefix
const ClientsHashPrefix string = "beaver.client"

// Client struct
type Client struct {
	Driver        *driver.Redis
	CorrelationID string
}

// ClientResult struct
type ClientResult struct {
	ID        string `json:"id"` // ident:uuid
	Token     string `json:"token"`
	CreatedAt int64  `json:"created_at"`
}

// LoadFromJSON load object from json
func (c *ClientResult) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON converts object to json
func (c *ClientResult) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Init initialize the redis connection
func (c *Client) Init() bool {
	c.Driver = driver.NewRedisDriver()

	result, err := c.Driver.Connect()
	if !result {
		logger.Errorf(`Error while connecting to redis: %s {"correlationId":"%s"}`, err.Error(), c.CorrelationID)
		return false
	}
	return true
}
