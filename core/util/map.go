// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"sync"
)

// Map type
type Map struct {
	sync.RWMutex
	items map[string]interface{}
}

// NewMap creates a new instance of Map
func NewMap() Map {
	return Map{items: make(map[string]interface{})}
}

// Get a key from a concurrent map
func (cm *Map) Get(key string) (interface{}, bool) {
	cm.Lock()
	defer cm.Unlock()

	value, ok := cm.items[key]

	return value, ok
}

// Set a key in a concurrent map
func (cm *Map) Set(key string, value interface{}) {
	cm.Lock()
	defer cm.Unlock()

	cm.items[key] = value
}

// Delete deletes a key
func (cm *Map) Delete(key string) {
	cm.Lock()
	defer cm.Unlock()

	delete(cm.items, key)
}
