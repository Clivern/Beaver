// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

import (
	"github.com/stretchr/testify/mock"
	"go.etcd.io/etcd/clientv3"
)

// EtcdMock type
type EtcdMock struct {
	mock.Mock
}

// Connect mock
func (e *EtcdMock) Connect() error {
	args := e.Called()
	return args.Error(0)
}

// IsConnected mock
func (e *EtcdMock) IsConnected() bool {
	args := e.Called()
	return args.Bool(0)
}

// Put mock
func (e *EtcdMock) Put(key, value string) error {
	args := e.Called(key, value)
	return args.Error(0)
}

// PutWithLease mock
func (e *EtcdMock) PutWithLease(key, value string, leaseID clientv3.LeaseID) error {
	args := e.Called(key, value, leaseID)
	return args.Error(0)
}

// Get mock
func (e *EtcdMock) Get(key string) (map[string]string, error) {
	args := e.Called(key)
	return args.Get(0).(map[string]string), args.Error(1)
}

// Delete mock
func (e *EtcdMock) Delete(key string) (int64, error) {
	args := e.Called(key)
	return args.Get(0).(int64), args.Error(1)
}

// CreateLease mock
func (e *EtcdMock) CreateLease(seconds int64) (clientv3.LeaseID, error) {
	args := e.Called(seconds)
	return args.Get(0).(clientv3.LeaseID), args.Error(1)
}

// RenewLease mock
func (e *EtcdMock) RenewLease(leaseID clientv3.LeaseID) error {
	args := e.Called(leaseID)
	return args.Error(0)
}

// GetKeys mock
func (e *EtcdMock) GetKeys(key string) ([]string, error) {
	args := e.Called(key)
	return args.Get(0).([]string), args.Error(1)
}

// Exists mock
func (e *EtcdMock) Exists(key string) (bool, error) {
	args := e.Called(key)
	return args.Bool(0), args.Error(1)
}

// Close mock
func (e *EtcdMock) Close() {
	//
}
