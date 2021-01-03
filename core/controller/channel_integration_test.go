// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build integration

package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TestIntegrationChannelController
func TestIntegrationChannelController(t *testing.T) {
	// Skip if -short flag exist
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	db := driver.NewEtcdDriver()
	db.Connect()
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)

	// Cleanup
	db.Delete(viper.GetString("app.database.etcd.databaseName"))

	g := goblin.Goblin(t)

	g.Describe("#CreateChannel", func() {
		g.It("It should create a channel", func() {
			r := gin.Default()
			r.POST("/api/v2/channel", CreateChannel)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/v2/channel", strings.NewReader(`{"name": "cc_test_001", "type": "private"}`))
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(201)
			g.Assert(strings.Contains(w.Body.String(), "cc_test_001")).Equal(true)
		})
	})

	g.Describe("#CreateChannel", func() {
		g.It("It should create a channel", func() {
			r := gin.Default()
			r.POST("/api/v2/channel", CreateChannel)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/v2/channel", strings.NewReader(`{"name": "cc_test_003", "type": "public"}`))
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(201)
			g.Assert(strings.Contains(w.Body.String(), "cc_test_003")).Equal(true)
		})
	})

	g.Describe("#GetChannelByName", func() {
		g.It("It should return 404", func() {
			r := gin.Default()
			r.GET("/api/v2/channel/:name", GetChannelByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v2/channel/cc_test_002", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(404)
			g.Assert(strings.Contains(w.Body.String(), "not found")).Equal(true)
		})

		g.It("It should return 200", func() {
			r := gin.Default()
			r.GET("/api/v2/channel/:name", GetChannelByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v2/channel/cc_test_001", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
			g.Assert(strings.Contains(w.Body.String(), "cc_test_001")).Equal(true)
		})
	})

	g.Describe("#UpdateChannelByName", func() {
		g.It("It should return 404", func() {
			r := gin.Default()
			r.PUT("/api/v2/channel/:name", UpdateChannelByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", "/api/v2/channel/cc_test_002", strings.NewReader(`{"type": "public"}`))
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(404)
			g.Assert(strings.Contains(w.Body.String(), "not found")).Equal(true)
		})

		g.It("It should return 200", func() {
			r := gin.Default()
			r.PUT("/api/v2/channel/:name", UpdateChannelByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", "/api/v2/channel/cc_test_001", strings.NewReader(`{"type": "public"}`))
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
			g.Assert(strings.Contains(w.Body.String(), "public")).Equal(true)
		})
	})

	g.Describe("#DeleteChannelByName", func() {
		g.It("It should return 404", func() {
			r := gin.Default()
			r.DELETE("/api/v2/channel/:name", DeleteChannelByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/api/v2/channel/cc_test_002", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(404)
			g.Assert(strings.Contains(w.Body.String(), "not found")).Equal(true)
		})

		g.It("It should return 204", func() {
			r := gin.Default()
			r.DELETE("/api/v2/channel/:name", DeleteChannelByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/api/v2/channel/cc_test_001", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(204)
		})
	})

	g.Describe("#GetChannelSubscribersByName", func() {
		g.It("It should return 404", func() {
			r := gin.Default()
			r.GET("/api/v2/channel/:name/subscribers", GetChannelSubscribersByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v2/channel/cc_test_002/subscribers", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(404)
			g.Assert(strings.Contains(w.Body.String(), "not found")).Equal(true)
		})

		g.It("It should return 200", func() {
			r := gin.Default()
			r.GET("/api/v2/channel/:name/subscribers", GetChannelSubscribersByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v2/channel/cc_test_003/subscribers", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
			g.Assert(strings.Contains(w.Body.String(), "0")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), "count")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), "subscribers")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), "[]")).Equal(true)
		})
	})

	g.Describe("#GetChannelListenersByName", func() {
		g.It("It should return 404", func() {
			r := gin.Default()
			r.GET("/api/v2/channel/:name/listeners", GetChannelListenersByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v2/channel/cc_test_002/listeners", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(404)
			g.Assert(strings.Contains(w.Body.String(), "not found")).Equal(true)
		})

		g.It("It should return 200", func() {
			r := gin.Default()
			r.GET("/api/v2/channel/:name/listeners", GetChannelListenersByName)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v2/channel/cc_test_003/listeners", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
			g.Assert(strings.Contains(w.Body.String(), "0")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), "count")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), "listeners")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), "[]")).Equal(true)
		})
	})
}
