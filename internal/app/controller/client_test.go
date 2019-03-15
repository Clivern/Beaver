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

// TestGetClient1 test case
func TestGetClient1(t *testing.T) {

	// Clean Before
	clientID := "c6da288b-9024-4578-a3c2-d23795fa1067"
	clientAPI := api.Client{}
	st.Expect(t, clientAPI.Init(), true)
	clientAPI.DeleteClientByID(clientID)

	router := gin.Default()
	router.GET("/api/client/:id", GetClientByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/client/%s", clientID), nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNotFound, w.Code)
}

// TestGetClient2 test case
func TestGetClient2(t *testing.T) {

	clientAPI := api.Client{}
	st.Expect(t, clientAPI.Init(), true)

	// Clean Before
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	clientResult := api.ClientResult{
		ID:        "c6da288b-9024-4578-a3c2-d23795fa1067",
		Token:     "eyJhbGciOiJIUzI1NiIs",
		Channels:  []string{"chat1"},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	incomingClientResult := api.ClientResult{}

	clientAPI.DeleteClientByID(clientResult.ID)
	clientAPI.CreateClient(clientResult)

	router := gin.Default()
	router.GET("/api/client/:id", GetClientByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/client/%s", clientResult.ID), nil)
	router.ServeHTTP(w, req)

	incomingClientResult.LoadFromJSON([]byte(w.Body.String()))

	st.Expect(t, http.StatusOK, w.Code)
	st.Expect(t, incomingClientResult.ID, clientResult.ID)
	st.Expect(t, incomingClientResult.Token, clientResult.Token)
	st.Expect(t, incomingClientResult.Channels, clientResult.Channels)
	st.Expect(t, incomingClientResult.CreatedAt, clientResult.CreatedAt)
	st.Expect(t, incomingClientResult.UpdatedAt, clientResult.UpdatedAt)

	// Clean After
	clientAPI.DeleteClientByID(clientResult.ID)
}

// TestDeleteClient1 test case
func TestDeleteClient1(t *testing.T) {

	// Clean Before
	clientID := "c6da288b-9024-4578-a3c2-d23795fa1067"
	clientAPI := api.Client{}
	st.Expect(t, clientAPI.Init(), true)
	clientAPI.DeleteClientByID(clientID)

	router := gin.Default()
	router.DELETE("/api/client/:id", DeleteClientByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/client/%s", clientID), nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNotFound, w.Code)
}

// TestDeleteClient2 test case
func TestDeleteClient2(t *testing.T) {

	clientAPI := api.Client{}
	st.Expect(t, clientAPI.Init(), true)

	// Clean Before
	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	clientResult := api.ClientResult{
		ID:        "c6da288b-9024-4578-a3c2-d23795fa1067",
		Token:     "eyJhbGciOiJIUzI1NiIs",
		Channels:  []string{"chat1"},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	clientAPI.DeleteClientByID(clientResult.ID)
	clientAPI.CreateClient(clientResult)

	router := gin.Default()
	router.DELETE("/api/client/:id", DeleteClientByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/client/%s", clientResult.ID), nil)
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusNoContent, w.Code)

	// Clean After
	clientAPI.DeleteClientByID(clientResult.ID)
}

// TestCreateClient1 test case
func TestCreateClient1(t *testing.T) {

	router := gin.Default()
	router.POST("/api/client", CreateClient)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/client", strings.NewReader(`{}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusCreated, w.Code)
}

// TestCreateClient2 test case
func TestCreateClient2(t *testing.T) {

	router := gin.Default()
	router.POST("/api/client", CreateClient)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/client", strings.NewReader(`{"channels":["chat"}`))
	router.ServeHTTP(w, req)

	st.Expect(t, http.StatusBadRequest, w.Code)
}
