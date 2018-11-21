// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"github.com/clivern/beaver/internal/app/driver"
	"github.com/clivern/beaver/internal/pkg/logger"
)

// ClientsHashPrefix is the hash prefix
const ClientsHashPrefix string = "beaver.client"

// Client struct
type Client struct {
	Driver *driver.Redis
}

// Init initialize the redis connection
func (c *Client) Init() bool {
	c.Driver = driver.NewRedisDriver()

	result, err := c.Driver.Connect()
	if !result {
		logger.Errorf("Error while connecting to redis: %s", err.Error())
		return false
	}
	return true
}
