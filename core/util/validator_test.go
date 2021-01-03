// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build unit

package util

import (
	"testing"

	"github.com/franela/goblin"
)

// TestValidator
func TestValidator(t *testing.T) {
	g := goblin.Goblin(t)
	v := &Validator{}

	g.Describe("#Validator.IsIn", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				item       string
				list       []string
				wantResult bool
			}{
				{"a", []string{"b", "a", "c", "d"}, true},
				{"b", []string{"b", "a", "c", "d"}, true},
				{"c", []string{"b", "a", "c", "d"}, true},
				{"o", []string{"b", "a", "c", "d"}, false},
				{"f", []string{"b", "a", "c", "d"}, false},
			}

			for _, tt := range tests {
				g.Assert(v.IsIn(tt.item, tt.list)).Equal(tt.wantResult)
			}
		})
	})

	g.Describe("#Validator.IsSlug", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				slug       string
				min        int
				max        int
				wantResult bool
			}{
				{"abdg_shdt", 1, 9, true},
				{"abdg_shdtki", 1, 9, false},
				{"abdG_shdt", 1, 9, false},
				{"ab78_s98t", 1, 9, true},
				{"", 1, 9, false},
			}

			for _, tt := range tests {
				g.Assert(v.IsSlug(tt.slug, tt.min, tt.max)).Equal(tt.wantResult)
			}
		})
	})

	g.Describe("#Validator.IsSlugs", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				slugs      []string
				min        int
				max        int
				wantResult bool
			}{
				{[]string{"abdg_shdt", "ab78_s98t"}, 1, 9, true},
				{[]string{"abdg_shdt", "ab78_s98t", "abdg_shdtki"}, 1, 9, false},
				{[]string{"abdg_shdt", "ab78_s98t", ""}, 1, 9, false},
				{[]string{"abdg_shdt", "ab78_s98t", "abdG_shdt"}, 1, 9, false},
			}

			for _, tt := range tests {
				g.Assert(v.IsSlugs(tt.slugs, tt.min, tt.max)).Equal(tt.wantResult)
			}
		})
	})

	g.Describe("#Validator.IsEmpty", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				item       string
				wantResult bool
			}{
				{"", true},
				{" ", true},
				{"\n", true},
				{"o", false},
			}

			for _, tt := range tests {
				g.Assert(v.IsEmpty(tt.item)).Equal(tt.wantResult)
			}
		})
	})

	g.Describe("#Validator.IsUUID4", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				item       string
				wantResult bool
			}{
				{"2c886455-24ff-4100-875b-04a95b7ce4a2", true},
				{"ea452311-0281-4caa-afe5-1a2d5042da67", true},
				{"8afee4b7-7b13-4880-8045-8b2c78cefb73", true},
				{"8afee4b7-7b13-4880-8045-8b2c78c", false},
			}

			for _, tt := range tests {
				g.Assert(v.IsUUID4(tt.item)).Equal(tt.wantResult)
			}
		})
	})

	g.Describe("#Validator.IsJSON", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				item       string
				wantResult bool
			}{
				{"{\"test\":\"test\"}", true},
				{"{\"test\":\"test}", false},
				{"{\"test\":\"test\", \"de\":1}", true},
			}

			for _, tt := range tests {
				g.Assert(v.IsJSON(tt.item)).Equal(tt.wantResult)
			}
		})
	})
}
