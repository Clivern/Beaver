// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"fmt"
	"os"
	"strings"

	"github.com/clivern/beaver/core/driver"

	"github.com/spf13/viper"
)

// Node type
type Node struct {
	db driver.Database
}

// NewNode creates a node instance
func NewNode(db driver.Database) *Node {
	result := new(Node)
	result.db = db

	return result
}

// Alive report the node as live to etcd
func (n *Node) Alive(seconds int64) error {
	hostname, err := n.GetHostname()

	if err != nil {
		return err
	}

	key := fmt.Sprintf(
		"%s/node/%s__%s",
		viper.GetString("app.database.etcd.databaseName"),
		hostname,
		viper.GetString("app.name"),
	)

	leaseID, err := n.db.CreateLease(seconds)

	if err != nil {
		return err
	}

	err = n.db.PutWithLease(fmt.Sprintf("%s/state", key), "alive", leaseID)

	if err != nil {
		return err
	}

	err = n.db.PutWithLease(fmt.Sprintf("%s/url", key), viper.GetString("app.url"), leaseID)

	if err != nil {
		return err
	}

	err = n.db.PutWithLease(fmt.Sprintf("%s/load", key), "0", leaseID)

	if err != nil {
		return err
	}

	err = n.db.RenewLease(leaseID)

	if err != nil {
		return err
	}

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
