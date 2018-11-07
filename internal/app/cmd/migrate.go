// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
)

// Migrate struct
type Migrate struct {
}

// Up runs the up migrations
func (e *Migrate) Up() (error, bool) {
	fmt.Println("Up")
	return nil, true
}

// Down runs the down migrations
func (e *Migrate) Down() (error, bool) {
	fmt.Println("Down")
	return nil, true
}

// Status shows migration status
func (e *Migrate) Status() (error, bool) {
	fmt.Println("Status")
	return nil, true
}
