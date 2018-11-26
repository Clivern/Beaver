// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/nbio/st"
	"os"
	"strings"
	"testing"
	"time"
)

// init setup stuff
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

// TestClientAPI test cases
func TestClientAPI(t *testing.T) {

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	clientResult := ClientResult{ID: "id", Ident: "ident", UUID: "uuid", Token: "token", Channels: []string{}, CreatedAt: createdAt, UpdatedAt: updatedAt}
	jsonValue, err := clientResult.ConvertToJSON()
	st.Expect(t, jsonValue, fmt.Sprintf(`{"id":"id","ident":"ident","uuid":"uuid","token":"token","channels":[],"created_at":%d,"updated_at":%d}`, createdAt, updatedAt))
	st.Expect(t, err, nil)

	ok, err := clientResult.LoadFromJSON([]byte(jsonValue))
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
	st.Expect(t, clientResult.ID, "id")
	st.Expect(t, clientResult.Ident, "ident")
	st.Expect(t, clientResult.UUID, "uuid")
	st.Expect(t, clientResult.Token, "token")
	st.Expect(t, clientResult.Channels, []string{})
	st.Expect(t, clientResult.CreatedAt, createdAt)
	st.Expect(t, clientResult.UpdatedAt, updatedAt)

	clientAPI := Client{}
	st.Expect(t, clientAPI.Init(), true)

	//Clear
	clientAPI.DeleteClientByID(clientResult.ID)

	ok, err = clientAPI.UpdateClientByID(clientResult)
	st.Expect(t, ok, false)
	st.Expect(t, err.Error(), "Trying to create non existent client id")

	ok, err = clientAPI.CreateClient(clientResult)
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	newClientResult, err := clientAPI.GetClientByID(clientResult.ID)
	st.Expect(t, clientResult.ID, newClientResult.ID)
	st.Expect(t, clientResult.Ident, newClientResult.Ident)
	st.Expect(t, clientResult.UUID, newClientResult.UUID)
	st.Expect(t, clientResult.Token, newClientResult.Token)
	st.Expect(t, clientResult.Channels, newClientResult.Channels)
	st.Expect(t, clientResult.CreatedAt, newClientResult.CreatedAt)
	st.Expect(t, clientResult.UpdatedAt, newClientResult.UpdatedAt)
	st.Expect(t, err, nil)

	newClientResult.Ident = "n-ident"

	ok, err = clientAPI.UpdateClientByID(newClientResult)
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	newClientResult, err = clientAPI.GetClientByID(clientResult.ID)
	st.Expect(t, clientResult.ID, newClientResult.ID)
	st.Expect(t, "n-ident", newClientResult.Ident)
	st.Expect(t, clientResult.UUID, newClientResult.UUID)
	st.Expect(t, clientResult.Token, newClientResult.Token)
	st.Expect(t, clientResult.Channels, newClientResult.Channels)
	st.Expect(t, clientResult.CreatedAt, newClientResult.CreatedAt)
	st.Expect(t, clientResult.UpdatedAt, newClientResult.UpdatedAt)
	st.Expect(t, err, nil)

	ok, err = clientAPI.DeleteClientByID(newClientResult.ID)
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
}
