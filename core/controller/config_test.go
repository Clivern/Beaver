// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/clivern/beaver/internal/app/api"
	"github.com/gin-gonic/gin"
	"github.com/nbio/st"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

// init setup stuff
func init() {
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.yml")

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf(
			"Error while loading config file [%s]: %s",
			configFile,
			err.Error(),
		))
	}

	os.Setenv("BeaverBasePath", fmt.Sprintf("%s/", basePath))
	os.Setenv("PORT", strconv.Itoa(viper.GetInt("app.port")))
}

// TestGetConfig1 test case
func TestGetConfig1(t *testing.T) {

	router := gin.Default()
	router.GET("/api/config/:key", GetConfigByKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/config/config_key", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNotFound, w.Code)
}

// TestGetConfig2 test case
func TestGetConfig2(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("config_key")
	configAPI.CreateConfig("config_key", "config_value")

	router := gin.Default()
	router.GET("/api/config/:key", GetConfigByKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/config/config_key", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusOK, w.Code)
	st.Expect(t, `{"key":"config_key","value":"config_value"}`, w.Body.String())

	// Clean After
	configAPI.DeleteConfigByKey("config_key")
}

// TestCreateConfig1 test case
func TestCreateConfig1(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("new_key")

	router := gin.Default()
	router.POST("/api/config", CreateConfig)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/config", strings.NewReader(`{"key":"new_key", "value":"new_value"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusCreated, w.Code)

	// Clean After
	configAPI.DeleteConfigByKey("new_key")
}

// TestCreateConfig2 test case
func TestCreateConfig2(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("new_key")

	router := gin.Default()
	router.POST("/api/config", CreateConfig)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/config", strings.NewReader(`{"key":"new_key", "value":`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	configAPI.DeleteConfigByKey("new_key")
}

// TestCreateConfig3 test case
func TestCreateConfig3(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("new_key")

	router := gin.Default()
	router.POST("/api/config", CreateConfig)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/config", strings.NewReader(`{"key":"", "value":""}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	configAPI.DeleteConfigByKey("new_key")
}

// TestCreateConfig4 test case
func TestCreateConfig4(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("new_key")
	configAPI.CreateConfig("new_key", "new_value")

	router := gin.Default()
	router.POST("/api/config", CreateConfig)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/config", strings.NewReader(`{"key":"new_key", "value":"new_value"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	configAPI.DeleteConfigByKey("new_key")
}

// TestDeleteConfig1 test case
func TestDeleteConfig1(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("config_key")

	router := gin.Default()
	router.DELETE("/api/config/:key", DeleteConfigByKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/config/config_key", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNotFound, w.Code)
}

// TestDeleteConfig2 test case
func TestDeleteConfig2(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("config_key")
	configAPI.CreateConfig("config_key", "config_value")

	router := gin.Default()
	router.DELETE("/api/config/:key", DeleteConfigByKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/config/config_key", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNoContent, w.Code)

	// Clean After
	configAPI.DeleteConfigByKey("config_key")
}

// TestUpdateConfig1 test case
func TestUpdateConfig1(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("new_key")
	configAPI.CreateConfig("new_key", "old_value")

	router := gin.Default()
	router.PUT("/api/config/:key", UpdateConfigByKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/config/new_key", strings.NewReader(`{"value":"new_value"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusOK, w.Code)

	// Clean After
	configAPI.DeleteConfigByKey("new_key")
}

// TestUpdateConfig2 test case
func TestUpdateConfig2(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("new_key")

	router := gin.Default()
	router.PUT("/api/config/:key", UpdateConfigByKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/config/new_key", strings.NewReader(`{"value":"new_value"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)
}

// TestUpdateConfig3 test case
func TestUpdateConfig3(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("new_key")
	configAPI.CreateConfig("new_key", "old_value")

	router := gin.Default()
	router.PUT("/api/config/:key", UpdateConfigByKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/config/new_key", strings.NewReader(`{"value":""}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	configAPI.DeleteConfigByKey("new_key")
}

// TestUpdateConfig4 test case
func TestUpdateConfig4(t *testing.T) {

	configAPI := api.Config{}
	st.Expect(t, configAPI.Init(), true)

	// Clean Before
	configAPI.DeleteConfigByKey("new_key")
	configAPI.CreateConfig("new_key", "old_value")

	router := gin.Default()
	router.PUT("/api/config/:key", UpdateConfigByKey)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/config/new_key", strings.NewReader(`{"value"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	configAPI.DeleteConfigByKey("new_key")
}
