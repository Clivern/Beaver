// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-config"
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

	err := config.LoadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	os.Setenv("BeaverBasePath", fmt.Sprintf("%s/", basePath))
	os.Setenv("PORT", config.Get("app", "port").String("8080"))
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
