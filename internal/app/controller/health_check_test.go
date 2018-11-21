// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nbio/st"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func init() {
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.json")

	config := utils.Config{}
	ok, err := config.Load(configFile)

	if !ok || err != nil {
		panic(err.Error())
	}
	config.Cache()
	config.GinEnv()
	if !strings.Contains(os.Getenv("LogPath"), basePath) {
		os.Setenv("LogPath", fmt.Sprintf("%s/%s", basePath, os.Getenv("LogPath")))
	}
}

// TestHealthCheckController test case
func TestHealthCheckController(t *testing.T) {

	router := gin.Default()
	router.GET("/_healthcheck", HealthCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_healthcheck", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, 200, w.Code)
	st.Expect(t, `{"status":"ok"}`, w.Body.String())
}
