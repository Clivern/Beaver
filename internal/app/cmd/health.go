// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
)

// HealthStatus check the current app health. Make it compatible with process managers like systemd & docker
func HealthStatus() (bool, error) {
	fmt.Println("HealthStatus")
	return true, nil
}
