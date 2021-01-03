// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build unit

package util

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/franela/goblin"
)

// TestMap
func TestMap(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#Map", func() {
		g.It("It should not panic and length is right", func() {

			var wg sync.WaitGroup
			cm := &Map{items: make(map[string]interface{})}

			for i := 0; i < 1000; i++ {
				wg.Add(1)
				go func(wg *sync.WaitGroup, cm *Map, k int) {
					cm.Set(fmt.Sprintf("record_%d", k), fmt.Sprintf("value_%d", k))
					cm.Get(fmt.Sprintf("record_%d", k))
					defer wg.Done()
				}(&wg, cm, i)
			}
			// Wait till all above go routines finish
			time.Sleep(4 * time.Second)

			for i := 0; i < 500; i++ {
				wg.Add(1)
				go func(wg *sync.WaitGroup, cm *Map, k int) {
					cm.Delete(fmt.Sprintf("record_%d", k))
					defer wg.Done()
				}(&wg, cm, i)
			}

			wg.Wait()

			g.Assert(len(cm.items)).Equal(500)
		})
	})
}
