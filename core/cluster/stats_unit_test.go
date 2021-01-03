// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build unit

package cluster

import (
	"fmt"
	"testing"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
)

// TestUnitStatsGetTotalNodes test cases
func TestUnitStatsGetTotalNodes(t *testing.T) {

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("#Stats.GetTotalNodes", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				inputKey   string
				inputKeys  []string
				inputError error
				wantCount  int
				wantError  bool
			}{
				{"beaver_v2/node", []string{"a", "b", "c"}, nil, 3, false},
				{"beaver_v2/node", []string{"a", "b", "c", "d"}, nil, 4, false},
				{"beaver_v2/node", []string{}, nil, 0, false},
				{"beaver_v2/node", []string{"a", "b", "c", "d"}, fmt.Errorf("Error1"), 0, true},
				{"beaver_v2/node", []string{}, fmt.Errorf("Error1"), 0, true},
			}

			for _, tt := range tests {
				e := new(driver.EtcdMock)
				s := NewStats(e)

				e.On("GetKeys", tt.inputKey).Return(tt.inputKeys, tt.inputError)
				count, err := s.GetTotalNodes()

				g.Assert(count).Equal(tt.wantCount)
				g.Assert(err != nil).Equal(tt.wantError)
			}
		})
	})
}

// TestUnitStatsGetTotalChannels test cases
func TestUnitStatsGetTotalChannels(t *testing.T) {

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("#Stats.GetTotalChannels", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				inputKey   string
				inputKeys  []string
				inputError error
				wantCount  int
				wantError  bool
			}{
				{"beaver_v2/channel", []string{"a", "b", "c"}, nil, 3, false},
				{"beaver_v2/channel", []string{"a", "b", "c", "d"}, nil, 4, false},
				{"beaver_v2/channel", []string{}, nil, 0, false},
				{"beaver_v2/channel", []string{"a", "b", "c", "d"}, fmt.Errorf("Error1"), 0, true},
				{"beaver_v2/channel", []string{}, fmt.Errorf("Error1"), 0, true},
			}

			for _, tt := range tests {
				e := new(driver.EtcdMock)
				s := NewStats(e)

				e.On("GetKeys", tt.inputKey).Return(tt.inputKeys, tt.inputError)
				count, err := s.GetTotalChannels()

				g.Assert(count).Equal(tt.wantCount)
				g.Assert(err != nil).Equal(tt.wantError)
			}
		})
	})
}

// TestUnitStatsGetTotalClients test cases
func TestUnitStatsGetTotalClients(t *testing.T) {

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("#Stats.GetTotalClients", func() {
		g.It("It should satisfy all provided test cases", func() {
			var tests = []struct {
				inputKey   string
				inputKeys  []string
				inputError error
				wantCount  int
				wantError  bool
			}{
				{"beaver_v2/client", []string{"a", "b", "c"}, nil, 3, false},
				{"beaver_v2/client", []string{"a", "b", "c", "d"}, nil, 4, false},
				{"beaver_v2/client", []string{}, nil, 0, false},
				{"beaver_v2/client", []string{"a", "b", "c", "d"}, fmt.Errorf("Error1"), 0, true},
				{"beaver_v2/client", []string{}, fmt.Errorf("Error1"), 0, true},
			}

			for _, tt := range tests {
				e := new(driver.EtcdMock)
				s := NewStats(e)

				e.On("GetKeys", tt.inputKey).Return(tt.inputKeys, tt.inputError)
				count, err := s.GetTotalClients()

				g.Assert(count).Equal(tt.wantCount)
				g.Assert(err != nil).Equal(tt.wantError)
			}
		})
	})
}
