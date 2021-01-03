// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/clivern/beaver/core/util"

	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
)

// Etcd driver
type Etcd struct {
	client *clientv3.Client
}

// NewEtcdDriver create a new instance
func NewEtcdDriver() Database {
	return new(Etcd)
}

// Connect connect to etcd server
func (e *Etcd) Connect() error {
	var err error

	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(viper.GetString("app.database.etcd.endpoints"), ","),
		DialTimeout: time.Duration(viper.GetInt("app.database.etcd.timeout")) * time.Second,
		Username:    viper.GetString("app.database.etcd.username"),
		Password:    viper.GetString("app.database.etcd.password"),
	})

	if err != nil {
		return err
	}

	return nil
}

// IsConnected checks if there is an etcd connection
func (e *Etcd) IsConnected() bool {
	return e.client != nil
}

// Put sets a record
func (e *Etcd) Put(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("app.database.etcd.timeout"))*time.Second)

	_, err := e.client.Put(ctx, key, value)

	defer cancel()

	if err != nil {
		return err
	}

	return nil
}

// PutWithLease sets a record
func (e *Etcd) PutWithLease(key, value string, leaseID clientv3.LeaseID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("app.database.etcd.timeout"))*time.Second)

	_, err := e.client.Put(ctx, key, value, clientv3.WithLease(leaseID))

	defer cancel()

	if err != nil {
		return err
	}

	return nil
}

// Get gets a record value
func (e *Etcd) Get(key string) (map[string]string, error) {
	result := make(map[string]string, 0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("app.database.etcd.timeout"))*time.Second)

	resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())

	defer cancel()

	if err != nil {
		return result, err
	}

	for _, ev := range resp.Kvs {
		result[string(ev.Key)] = string(ev.Value)
	}

	return result, nil
}

// Delete deletes a record
func (e *Etcd) Delete(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("app.database.etcd.timeout"))*time.Second)

	dresp, err := e.client.Delete(ctx, key, clientv3.WithPrefix())

	defer cancel()

	if err != nil {
		return 0, err
	}

	return dresp.Deleted, nil
}

// CreateLease creates a lease
func (e *Etcd) CreateLease(seconds int64) (clientv3.LeaseID, error) {
	var result clientv3.LeaseID

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("app.database.etcd.timeout"))*time.Second)

	resp, err := e.client.Grant(ctx, seconds)

	defer cancel()

	if err != nil {
		return result, err
	}

	return resp.ID, nil
}

// RenewLease renews a lease
func (e *Etcd) RenewLease(leaseID clientv3.LeaseID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("app.database.etcd.timeout"))*time.Second)

	_, err := e.client.KeepAliveOnce(ctx, leaseID)

	defer cancel()

	if err != nil {
		return err
	}

	return nil
}

// GetKeys gets a record sub keys
// This method will return only the keys under one key
func (e *Etcd) GetKeys(key string) ([]string, error) {
	result := []string{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("app.database.etcd.timeout"))*time.Second)

	resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())

	defer cancel()

	if err != nil {
		return result, err
	}

	for _, ev := range resp.Kvs {
		sub := strings.Replace(string(ev.Key), util.EnsureTrailingSlash(key), "", -1)
		subKeys := strings.Split(sub, "/")
		newKey := fmt.Sprintf("%s%s", util.EnsureTrailingSlash(key), subKeys[0])

		if !util.InArray(newKey, result) {
			result = append(result, newKey)
		}
	}

	return result, nil
}

// Exists checks if a record exists
func (e *Etcd) Exists(key string) (bool, error) {
	result, err := e.Get(key)

	return len(result) > 0, err
}

// Close closes the etcd connection
func (e *Etcd) Close() {
	e.client.Close()
}
