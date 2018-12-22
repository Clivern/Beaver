// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"github.com/nbio/st"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// init setup stuff
func init() {
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.yml")

	config.Load(file.NewSource(
		file.WithPath(configFile),
	))

	if config.Get("app", "mode").String("") == "" {
		panic("Error! Config file not loaded")
	}

	os.Setenv("BeaverBasePath", fmt.Sprintf("%s/", basePath))
	os.Setenv("PORT", config.Get("app", "port").String("8080"))
}

// TestMetricsController test case
func TestMetricsController(t *testing.T) {

	router := gin.Default()
	router.GET("/api/metrics", GetMetrics)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/metrics", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, 200, w.Code)
	st.Expect(t, `{"status":"ok"}`, w.Body.String())
}
