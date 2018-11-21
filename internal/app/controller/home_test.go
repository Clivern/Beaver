// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nbio/st"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomeController test case
func TestHomeController(t *testing.T) {

	router := gin.Default()
	router.GET("/", HealthCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, 200, w.Code)
	st.Expect(t, `{"status":"ok"}`, w.Body.String())
}
