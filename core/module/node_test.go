// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build integration

package module

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
)

// TestNodeModule test cases
func TestNodeModule(t *testing.T) {
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.test.yml", baseDir))

	g := goblin.Goblin(t)

	g.Describe("#TestInsert", func() {
		g.It("It should satisfy all provided test cases", func() {

			g.Assert(Sum(tt.input)).Equal(tt.wantResult)

		})
	})
}
