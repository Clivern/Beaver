// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"github.com/clivern/beaver/core/driver"
)

// Stats type
type Stats struct {
	db driver.Cassandra
}

// NewStats creates a stats instance
func NewStats(db driver.Cassandra) *Stats {
	result := new(Stats)
	result.db = db

	return result
}

// GetTotalNodes gets total nodes count
func (s *Stats) GetTotalNodes() (int, error) {
	return 0, nil
}

// GetTotalChannels gets total channels count
func (s *Stats) GetTotalChannels() (int, error) {
	return 0, nil
}

// GetTotalClients gets total clients count
func (s *Stats) GetTotalClients() (int, error) {
	return 0, nil
}

// GetTotalMessages gets total messages count
func (s *Stats) GetTotalMessages() (int, error) {
	return 0, nil
}
