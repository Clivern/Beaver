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
	"time"
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

// TestGetChannel1 test case
func TestGetChannel1(t *testing.T) {

	router := gin.Default()
	router.GET("/api/channel/:name", GetChannelByName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/channel/chan_name", nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNotFound, w.Code)
}

// TestGetChannel2 test case
func TestGetChannel2(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	channelResult := api.ChannelResult{
		Name:      "chan_name",
		Type:      "type",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	incomingChannelResult := api.ChannelResult{}

	channelAPI.DeleteChannelByName("chan_name")
	channelAPI.CreateChannel(channelResult)

	router := gin.Default()
	router.GET("/api/channel/:name", GetChannelByName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/channel/chan_name", nil)
	router.ServeHTTP(w, req)

	incomingChannelResult.LoadFromJSON([]byte(w.Body.String()))

	st.Expect(t, http.StatusOK, w.Code)
	st.Expect(t, incomingChannelResult.Name, channelResult.Name)
	st.Expect(t, incomingChannelResult.Type, channelResult.Type)
	st.Expect(t, incomingChannelResult.CreatedAt, channelResult.CreatedAt)
	st.Expect(t, incomingChannelResult.UpdatedAt, channelResult.UpdatedAt)

	// Clean After
	channelAPI.DeleteChannelByName("chan_name")
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

	channelResult := api.ChannelResult{
		Name:      "chan_name",
		Type:      "type",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
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

// TestCreateChannel1 test case
func TestCreateChannel1(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	channelAPI.DeleteChannelByName("new_chan")

	router := gin.Default()
	router.POST("/api/channel", CreateChannel)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/channel", strings.NewReader(`{"name":"new_chan", "type":"public"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusCreated, w.Code)

	// Clean After
	channelAPI.DeleteChannelByName("new_chan")
}

// TestCreateChannel2 test case
func TestCreateChannel2(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	channelAPI.DeleteChannelByName("new_chan")

	router := gin.Default()
	router.POST("/api/channel", CreateChannel)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/channel", strings.NewReader(`{"name":"new_chan", "type":`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	channelAPI.DeleteChannelByName("new_chan")
}

// TestCreateChannel3 test case
func TestCreateChannel3(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	channelAPI.DeleteChannelByName("new_chan")

	router := gin.Default()
	router.POST("/api/channel", CreateChannel)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/channel", strings.NewReader(`{"name":"", "type":""}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	channelAPI.DeleteChannelByName("new_chan")
}

// TestCreateChannel4 test case
func TestCreateChannel4(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	channelResult := api.ChannelResult{
		Name:      "new_chan",
		Type:      "type",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	channelAPI.DeleteChannelByName("new_chan")
	channelAPI.CreateChannel(channelResult)

	router := gin.Default()
	router.POST("/api/channel", CreateChannel)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/channel", strings.NewReader(`{"name":"new_chan", "type":"public"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	channelAPI.DeleteChannelByName("new_chan")
}

// TestUpdateChannel1 test case
func TestUpdateChannel1(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	channelResult := api.ChannelResult{
		Name:      "new_chan",
		Type:      "type",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	channelAPI.DeleteChannelByName("new_chan")
	channelAPI.CreateChannel(channelResult)

	router := gin.Default()
	router.PUT("/api/channel/:name", UpdateChannelByName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/channel/new_chan", strings.NewReader(`{"type":"private"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusOK, w.Code)

	// Clean After
	channelAPI.DeleteChannelByName("new_chan")
}

// TestUpdateChannel2 test case
func TestUpdateChannel2(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	channelAPI.DeleteChannelByName("new_chan")

	router := gin.Default()
	router.PUT("/api/channel/:name", UpdateChannelByName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/channel/new_chan", strings.NewReader(`{"type":"private"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNotFound, w.Code)
}

// TestUpdateChannel3 test case
func TestUpdateChannel3(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	channelResult := api.ChannelResult{
		Name:      "new_chan",
		Type:      "type",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	channelAPI.DeleteChannelByName("new_chan")
	channelAPI.CreateChannel(channelResult)

	router := gin.Default()
	router.PUT("/api/channel/:name", UpdateChannelByName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/channel/new_chan", strings.NewReader(`{"type":""}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	channelAPI.DeleteChannelByName("new_chan")
}

// TestUpdateChannel4 test case
func TestUpdateChannel4(t *testing.T) {

	channelAPI := api.Channel{}
	st.Expect(t, channelAPI.Init(), true)

	// Clean Before
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	channelResult := api.ChannelResult{
		Name:      "new_chan",
		Type:      "type",
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	channelAPI.DeleteChannelByName("new_chan")
	channelAPI.CreateChannel(channelResult)

	router := gin.Default()
	router.PUT("/api/channel/:name", UpdateChannelByName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/channel/new_chan", strings.NewReader(`{"type"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)

	// Clean After
	channelAPI.DeleteChannelByName("new_chan")
}
