// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"os"
	"strings"

	"github.com/clivern/beaver/core/driver"
)

// Node type
type Node struct {
	db driver.Cassandra
}

// NewNode creates a node instance
func NewNode(db driver.Cassandra) *Node {
	result := new(Node)
	result.db = db

	return result
}

// Alive report the node as live to etcd
func (n *Node) Alive(seconds int64) error {
	return nil
}

// GetHostname gets the hostname
func (n *Node) GetHostname() (string, error) {
	hostname, err := os.Hostname()

	if err != nil {
		return "", err
	}

	return strings.ToLower(hostname), nil
}
