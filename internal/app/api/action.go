// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"github.com/clivern/beaver/internal/app/driver"
	"github.com/clivern/beaver/internal/pkg/logger"
)

// Action struct
type Action struct {
	Driver        *driver.Redis
	CorrelationID string
}

// Init initialize the redis connection
func (c *Action) Init() bool {
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
