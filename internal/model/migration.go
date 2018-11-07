// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package model

// Migration struct
type Migration struct {
	ID        int
	Migration string
	CreatedAt int
}

// Insert inserts records
func (e *Migration) Insert() (error, bool) {
	return nil, true
}

// Get gets records
func (e *Migration) Get() (error, bool) {
	return nil, true
}

// Count counts records
func (e *Migration) Count() (error, bool) {
	return nil, true
}

// Delete deletes records
func (e *Migration) Delete() (error, bool) {
	return nil, true
}
