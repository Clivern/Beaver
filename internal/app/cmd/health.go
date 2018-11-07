// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
)

// Health struct
type Health struct {
}

// Status check the current app health. Make it compatible with process managers like systemd & docker
func (e *Health) Status() (error, bool) {
	fmt.Println("Status")
	return nil, true
}
