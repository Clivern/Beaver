// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build unit

package util

import (
	"testing"

	"github.com/franela/goblin"
)

// TestInArray
func TestInArray(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#InArray", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				value      string
				array      []string
				wantResult bool
			}{
				{"a", []string{"b", "a", "c", "d"}, true},
				{"b", []string{"b", "a", "c", "d"}, true},
				{"c", []string{"b", "a", "c", "d"}, true},
				{"o", []string{"b", "a", "c", "d"}, false},
				{"f", []string{"b", "a", "c", "d"}, false},
			}

			for _, tt := range tests {
				g.Assert(InArray(tt.value, tt.array)).Equal(tt.wantResult)
			}
		})
	})
}

// TestGenerateUUID4
func TestGenerateUUID4(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#GenerateUUID4", func() {
		g.It("It should satisfy all provided test cases", func() {
			g.Assert(GenerateUUID4() != "").Equal(true)
			g.Assert(GenerateUUID4() != "").Equal(true)
			g.Assert(GenerateUUID4() != "").Equal(true)
		})
	})
}
