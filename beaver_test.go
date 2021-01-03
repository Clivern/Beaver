// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// +build integration

package main

import (
	"fmt"
	"testing"

	"github.com/clivern/beaver/core/driver"
	"github.com/clivern/beaver/pkg"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TestMain
func TestMain(m *testing.M) {
	fmt.Println("====> Setup")
	baseDir := pkg.GetBaseDir("cache")
	pkg.LoadConfigs(fmt.Sprintf("%s/config.dist.yml", baseDir))

	db := driver.NewEtcdDriver()
	db.Connect()
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)

	// Cleanup
	db.Delete(viper.GetString("app.database.etcd.databaseName"))
}
