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
	"github.com/clivern/beaver/core/module"
	"github.com/clivern/beaver/pkg"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TestIntegrationClientController
func TestIntegrationClientController(t *testing.T) {
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
			req, _ := http.NewRequest("POST", "/api/v2/channel", strings.NewReader(`{"name": "client_test_01", "type": "private"}`))
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(201)
			g.Assert(strings.Contains(w.Body.String(), "client_test_01")).Equal(true)
		})
	})

	g.Describe("#CreateClient", func() {
		g.It("It should create a client", func() {
			r := gin.Default()
			r.POST("/api/v2/client", CreateClient)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/v2/client", strings.NewReader(`{"channels": []}`))
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(201)
			g.Assert(strings.Contains(w.Body.String(), "id")).Equal(true)
		})
	})

	g.Describe("#GetClientByID", func() {
		g.It("It should return 400", func() {
			r := gin.Default()
			r.GET("/api/v2/client/:id", GetClientByID)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v2/client/x-x-x-x", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(400)
		})

		g.It("It should return 404", func() {
			r := gin.Default()
			r.GET("/api/v2/client/:id", GetClientByID)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/v2/client/ce70de31-cb08-48f6-b849-7fdf0b02b722", nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(404)
		})
	})

	client := module.NewClient(db)
	newClient := module.GenerateClient([]string{})
	client.CreateClient(*newClient)

	g.Describe("#GetClientByID", func() {
		g.It("It should return 200", func() {
			r := gin.Default()
			r.GET("/api/v2/client/:id", GetClientByID)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v2/client/%s", newClient.ID), nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
			g.Assert(strings.Contains(w.Body.String(), "id")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), newClient.ID)).Equal(true)
		})
	})

	g.Describe("#Subscribe", func() {
		g.It("It should return 200", func() {
			r := gin.Default()
			r.PUT("/api/v2/client/:id/subscribe", Subscribe)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v2/client/%s/subscribe", newClient.ID), strings.NewReader(`{"channels": ["client_test_01"]}`))
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
		})

		g.It("It should return 200", func() {
			r := gin.Default()
			r.GET("/api/v2/client/:id", GetClientByID)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v2/client/%s", newClient.ID), nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
			g.Assert(strings.Contains(w.Body.String(), "id")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), newClient.ID)).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), "client_test_01")).Equal(true)
		})
	})

	g.Describe("#Unsubscribe", func() {
		g.It("It should return 200", func() {
			r := gin.Default()
			r.PUT("/api/v2/client/:id/unsubscribe", Unsubscribe)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v2/client/%s/unsubscribe", newClient.ID), strings.NewReader(`{"channels": ["client_test_01"]}`))
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
		})

		g.It("It should return 200", func() {
			r := gin.Default()
			r.GET("/api/v2/client/:id", GetClientByID)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v2/client/%s", newClient.ID), nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(200)
			g.Assert(strings.Contains(w.Body.String(), "id")).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), newClient.ID)).Equal(true)
			g.Assert(strings.Contains(w.Body.String(), "client_test_01")).Equal(false)
		})
	})

	g.Describe("#DeleteClientByID", func() {
		g.It("It should return 204", func() {
			r := gin.Default()
			r.DELETE("/api/v2/client/:id", DeleteClientByID)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v2/client/%s", newClient.ID), nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(204)
		})

		g.It("It should return 404", func() {
			r := gin.Default()
			r.DELETE("/api/v2/client/:id", DeleteClientByID)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v2/client/%s", newClient.ID), nil)
			r.ServeHTTP(w, req)

			g.Assert(w.Code).Equal(404)
		})
	})
}
