// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

import (
	"sync"
)

// ConcurrentMap type
type ConcurrentMap struct {
	sync.RWMutex
	items map[string]interface{}
}

// NewConcurrentMap creates a new instance of ConcurrentMap
func NewConcurrentMap() ConcurrentMap {
	return ConcurrentMap{items: make(map[string]interface{})}
}

// Gets a key from a concurrent map
func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	cm.Lock()
	defer cm.Unlock()

	value, ok := cm.items[key]

	return value, ok
}

// Sets a key in a concurrent map
func (cm *ConcurrentMap) Set(key string, value interface{}) {
	cm.Lock()
	defer cm.Unlock()

	cm.items[key] = value
}

// Delete deletes a key
func (cm *ConcurrentMap) Delete(key string) {
	cm.Lock()
	defer cm.Unlock()

	delete(cm.items, key)
}
