// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package api

import (
	"fmt"
	"github.com/clivern/beaver/internal/pkg/utils"
	"github.com/nbio/st"
	"os"
	"testing"
)

// TestClientAPI test cases
func TestClientAPI(t *testing.T) {
	// Setup Env Vars
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.json")

	config := &utils.Config{}
	ok, err := config.Load(configFile)

	if !ok || err != nil {
		panic(err.Error())
	}
	config.Cache()
	config.GinEnv()
	os.Setenv("LogPath", fmt.Sprintf("%s/%s", basePath, os.Getenv("LogPath")))

	clientAPI := &Client{}
	st.Expect(t, clientAPI.Init(), true)
}
