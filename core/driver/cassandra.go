// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"context"
	"time"

	"github.com/gocql/gocql"
)

// Cassandra type
type Cassandra struct {
	Session  *gocql.Session
	Hosts    []string
	Timeout  int
	Username string
	Password string
}

// NewCassandra creates a new instance
func NewCassandra() *Cassandra {
	return &Cassandra{}
}

// WithHosts define hosts
func (c *Cassandra) WithHosts(hosts []string) *Cassandra {
	c.Hosts = hosts
	return c
}

// WithTimeout define timeout
func (c *Cassandra) WithTimeout(timeout int) *Cassandra {
	c.Timeout = timeout
	return c
}

// WithAuth define auth configs
func (c *Cassandra) WithAuth(username, password string) *Cassandra {
	c.Username = username
	c.Password = password
	return c
}

// CreateSession creates a new session
func (c *Cassandra) CreateSession() error {
	var err error

	// https://github.com/gocql/gocql/blob/master/cluster.go#L31
	cluster := gocql.NewCluster()
	cluster.Hosts = c.Hosts
	cluster.ConnectTimeout = time.Second * time.Duration(c.Timeout)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: c.Username, Password: c.Password}
	c.Session, err = cluster.CreateSession()

	return err
}

// Query query the database
func (c *Cassandra) Query(ctx context.Context, query string) *gocql.Query {
	return c.Session.Query(query).WithContext(ctx)
}

// GetSession gets the session
func (c *Cassandra) GetSession() *gocql.Session {
	return c.Session
}

// Close closes a session
func (c *Cassandra) Close() {
	c.Session.Close()
}
