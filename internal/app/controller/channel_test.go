// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"github.com/clivern/beaver/internal/app/api"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/nbio/st"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

// init setup driver
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

// TestDeleteChannel1 test case
func TestDeleteChannel1(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	channelAPI.DeleteChannelByName("chan_name")

	router := gin.Default()
	router.DELETE("/api/channel/:name", DeleteChannelByName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/channel/chan_name", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNotFound, w.Code)
}

// TestDeleteChannel2 test case
func TestDeleteChannel2(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	channelResult := api.ChannelResult{Name: "chan_name", Type: "type", Listeners: 1, Subscribers: 1, CreatedAt: createdAt, UpdatedAt: updatedAt}
	channelAPI.DeleteChannelByName("chan_name")
	channelAPI.CreateChannel(channelResult)

	router := gin.Default()
	router.DELETE("/api/channel/:name", DeleteChannelByName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/channel/chan_name", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNoContent, w.Code)

	// Clean After
	channelAPI.DeleteChannelByName("chan_name")
}
