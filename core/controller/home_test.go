// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build unit

package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
)

// TestIndex test cases
func TestIndex(t *testing.T) {

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	gin.SetMode(gin.ReleaseMode)

	g := goblin.Goblin(t)

	g.Describe("#Index", func() {
		g.It("It should return the expected response and status code", func() {
			r := gin.Default()
			r.GET("/", Index)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
		})
	})
}
