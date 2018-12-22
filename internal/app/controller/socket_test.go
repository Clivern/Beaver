// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	_ "github.com/nbio/st"
	_ "net/http"
	_ "net/http/httptest"
	"os"
	_ "testing"
)

// init setup stuff
func init() {
	basePath := fmt.Sprintf("%s/src/github.com/clivern/beaver", os.Getenv("GOPATH"))
	configFile := fmt.Sprintf("%s/%s", basePath, "config.test.yml")

	config.Load(file.NewSource(
		file.WithPath(configFile),
	))

	if config.Get("app", "mode").String("") == "" {
		panic("Error! Config file not loaded")
	}

	os.Setenv("BeaverBasePath", fmt.Sprintf("%s/", basePath))
	os.Setenv("PORT", config.Get("app", "port").String("8080"))
}
