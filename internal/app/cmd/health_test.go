// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/nbio/st"
	"testing"
)

// TestHealthStatus test cases
func TestHealthStatus(t *testing.T) {
	ok, err := HealthStatus()
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
}
