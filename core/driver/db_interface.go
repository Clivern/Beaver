// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"go.etcd.io/etcd/clientv3"
)

// Database interface
type Database interface {
	Connect() error
	IsConnected() bool
	Put(key, value string) error
	PutWithLease(key, value string, leaseID clientv3.LeaseID) error
	Get(key string) (map[string]string, error)
	Delete(key string) (int64, error)
	CreateLease(seconds int64) (clientv3.LeaseID, error)
	RenewLease(leaseID clientv3.LeaseID) error
	GetKeys(key string) ([]string, error)
	Exists(key string) (bool, error)
	Close()
}
