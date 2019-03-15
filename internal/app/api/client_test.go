// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"github.com/nbio/st"
	"github.com/spf13/viper"
	"os"
	"strconv"
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

// TestClientAPI test cases
func TestClientAPI(t *testing.T) {

	createdAt := time.Now().Unix()
	updatedAt := time.Now().Unix()

	clientResult := ClientResult{ID: "id", Token: "token", Channels: []string{}, CreatedAt: createdAt, UpdatedAt: updatedAt}
	jsonValue, err := clientResult.ConvertToJSON()
	st.Expect(t, jsonValue, fmt.Sprintf(`{"id":"id","token":"token","channels":[],"created_at":%d,"updated_at":%d}`, createdAt, updatedAt))
	st.Expect(t, err, nil)

	ok, err := clientResult.LoadFromJSON([]byte(jsonValue))
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
	st.Expect(t, clientResult.ID, "id")
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
	st.Expect(t, clientResult.Token, newClientResult.Token)
	st.Expect(t, clientResult.Channels, newClientResult.Channels)
	st.Expect(t, clientResult.CreatedAt, newClientResult.CreatedAt)
	st.Expect(t, clientResult.UpdatedAt, newClientResult.UpdatedAt)
	st.Expect(t, err, nil)

	newClientResult.Token = "n-Token"

	ok, err = clientAPI.UpdateClientByID(newClientResult)
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	newClientResult, err = clientAPI.GetClientByID(clientResult.ID)
	st.Expect(t, clientResult.ID, newClientResult.ID)
	st.Expect(t, "n-Token", newClientResult.Token)
	st.Expect(t, clientResult.Channels, newClientResult.Channels)
	st.Expect(t, clientResult.CreatedAt, newClientResult.CreatedAt)
	st.Expect(t, clientResult.UpdatedAt, newClientResult.UpdatedAt)
	st.Expect(t, err, nil)

	ok, err = clientAPI.DeleteClientByID(newClientResult.ID)
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
}
