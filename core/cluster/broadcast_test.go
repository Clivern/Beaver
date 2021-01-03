// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build unit

package cluster

import (
	"testing"

	"github.com/franela/goblin"
)

// TestBroadcast test cases
func TestBroadcast(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Func", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				input      int
				wantOutput int
			}{
				{1, 1},
			}

			for _, tt := range tests {
				g.Assert(tt.input).Equal(tt.wantOutput)
			}
		})
	})
}
