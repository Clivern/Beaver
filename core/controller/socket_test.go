// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/nbio/st"
	"github.com/spf13/viper"
	_ "net/http"
	_ "net/http/httptest"
	"os"
	"strconv"
	_ "testing"
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
