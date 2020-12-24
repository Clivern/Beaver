// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package driver

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

// TestRedisDriver test cases
func TestRedisDriver(t *testing.T) {

	driver := NewRedisDriver()

	ok, err := driver.Connect()
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = driver.Ping()
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	// Do Clean
	driver.Del("app_name")
	driver.HTruncate("configs")

	count, err := driver.Del("app_name")
	st.Expect(t, int(count), 0)
	st.Expect(t, err, nil)

	ok, err = driver.Set("app_name", "Beaver", 0)
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = driver.Exists("app_name")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	value, err := driver.Get("app_name")
	st.Expect(t, value, "Beaver")
	st.Expect(t, err, nil)

	count, err = driver.HDel("configs", "app_name")
	st.Expect(t, int(count), 0)
	st.Expect(t, err, nil)

	ok, err = driver.HSet("configs", "app_name", "Beaver")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	ok, err = driver.HExists("configs", "app_name")
	st.Expect(t, ok, true)
	st.Expect(t, err, nil)

	value, err = driver.HGet("configs", "app_name")
	st.Expect(t, value, "Beaver")
	st.Expect(t, err, nil)

	count, err = driver.HLen("configs")
	st.Expect(t, int(count), 1)
	st.Expect(t, err, nil)

	count, err = driver.HDel("configs", "app_name")
	st.Expect(t, int(count), 1)
	st.Expect(t, err, nil)

	count, err = driver.HTruncate("configs")
	st.Expect(t, int(count), 0)
	st.Expect(t, err, nil)
}
