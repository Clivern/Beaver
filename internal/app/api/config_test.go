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
)

// TestConfigsAPI test cases
func TestConfigsAPI(t *testing.T) {
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
	if !strings.Contains(os.Getenv("LogPath"), basePath) {
		os.Setenv("LogPath", fmt.Sprintf("%s/%s", basePath, os.Getenv("LogPath")))
	}

	configResult := &ConfigResult{Key: "key", Value: "value"}
	jsonValue, err := configResult.ConvertToJSON()
	st.Expect(t, jsonValue, `{"key":"key","value":"value"}`)
	st.Expect(t, err, nil)

	ok, err = configResult.LoadFromJSON([]byte(jsonValue))
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
	st.Expect(t, configResult.Key, "key")
	st.Expect(t, configResult.Value, "value")

	configAPI := &Config{}
	st.Expect(t, configAPI.Init(), true)

	ok, err = configAPI.CreateConfig("key", "value")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = configAPI.CreateConfig("key", "value")
	st.Expect(t, ok, false)

	value, err := configAPI.GetConfigByKey("key")
	st.Expect(t, value, "value")
	st.Expect(t, err, nil)

	value, err = configAPI.GetConfigByKey("not_exist")
	st.Expect(t, value, "")

	ok, err = configAPI.UpdateConfigByKey("key", "new_value")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = configAPI.UpdateConfigByKey("not_exist", "new_value")
	st.Expect(t, ok, false)

	value, err = configAPI.GetConfigByKey("key")
	st.Expect(t, value, "new_value")
	st.Expect(t, err, nil)

	ok, err = configAPI.DeleteConfigByKey("key")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = configAPI.DeleteConfigByKey("not_exist")
	st.Expect(t, ok, false)
}
