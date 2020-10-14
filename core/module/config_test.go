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

// TestConfigAPI test cases
func TestConfigAPI(t *testing.T) {

	configResult := ConfigResult{Key: "key", Value: "value"}
	jsonValue, err := configResult.ConvertToJSON()
	st.Expect(t, jsonValue, `{"key":"key","value":"value"}`)
	st.Expect(t, err, nil)

	ok, err := configResult.LoadFromJSON([]byte(jsonValue))
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)
	st.Expect(t, configResult.Key, "key")
	st.Expect(t, configResult.Value, "value")

	configAPI := Config{}
	st.Expect(t, configAPI.Init(), true)

	ok, err = configAPI.CreateConfig("key", "value")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = configAPI.CreateConfig("key", "value")
	st.Expect(t, ok, false)
	st.Expect(t, err.Error(), "Trying to create existent config key")

	value, err := configAPI.GetConfigByKey("key")
	st.Expect(t, value, "value")
	st.Expect(t, err, nil)

	value, err = configAPI.GetConfigByKey("not_exist")
	st.Expect(t, value, "")
	st.Expect(t, err.Error(), "Trying to get non existent config not_exist")

	ok, err = configAPI.UpdateConfigByKey("key", "new_value")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = configAPI.UpdateConfigByKey("not_exist", "new_value")
	st.Expect(t, ok, false)
	st.Expect(t, err.Error(), "Trying to update non existent config not_exist")

	value, err = configAPI.GetConfigByKey("key")
	st.Expect(t, value, "new_value")
	st.Expect(t, err, nil)

	ok, err = configAPI.DeleteConfigByKey("key")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = configAPI.DeleteConfigByKey("not_exist")
	st.Expect(t, ok, false)
	st.Expect(t, err.Error(), "Trying to delete non existent config not_exist")
}
