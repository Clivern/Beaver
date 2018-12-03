// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"github.com/clivern/beaver/internal/app/driver"
	"github.com/clivern/beaver/internal/pkg/logger"
)

// Metrics struct
type Metrics struct {
	Driver        *driver.Redis
	CorrelationID string
	Configs       int
	Channels      int
	Subscribers   int
	Clients       int
	MessageSent   int
}

// Init initialize the redis connection
func (m *Metrics) Init() bool {
	m.Driver = driver.NewRedisDriver()

	result, err := m.Driver.Connect()
	if !result {
		logger.Errorf(
			`Error while connecting to redis: %s {"correlationId":"%s"}`,
			err.Error(),
			m.CorrelationID,
		)
		return false
	}
	return true
}

// Trace get all metrics values
func (m *Metrics) Trace() bool {
	return true
}

// GetConfigs get configs count value
func (m *Metrics) GetConfigs() int {
	return m.Configs
}

// GetChannels get channels count value
func (m *Metrics) GetChannels() int {
	return m.Channels
}

// GetSubscribers get subscribers count value
func (m *Metrics) GetSubscribers() int {
	return m.Subscribers
}

// GetClients get clients count value
func (m *Metrics) GetClients() int {
	return m.Clients
}
